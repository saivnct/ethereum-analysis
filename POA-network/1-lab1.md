# **Running a Proof of authority network**

## **Lab 1 - basic nodes**

- Clique is a proof-of-authority system where new blocks can be created by authorized ‘signers’ only. The initial set of authorized signers is configured in the genesis block. Signers can be authorized and de-authorized using a voting mechanism, thus allowing the set of signers to change while the blockchain operates

- Signing blocks in Clique networks classically uses the "unlock" feature of Geth so that each node is always ready to sign without requiring a user to manually provide authorization

- In this lab we're going to run:
  - 1 bootnode
  - 1 full node as validator
  - 1 full node

## **Steps**

**1. Clone go-ethereum**

> git clone https://github.com/ethereum/go-ethereum.git

**2. Build geth**

> make all

**3. Create folders for storing chain's data**

> mkdir node1 node2

**4. Create accounts**

- Create account for node 1 => _this is validator_

  > geth --datadir node1 account new
  > <br> => 0xF26fD23524A16249Bd7Cfb25376611EEb5143024

- Create account for node 2

  > geth --datadir node2 account new
  > <br> => 0xad02DC5fC5B958658833880b65eCeB04Cf0A1Ac8

- Note: Save account's password in password.txt

**5. Create genesis.json**

- set node 1 as validator
- set gaslimit 30M
- premined some eth (10^18)

                    {
                         "config": {
                              "chainId": 123451111,
                              "homesteadBlock": 0,
                              "eip150Block": 0,
                              "eip155Block": 0,
                              "eip158Block": 0,
                              "byzantiumBlock": 0,
                              "constantinopleBlock": 0,
                              "petersburgBlock": 0,
                              "istanbulBlock": 0,
                              "muirGlacierBlock": 0,
                              "berlinBlock": 0,
                              "londonBlock": 0,
                              "arrowGlacierBlock": 0,
                              "grayGlacierBlock": 0,
                              "clique": {
                                   "period": 5,
                                   "epoch": 30000
                              }
                         },
                         "difficulty": "1",
                         "gasLimit": "30000000",
                         "extradata": "0x0000000000000000000000000000000000000000000000000000000000000000F26fD23524A16249Bd7Cfb25376611EEb51430240000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
                         "alloc": {
                              "eD19d9CE7392A7f2262a67817E0DBAF68dF042b8": {
                                   "balance": "1000000000000000000000000"
                              },
                              "146142AFF24195Aea62C84c88017e543c7342911": {
                                   "balance": "1000000000000000000000000"
                              },
                              "F26fD23524A16249Bd7Cfb25376611EEb5143024": {
                                   "balance": "500000000000000000000"
                              }
                         }
                    }

**6. Initializing the Geth Database**

- To create a blockchain node that uses this genesis block, first use **geth init** to import and sets the canonical genesis block for the new chain. This requires the path to **genesis.json** to be passed as an argument.

  > geth init --datadir node1 genesis.json
  >
  > geth init --datadir node2 genesis.json

**7. Configure a bootnode**

- Create bootnode key => This key can then be used to generate a bootnode.
  > bootnode -genkey bootnode/boot.key
- Start bootnode, open new terminal:

  > bootnode -nodekey bootnode/boot.key -verbosity 9 -addr :36303
  >
  > => _enode://22d52bd5bf605f94b8cade150ad991ddfb0974a1d40fa7aceff8b6fd8aebb3b8342b18a284accb441a940135d41aa0e02d536180a2e24f2311a62fb585bcf5d5@127.0.0.1:0?discport=36303_

- **Note**: The choice of port passed to -addr is arbitrary, but public Ethereum networks use **30303**, so this is best avoided

**8. Start eth nodes**

- Open 2 terminals
- Run node 1 as miner:

  > geth --datadir node1 --port 30306 --bootnodes "enode://22d52bd5bf605f94b8cade150ad991ddfb0974a1d40fa7aceff8b6fd8aebb3b8342b18a284accb441a940135d41aa0e02d536180a2e24f2311a62fb585bcf5d5@127.0.0.1:0?discport=36303" --networkid 123451111 --unlock 0xF26fD23524A16249Bd7Cfb25376611EEb5143024 --password password.txt --authrpc.port 8551 **--miner.etherbase** 0xF26fD23524A16249Bd7Cfb25376611EEb5143024 **--mine**

- Run node 2:
  > geth --datadir node2 --port 30308 --bootnodes "enode://22d52bd5bf605f94b8cade150ad991ddfb0974a1d40fa7aceff8b6fd8aebb3b8342b18a284accb441a940135d41aa0e02d536180a2e24f2311a62fb585bcf5d5@127.0.0.1:0?discport=36303" --networkid 123451111 --unlock 0xad02DC5fC5B958658833880b65eCeB04Cf0A1Ac8 --password password.txt --authrpc.port 8552
- **Note**:
  - Node 1: We must set _--miner.etherbase_ & _--mine_ in order to make node 1 mine & validate block
  - Every node must be run on dedicated ports
  - If we try to run node 2 as miner => still run, but get error "Block sealing failed err="unauthorized signer" => it cannot seal block

**9. Testing**

- Connect to node 1
  > geth attach node1/geth.ipc
  >
  > net.peerCount
  >
  > admin.peers
  >
  > eth.getBalance(eth.accounts[0])
  >
  > eth.blockNumber
  >
  > web3.fromWei(eth.getBalance(eth.accounts[0]), 'ether');
  > web3.fromWei(eth.getBalance('0xad02DC5fC5B958658833880b65eCeB04Cf0A1Ac8'), 'ether');
  >
  > web3.eth.getGasPrice(function(error, result){
  > console.log(result);
  > });
  >
  > eth.sendTransaction({
  > to: '0xad02DC5fC5B958658833880b65eCeB04Cf0A1Ac8',
  > from: eth.accounts[0],
  > value: web3.toWei(1, 'ether')
  > });
