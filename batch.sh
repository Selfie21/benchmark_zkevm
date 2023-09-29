for f in smartcontracts/*;
do
echo "################# NEW BENCHMARK OPCODE ######################## : " $f;
rm processing/Analysis.sol;
cp $f processing/Analysis.sol;
solc --abi processing/Analysis.sol -o processing/ --overwrite;
solc --bin processing/Analysis.sol -o processing/ --overwrite;
abigen --abi processing/Analysis.abi --pkg Analysis --type Analysis --bin processing/Analysis.bin --out processing/Analysis.go;
cp processing/Analysis.go go-ethereum/tests/testdata/Analysis/;
cp processing/Analysis.go zkevm-node/test/testdata/Analysis/;
cd zkevm-node/test/testdata/;
go run main.go $f;
cd ../../../;
done