# erc20-airdrop
erc20 airdrop tools

# install build tools
```
npm install -g solc
solcjs --help
```

install abigen
```
# https://geth.ethereum.org/downloads ## download last (Geth & Tools)
wget https://gethstore.blob.core.windows.net/builds/geth-alltools-darwin-amd64-1.11.5-a38f4108.tar.gz
tar zxvf geth-alltools-darwin-amd64-1.11.5-a38f4108.tar.gz && cd geth-alltools-darwin-amd64*
sudo mv abigen /usr/local/bin
```

# from sol to build go
```
npm install @openzeppelin/contracts

solcjs --base-path node_modules --optimize --abi contracts/TtToken.sol -o build
solcjs --base-path node_modules --optimize --bin contracts/TtToken.sol -o build
mv ./build/*TtToken.abi ./build/TtToken.abi
mv ./build/*TtToken.bin ./build/TtToken.bin
cp build/TtToken.* contracts
abigen --abi=./build/TtToken.abi --bin=./build/TtToken.bin --pkg=contracts --out=./contracts/TtToken.go

solcjs --base-path node_modules --optimize --abi contracts/AirDropper.sol -o build
solcjs --base-path node_modules --optimize --bin contracts/AirDropper.sol -o build
mv ./build/*AirDropper.abi ./build/AirDropper.abi
mv ./build/*AirDropper.bin ./build/AirDropper.bin
cp build/AirDropper.* contracts
abigen --abi=./build/AirDropper.abi --bin=./build/AirDropper.bin --pkg=contracts --out=./contracts/AirDropper.go
```

# build airdrop
```
go build -o airdropErc20 main.go
```
