# erc20-airdrop

erc20 airdrop tools

## install build tools

```shell
npm install -g solc
solcjs --help
```

install abigen

```shell
# https://geth.ethereum.org/downloads ## download last (Geth & Tools)
wget https://gethstore.blob.core.windows.net/builds/geth-alltools-darwin-amd64-1.11.5-a38f4108.tar.gz
tar zxvf geth-alltools-darwin-amd64-1.11.5-a38f4108.tar.gz && cd geth-alltools-darwin-amd64*
sudo mv abigen /usr/local/bin
```

## from sol to build go

```shell
npm install @openzeppelin/contracts

solcjs --base-path node_modules --optimize --abi contracts/TtToken.sol -o build
solcjs --base-path node_modules --optimize --bin contracts/TtToken.sol -o build
mv ./build/*TtToken.abi ./build/TtToken.abi
mv ./build/*TtToken.bin ./build/TtToken.bin
cp build/TtToken.* contracts/TtToken/
abigen --abi=./build/TtToken.abi --bin=./build/TtToken.bin --pkg=TtToken --out=./contracts/TtToken/TtToken.go

solcjs --base-path node_modules --optimize --abi contracts/AirDropper.sol -o build
solcjs --base-path node_modules --optimize --bin contracts/AirDropper.sol -o build
mv ./build/*AirDropper.abi ./build/AirDropper.abi
mv ./build/*AirDropper.bin ./build/AirDropper.bin
cp build/AirDropper.* contracts/AirDropper/
abigen --abi=./build/AirDropper.abi --bin=./build/AirDropper.bin --pkg=AirDropper --out=./contracts/AirDropper/AirDropper.go
```

## build airdrop

```shell
go build -o airdropErc20 main.go
```

## modify config files

```shell
cp env_template .env
vim .env
```

## deploy contract

token

```shell
./airdropErc20 deploy token
```

airdrop

```shell
./airdropErc20 deploy airdrop
```

## airdrop

add csv file

```shell
cp airdrop.csv_template airdrop.csv
vim airdrop.csv
```

execute the airdrop

```shell
./airdropErc20 airdrop
```

## other features

get eth balance

```shell
./airdropErc20 getAddrBalance <address>
```

get token balance

```shell
./airdropErc20 getTokenBalance <address>
```
