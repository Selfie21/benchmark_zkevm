package main

import (
	"os"
	"fmt"
	"time"
	"context"
	"strconv"
	"math/big"
	"encoding/csv"

	"github.com/0xPolygonHermez/zkevm-node/hex"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/test/operations"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xPolygonHermez/zkevm-node/test/testdata/Analysis"
)

const (
	txTimeout = 60 * time.Second
)

func main() {
	var networks = []struct {
		Name       string
		URL        string
		ChainID    uint64
		PrivateKey string
	}{
		{Name: "Local L1", URL: operations.DefaultL1NetworkURL, ChainID: operations.DefaultL1ChainID, PrivateKey: operations.DefaultSequencerPrivateKey},
		{Name: "Local L2", URL: operations.DefaultL2NetworkURL, ChainID: operations.DefaultL2ChainID, PrivateKey: operations.DefaultSequencerPrivateKey},
	}

	args := os.Args
	opcode := args[1]
	log.Debugf("Opcode %v", opcode)
	limits := []*big.Int{big.NewInt(0), big.NewInt(100), big.NewInt(1000), big.NewInt(2000)}
	numProbes := len(limits)
	var timings = make([]int64, numProbes)
	var gasUsed = make([]uint64, numProbes)
	var perOpcode = make([]float64, numProbes)

	ctx := context.Background()
	polygon := networks[1]
	log.Infof("connecting to %v: %v", networks[0].Name, networks[0].URL)
	log.Infof("connecting to %v: %v", polygon.Name, polygon.URL)
	client, err := ethclient.Dial(polygon.URL)
	chkErr(err)
	log.Infof("connected")

	auth := operations.MustGetAuth(polygon.PrivateKey, polygon.ChainID)
	chkErr(err)

	balance, err := client.BalanceAt(ctx, auth.From, nil)
	chkErr(err)
	log.Debugf("ETH Balance for %v: %v", auth.From, balance)

	log.Debugf("Sending TX to deploy Analysis SC")
	_, tx, analysisSC, err := Analysis.DeployAnalysis(auth, client)
	chkErr(err)
	err = operations.WaitTxToBeMined(ctx, client, tx, txTimeout)
	chkErr(err)

	for i, limit := range limits {
		log.Debugf("Calling Benchmark method from Analysis with limit %v", limit)
		start := time.Now()
		tx, err = analysisSC.Benchmark(auth, limit)
		chkErr(err)
		err = operations.WaitTxToBeMined(ctx, client, tx, txTimeout)
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

func ethTransfer(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts, to common.Address, amount *big.Int, nonce *uint64) *types.Transaction {
	if nonce == nil {
		log.Infof("reading nonce for account: %v", auth.From.Hex())
		var err error
		n, err := client.NonceAt(ctx, auth.From, nil)
		log.Infof("nonce: %v", n)
		chkErr(err)
		nonce = &n
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	chkErr(err)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{To: &to})
	chkErr(err)

	tx := types.NewTransaction(*nonce, to, amount, gasLimit, gasPrice, nil)

	signedTx, err := auth.Signer(auth.From, tx)
	chkErr(err)

	log.Infof("sending transfer tx")
	err = client.SendTransaction(ctx, signedTx)
	chkErr(err)
	log.Infof("tx sent: %v", signedTx.Hash().Hex())

	rlp, err := signedTx.MarshalBinary()
	chkErr(err)

	log.Infof("tx rlp: %v", hex.EncodeToHex(rlp))

	return signedTx
}

func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func writeDataToCSV(limits []*big.Int, timings []int64, gasUsed []uint64, perOpcode []float64, opcode string) error {

	file, err := os.OpenFile("../../../findings/data_zk.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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