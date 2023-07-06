# **Block Structure**

## **1. ETHEREUM 1**

### **a. Header**

- Parent Block’s hash
- Uncle Hash

  - An Uncle Hash is a reference to the hash of a block that is not included in the main blockchain but is still considered valid.
  - In Ethereum, when a miner finds a new block, other miners may also be working on finding a new block at the same time. If two miners find a new block at the same time, the one whose block gets added to the main blockchain first is called the “main block”, while the other is called an “uncle block”.

  - The Uncle Hash is included in the header of the main block and is a reference to the hash of the uncle block.
  - Uncle blocks are rewarded with a smaller amount of Ether than main blocks, as a way to incentivize miners to continue to mine even if their blocks do not make it into the main blockchain.
  - By including the Uncle Hash in the block header, Ethereum allows for the recognition of work done by miners even if the block is not included in the main chain, and also helps to promote network security.

- State Root Hash
  - Root node of world state.
- Transaction Hash Root
  - Root node of transaction hash.
- Receipt Root Hash
  - Root node of receipt hash.
- Logs bloom

  - The logs bloom is a filter that is included in the header of each block in the Ethereum blockchain.
  - It is used to efficiently check if a log event from a contract execution is included in the block.
  - A log event is a record of an event that occurred during the execution of a smart contract, such as a transfer of funds or a change in the state of the contract.

  - The logs bloom is a bit array that is constructed by taking the hashes of the log topics and setting the corresponding bits in the array.
  - The log topics are the indexed data, such as the address of the contract that emitted the event and the event signature.
    - => The bloom filter allows for quick checking of whether a log event is included in a block, without the need to scan through all the logs in the block.
    - => By including the logs bloom in the block header, Ethereum allows for efficient searching of logs events without having to scan through all the logs in the block. This makes it easy to look up events and improves the performance of searching.

- Timestamp
  - Unix timestamp that represents the time when mining ends.
- Difficulty
  - Difficulty is a value that shows how difficult to find a hash.
- Nonce
  - A nonce is a number that can only be used once. In cryptography, a nonce is a one-time code chosen randomly to transmit password securely and prevent replay attacks.
- Gas Limit

  - Gas limit set for the block.
  - Block gas limit is a maximum amount of gas allowed in a block. It determines how many transactions can be stored in a block based on the sum of gas.

    > Ex:

        - block gas limit = 100
        - we have transactions which gas limits are 50, 50, and 10.
        	=> Block can store only the first two transactions(the 50s) but not the last one(10).
        	=> When miner tries to append the last transaction, Ethereum emits “Transaction exceeds block gas limit” error.

  - The gas limit is used to prevent a situation where the network becomes overwhelmed by too many transactions, causing delays and increased fees.

    - => By including the gas limit in the block header, Ethereum allows for a mechanism to control the rate at which new transactions are processed and to ensure that the network remains secure and efficient.

  - The gas limit can be adjusted by the miner who mines the block, and it can be changed based on the current state of the network. The Ethereum protocol also has a mechanism called block gas limit, which is a dynamic algorithm that adjusts the gas limit based on the gas usage of recent blocks.

- Gas Used

  - Sum of all the gas used by all transaction in the block.
    - => By including the gas used in the block header, Ethereum allows for a mechanism to track the total amount of gas consumed by the transactions in a block and to ensure that the network remains secure and efficient.
    - => This information can be used to monitor network usage and to make decisions about adjustments to the gas limit or other network parameters.

- Extra Data

  - An optional and free field to store 32-byte extra data.
    - => The extra data field is not used by the Ethereum protocol for any specific purpose and is intended for use by miners or other users of the network. It can be used to include a message, signature, or other data that may be useful for the miner or other users of the network.

- number - Counting number of the block. The number increments sequentially. 0 is a genesis block.

### **b. Body(block payload)**

- The main components of the block body are the list of transactions and the list of uncles (stale blocks).

- The list of transactions in the block body contains all the transactions that were included in the block. Each transaction includes information such as the sender’s address, the recipient’s address, the amount of Ether to be transferred, and the amount of gas to be consumed.

- The list of uncles in the block body contains all the stale blocks that were included in the block. These stale blocks are included in the block as a reward for the miner who mined them, even though they were not included in the main blockchain.

- Block Reward: A field that contains the reward given to the miner who mined the block. This reward is a combination of the block reward, the uncle reward, and the transaction fee rewards.

### **c. Block creation diagram**

<div class="image-container" align="center">
<img src="img/eth-block-arch-1.jpeg" alt="Image 1"  >
</div>

### **d. Block**

<div class="image-container" align="center">
<img src="img/eth-block-arch-2.jpeg" alt="Image 1"  >
</div>

### **e. Execution Workflow**

<div class="image-container" align="center">
<img src="img/eth-block-arch-3-2003.png" alt="Image 1"  >
</div>
