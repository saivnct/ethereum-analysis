# **Accounts - Transaction & Message**

## **Accounts**

### **1. Externally owned accounts (EOAs) - An externally controlled account**

- has an ether balance,
- can send transactions (ether transfer or trigger contract code),
- is controlled by private keys,
- has no associated code.

### **2. Contract accounts - A contract**

- has an ether balance,
- has associated code,
- code execution is triggered by transactions or messages (calls) received from other contracts.
- when executed - perform operations of arbitrary complexity (Turing completeness) - manipulate its own persistent storage, i.e. can have its own permanent state - can call other contracts

---

## **Transaction & Message**

- All action on the Ethereum block chain is set in motion by transactions fired from accounts.
- Every time a contract account receives a transaction, its code is executed as instructed by the input parameters sent as part of the transaction.
- The contract code is executed by the Ethereum Virtual Machine on each node participating in the network as part of their verification of new blocks.

### **1. Transactions**

- The term “transaction” is used in Ethereum to refer to the signed data package that stores a message to be sent from an externally owned account to another account on the blockchain.
- Transactions contain:

  - The **recipient** of the message

  - A **signature identifying the sender** and proving their intention to send the message via the blockchain to the recipient

  - **VALUE** field - The amount of wei to transfer from the sender to the recipient

  - An optional **data** field, which can contain the message sent to a contract

  - **GASLIMIT** value, Gas limit refers to the maximum amount of gas you are willing to consume on a transaction.

    - More complicated transactions involving smart contracts require more computational work, so they require a higher gas limit than a simple payment.

    - A standard ETH transfer requires a gas limit of 21,000 units of gas.

    - For example, if you put a gas limit of 50,000 for a simple ETH transfer, the EVM would consume 21,000, and you would get back the remaining 29,000. However, if you specify too little gas, for example, a gas limit of 20,000 for a simple ETH transfer, the EVM will consume your 20,000 gas units attempting to fulfill the transaction, but it will not complete. The EVM then reverts any changes, but since the miner has already done 20k gas units worth of work, that gas is consumed.

  - **GASPRICE** value, representing the fee the sender is willing to pay for gas. One unit of gas corresponds to the execution of one atomic instruction, i.e. a computational step.

### **2. Messages**

- Contracts have the ability to send “messages” to other contracts. Messages are virtual objects that are never serialized and exist only in the Ethereum execution environment. They can be conceived of as function calls.
- A message is like a transaction, except it is produced by a contract and not an external actor
- A message is produced when a contract currently executing code executes the **'CALL'** or **'DELEGATECALL'** opcodes, which produces and executes a message.
- Messages are also sometimes called **“internal transactions.”**
- Like a transaction, a message leads to the recipient account running its code. Thus, contracts can have relationships with other contracts in exactly the same way that external actors can

- A message contains:
  - The **sender** of the message (implicit).
  - The **recipient** of the message
  - **VALUE** field - The amount of wei to transfer alongside the message to the contract address,
  - An optional **data** field, that is the actual input data to the contract
  - **GASLIMIT** value, which limits the maximum amount of gas the code execution triggered by the message can incur.
