# Hash Collisions

Probabilistically exploring hash collisions in vanity wallets on Ethereum.

I wrote a program to generate random Private keys, the corresponding public keys, and the hashed wallet addresses. 

The program then checks the wallet balance of the generated key pair through the Etherscan API.

The keyspace is 2^160. The Etherscan API only accepts five queries per second with a maximum of three free API keys per account. A single account can query 15 keys per second and about 1,296,000 a day.
This barely puts a dent in the keyspace as 1,296,000/2^160 is around 0.0000000000000000000000000000000000000000008 of the total keyspace.

As of August 2021, there are a little under 1,000,000 active Ethereum wallets with Eth or Eth-based tokens in them. Some of these may be smart contracts, some of these may be lost wallets. 1,000,000/(2^160) is about 0.0000000000000000000000000000000000000000006. Multiplied by the number of searches a day 1,296,000 is about 0.00000000000000000000000000000000000008, which is about ten magnitudes greater than searching the whole keyspace.

## Potential Optimizations

While searching the entire keyspace would be a fruitless endeavor, there are four optimizations to propose for this research to be worthwhile. There are two domain abstractions and two computational optimizations that worthwhile

## Computational Optimizations

Currently, the computational restriction is the use of Etherscans API. To eliminate the requirement of this API, one could run a full Ethereum node locally. While this will require some hardware, it will allow us to query the entire blockchain data locally rather than through the API. This will also allow multiple instances of the searcher program to run and check against the local Blockchain data. 

The Second computational optimization is in regards to programmatic optimization. Python, an interpreted language, is not as fast as a precompiled language low-level language like C, Rust, or Go. Utilizing these languages and process optimization, large magnitudes of improvement could be made. 

## Domain Abstraction Optimizations

Wallets with leading 0s produce lower transaction fees for related transactions. In short, this is because an address with leading zeros takes up less memory on the distributed network. As a result, there is a high demand for these "Vanity wallets" with leading zeros. If we can restrict the searched domain for wallets with leading zeros in their addresses, this could significantly reduce the search space. 

The second domain optimization is regarding the phenomena of two private keys sharing an address. I still need to look further into this. This could theoretically reduce the search space by a significant amount. 

## Resources

- [Efficient Ethereum Addresses](https://medium.com/coinmonks/on-efficient-ethereum-addresses-3fef0596e263)
- [Pigon Hole Principle proof of two private keys to one addresss](https://crypto.stackexchange.com/questions/72741/what-is-the-possibility-of-collision-of-trailing-160-bits-of-keccak-256-for-any/72753#72753)
- [Vanity Wallet Generator JavaScript](https://github.com/MyEtherWallet/VanityEth)
- [Two private keys --> one address](https://ethereum.stackexchange.com/questions/10055/is-each-ethereum-address-shared-by-theoretically-2-96-private-keys)
- [Week random source for ECDSA](https://web.archive.org/web/20160308014317/http://www.nilsschneider.net/2013/01/28/recovering-bitcoin-private-keys.html)
- [Python Resource](https://www.arthurkoziel.com/generating-ethereum-addresses-in-python/)
