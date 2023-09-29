package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"errors"
	"context"
	"strconv"
    "strings"
	"math/big"
	"encoding/csv"

	"github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/tests/testdata/Analysis"

)

type ethClienter interface {
	ethereum.TransactionReader
	ethereum.ContractCaller
	bind.DeployBackend
}

func main() {

	var networks = []struct {
		Name       string
		URL        string
		ChainID    uint64
		PrivateKey string
	}{
		{Name: "Local L1", URL: "http://localhost:8545", ChainID: 1337, PrivateKey: "0xd16a1b4f7bb1948a4d37d4b5283c9b8c97f7b007d0617501293f58f5c89d25ba"},
	}
	args := os.Args
	opcode := args[1]
	log.Printf("Opcode %v", opcode)
	limits := []*big.Int{big.NewInt(0), big.NewInt(100), big.NewInt(1000), big.NewInt(10000)}
	numProbes := len(limits)
	var timings = make([]int64, numProbes)
	var gasUsed = make([]uint64, numProbes)
	var perOpcode = make([]float64, numProbes)

	txTimeout := 60 * time.Second
	ctx := context.Background()
	ethnetwork := networks[0]
	client, err := ethclient.Dial(ethnetwork.URL)
	chkErr(err)

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(ethnetwork.PrivateKey, "0x"))
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(0).SetUint64(ethnetwork.ChainID))
	log.Print(auth)
	chkErr(err)

	balance, err := client.BalanceAt(ctx, auth.From, nil)
	chkErr(err)
	log.Printf("ETH Balance for %v: %v", auth.From, balance)

	log.Printf("Sending TX to deploy Analysis SC")
	_, tx, analysisSC, err := Analysis.DeployAnalysis(auth, client)
	chkErr(err)
	err = WaitTxToBeMined(ctx, client, tx, txTimeout)
	chkErr(err)

	for i, limit := range limits {
		log.Printf("Calling Benchmark method from Analysis with limit %v", limit)
		start := time.Now()
		tx, err = analysisSC.Benchmark(auth, limit)
		chkErr(err)
		err = WaitTxToBeMined(ctx, client, tx, txTimeout)
		timeElapsed := time.Since(start)
		chkErr(err)

		gasUsed[i] = tx.Gas()
		timings[i] = int64(timeElapsed / time.Microsecond)
		fmt.Println("Counter function took", timeElapsed, "time")
	}

	overhead := timings[0]
	for i := 1; i < numProbes; i++ {
		timings[i] -= overhead
		perOpcode[i] = float64(timings[i]) / float64(limits[i].Int64())
	}

	fmt.Println("Limits:", limits)
	fmt.Println("Timings [microseconds]:", timings)
	fmt.Println("Per Opcode [microseconds]:", perOpcode)
	fmt.Println("Gas Used: ", gasUsed)
	err = writeDataToCSV(limits, timings, gasUsed, perOpcode, opcode)
	chkErr(err)
}



func WaitTxToBeMined(parentCtx context.Context, client ethClienter, tx *types.Transaction, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()
	_, err := bind.WaitMined(ctx, client, tx)
	if errors.Is(err, context.DeadlineExceeded) {
		return err
	} else if err != nil {
		log.Printf("error waiting tx %s to be mined: %w", tx.Hash(), err)
		return err
	}
	log.Printf("Transaction successfully mined: ", tx.Hash())
	return nil
}


func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeDataToCSV(limits []*big.Int, timings []int64, gasUsed []uint64, perOpcode []float64, opcode string) error {

	file, err := os.OpenFile("../../../findings/data_nat.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"Limits", "Timings [μs]", "GasUsed", "PerOpcode [μs]", opcode}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	numProbes := len(limits)
	for i := 0; i < numProbes; i++ {
		data := []string{
			limits[i].String(),
			strconv.FormatInt(timings[i], 10),
			strconv.FormatUint(gasUsed[i], 10),
			strconv.FormatFloat(perOpcode[i], 'f', 2, 64),
		}
		err := writer.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}