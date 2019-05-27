test:
	cd ./smart-contracts && truffle test
	cd ./server && go test -v -timeout 5s
