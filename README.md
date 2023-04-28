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
cp build/TtToken.* contracts/TtToken/
abigen --abi=./build/TtToken.abi --bin=./build/TtToken.bin --pkg=TtToken --out=./contracts/TtToken/TtToken.go

solcjs --base-path node_modules --optimize --abi contracts/AirDropper.sol -o build
solcjs --base-path node_modules --optimize --bin contracts/AirDropper.sol -o build
mv ./build/*AirDropper.abi ./build/AirDropper.abi
mv ./build/*AirDropper.bin ./build/AirDropper.bin
cp build/AirDropper.* contracts/AirDropper/
abigen --abi=./build/AirDropper.abi --bin=./build/AirDropper.bin --pkg=AirDropper --out=./contracts/AirDropper/AirDropper.go
```

# build airdrop
```
go build -o airdropErc20 main.go
```

# modify config files
```
cp env_template .env
vim .env
```

# deploy contract
token
```
./airdropErc20 deploy token
```

airdrop
```
./airdropErc20 deploy airdrop
```

# airdrop

add csv file
```
cp airdrop.csv_template airdrop.csv
vim airdrop.csv
```

execute the airdrop
```
./airdropErc20 airdrop
```

# other features
get eth balance
```
./airdropErc20 getAddrBalance <address>
```

get token balance
```
./airdropErc20 getTokenBalance <address>
```