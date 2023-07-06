# Go-Ethereum Analysic

## 1. accounts module

- used for wallet management

## 2. cmd module

- the CLI of Geth
- cmd/geth/main.go
  - -> this file is executed when you start the CLI. This can be a fun to follow the rabbit hole and follow the execution path from there.

## 3. consensus module

- where you have the mining algorithm ethhash

## 4. core module

- manages the data structure of the Blockchains, like blocks, transactions
- Ethereum Virtual Machine or EVM, which executes the smart contract.
  - -> That’s one of the most interesting part of the code

## 5. eth module

- keep the local Blockchain synced with the rest of the network

## 6. ethdb module

- the database where the Ethereum Blockchain is stored, which uses a key value db used by google called LevelDB

## 7. miner module

- for mining Ethereum. It makes use of the ethhash mining algorithm in the consensus module

## 8. p2p module

- low-level networking library to exchange data between nodes

## 9. params module

- where you have various configuration. For example there is the block number for the different forks.

## 10. rpc module

- which handles all the interfaces of geth, such as the json rpc API, the web socket API and the IPC API, a very fast way to interact with a GETH client if you are on the same machine.

## 11. tests module

- integration tests to make sure that the implementation of geth respects the specification of the Ethereum protocol

# NOTE:

    - Solidity is not part of Geth, it’s a separate repo
