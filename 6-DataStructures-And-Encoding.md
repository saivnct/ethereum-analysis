# **Data Structure And Encoding**

## **1. Merkle Patricia Tries (MPT)**

https://medium.com/codechain/modified-merkle-patricia-trie-how-ethereum-saves-a-state-e6d7555078dd

https://www.youtube.com/watch?v=QlawpoK4g5A

### **a. 4 different merkles Tries**

- **Transaction Trie**

  - its purpose is to record transactions.
  - Since the ordering of the data is mostly decided upon by the miner, we do not know what the data looks like until it is mined.
  - We do know what parameters are used to compose the transaction trie, and they’re as follows:
    - Account nonce
    - Gas price
    - Gas limit
    - Recipient
    - Transfer value
    - Transaction signature values
    - Account initialization (if transaction is of contract creation type), or transaction data (if transaction is a message call)

  NOTE: Transaction Trie is specific to a particular block -> Once the block is mined, the transaction trie is never updated

- **Receipt Trie**

  - its purpose it to record the outcome of a transaction.
  - According to the yellow paper, the receipt trie can also be useful for zero-knowledge proofs or searches.
  - The parameters that make up the receipt trie are as follows:
    - Post-transaction state
    - Cumulative gas used
    - Logs
    - Bloom filter created from the information of the above logs

  **NOTE**: Receipt Trie is specific to a particular block -> Once the block is mined, the Receipt trie is never updated

- **State Trie**

  - The one and only one global state trie.
  - It contains a key-value pair for every Ethereum account on the network
    - key is an ethereum address
    - value is RLP encoded of account's details. An ethereum account and the state trie is comprised of the following fields:
      - Nonce
      - Balance
      - Storage Root
      - Codehash
  - Unlike the transaction and receipt tries, the state trie updates over time….constantly.

- **Storage Trie** (also known as the Contract Storage Trie)
  - is a data structure that represents the storage of an individual smart contract.
  - It stores the key-value pairs that define the persistent state of a contract
  - Each smart contract on Ethereum has its own Storage Trie, and it is specific to that contract
  - When a smart contract is deployed or executed, its storage state is stored and managed within the Storage Trie.
  - Unlike the transaction and receipt tries, the Storage Trie is not specific to a particular block, the state trie updates over time….constantly.

https://medium.com/@eiki1212/ethereum-state-trie-architecture-explained-a30237009d4e
https://medium.com/cybermiles/diving-into-ethereums-world-state-c893102030ed

- -> The root for each of the tries is a Keccak 256-bit hash
- -> The Storage Tree’s root lives within the RLP encoded data value within the State Trie

### **b. Merkles roots nodes included in block's header:**

- state Root
- transaction Root
- receipts Root

### **c. Use cases:**

- Has this transaction been included in a particular block? -> Transaction Tree
- Tell me all instances of an event of type X (eg. a crowdfunding contract reaching its goal) emitted by this address in the past 30 days -> Receipt tree.
- What is the current balance of my account? -> State Tree
- Does this account exist? -> State Tree
- Pretend to run this transaction on this contract. What would the output be? -> State Tree

---

## **2. Tries Architecture**

<div class="image-container" align="center">
<img src="img/eth-block-arch-1.jpeg" alt="Image 1"  >
</div>

---

## **3. Merkle Patricia Tries (MPT)**

### **a. Node type:**

- empty nodes
  - a blank node
- branch nodes
  - any node with 1 or more branches
  - branch node can have leaf node and extension node in branch
- extension nodes
  - branch node with 1 child is compressed
  - store key-value
- leaf nodes

  - node that doesnt have child
  - store key-value

<div class="image-container" align="center">
<img src="img/Merkle-Patricia-Trie-nodes-types.png" alt="Image 1"  >
</div>

- Nibble:
  - leaf & extension nodes are same as both contain key-value pair -> how we recognize whether the node is leaf or extension => Nibble help us there

<div class="image-container" align="center">
<img src="img/Merkle-Patricia-Trie-nodes-types-example.png" alt="Image 1"  >
</div>
