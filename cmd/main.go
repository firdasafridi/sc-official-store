package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/firdasafridi/sc-official-store/api" // this would be your generated smart contract bindings

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	privateKey      = "put_your_privatekey_here"
	addressMainPool = "put_your_main_address"
)

type requestContract struct {
	Name             string   `json:"name"`
	MainBrandAddress string   `json:"mainBrandAddress"`
	District         string   `json:"district"`
	ExpiredDate      *big.Int `json:"expiredDate"`
	Active           bool     `json:"active"`
}

type extendContract struct {
	Address     string   `json:"address"`
	ExpiredDate *big.Int `json:"expiredDate"`
}

func main() {
	// address of etherum env
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

	//create auth and transaction package for deploying smart contract
	auth := getAccountAuth(client, privateKey)

	//once deploy shutdown
	//deploying smart contract
	// auth := getAccountAuth(client, privateKey)

	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}

	//auth
	// address := common.BytesToAddress([]byte(addressMainPool))

	fmt.Println(address.Hex())

	_, _ = instance, tx
	fmt.Println("instance->", instance)
	fmt.Println("tx->", tx.Hash().Hex())

	//creating api object to intract with smart contract function
	conn, err := api.NewApi(common.HexToAddress(address.Hex()), client)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/set/brand/contract/:address/:brand", func(c echo.Context) error {
		addressStr := c.Param("address")
		brand := c.Param("brand")

		address := common.BytesToAddress([]byte(addressStr))

		auth := getAccountAuth(client, privateKey)

		reply, err := conn.SetBrandContract(auth, address, brand)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/get/brand/:address", func(c echo.Context) error {
		addressStr := c.Param("address")

		address := common.BytesToAddress([]byte(addressStr))

		reply, err := conn.GetBrandFromAddress(&bind.CallOpts{}, address)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, reply)
	})

	e.POST("/set/contract/os/:address", func(c echo.Context) error {

		addressStr := c.Param("address")

		address := common.BytesToAddress([]byte(addressStr))

		var v requestContract

		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		mainBrandAddress := common.BytesToAddress([]byte(v.MainBrandAddress))

		auth := getAccountAuth(client, privateKey)

		reply, err := conn.SetContractOS(auth, address, api.OfficialStoreContractBrand{
			Name:             v.Name,
			MainBrandAddress: mainBrandAddress,
			District:         v.District,
			ExpiredDate:      v.ExpiredDate,
			Active:           v.Active,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/get/contract/os/:address", func(c echo.Context) error {

		addressStr := c.Param("address")

		address := common.BytesToAddress([]byte(addressStr))

		reply, err := conn.GetContractBrandFromAddress(&bind.CallOpts{}, address)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/enable/os/:address", func(c echo.Context) error {

		addressStr := c.Param("address")

		address := common.BytesToAddress([]byte(addressStr))

		auth := getAccountAuth(client, privateKey)

		reply, err := conn.EnableContractOS(auth, address)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/disable/os/:address", func(c echo.Context) error {

		addressStr := c.Param("address")

		address := common.BytesToAddress([]byte(addressStr))

		auth := getAccountAuth(client, privateKey)

		reply, err := conn.DisableContractOS(auth, address)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/is/active/:address/:brand", func(c echo.Context) error {

		addressStr := c.Param("address")
		brand := c.Param("brand")

		address := common.BytesToAddress([]byte(addressStr))

		tNow := new(big.Int).SetUint64(uint64(time.Now().Unix()))

		reply, err := conn.IsContractBrandActive(&bind.CallOpts{}, address, brand, tNow)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.POST("/extend/contract", func(c echo.Context) error {

		var v extendContract

		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		address := common.BytesToAddress([]byte(v.Address))

		auth := getAccountAuth(client, privateKey)

		reply, err := conn.ExtendContractOS(auth, address, v.ExpiredDate)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

//function to create auth for any account from its private key
func getAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	fmt.Println("from_address=", fromAddress)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(3000000) // in wei
	auth.GasLimit = uint64(3000000)  // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
