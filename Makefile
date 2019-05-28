all: 
	solc --abi smart-contracts/contracts/TwitterWallet.sol -o smart-contracts/build --allow-paths . --overwrite
	abigen --abi=smart-contracts/build/TwitterWallet.abi --pkg=wallet --out=server/wallet/TwitterWallet.go

test:
	cd ./smart-contracts && truffle test
	cd ./server && go test -v -timeout 5s
