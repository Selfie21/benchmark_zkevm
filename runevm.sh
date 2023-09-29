rm processing/Analysis.sol;
cp smartcontracts/Analysis.sol processing/Analysis.sol;
solc --abi processing/Analysis.sol -o processing/ --overwrite
solc --bin processing/Analysis.sol -o processing/ --overwrite
abigen --abi processing/Analysis.abi --pkg Analysis --type Analysis --bin processing/Analysis.bin --out processing/Analysis.go
cp processing/Analysis.go go-ethereum/tests/testdata/Analysis/
cd go-ethereum/tests/testdata/
go run main.go Analysis.sol
