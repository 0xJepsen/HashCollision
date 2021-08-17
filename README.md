# HashCollision

Probabilistically exploring hash collisions in vanity wallets on Eth

I wrote a program to generate random private keys and the corresponding address from the public keys. 

The program then checks the wallet balance of the generate key pair through the Etherscan API.

Th key space is 2^160. The Etherscan API only excepts 5 querries per second with a maximum of three free API keys per account. This means a single account can querry 15 keys per second which is about 1,296,000 a day.
This barelly puts a dent in the keyspace as 1,296,000/2^160 is around 0.0000000000000000000000000000000000000000008 of the total keyspace.

As of August 2021 there are a little under 1,000,000 active ethereum wallets with Eth or Eth based tokens in them. Some of these may be smart contracts, some of these may be lost wallets. 1,000,000/(2^160) is about 0.0000000000000000000000000000000000000000006. Multiplied by the amount of searches a day 1,296,000 is about 0.00000000000000000000000000000000000008 which is about 10 magnitudes grater than searching the whole keyspace.

## Potential Optimizations

While searching the entire keyspace would be a fruitless endeavor there are four optimizations to prespose for this research to be worthwhile. There are two domain abstractions and two computational optimizations that worth while

## Computational Optimizations

Currently the computatational restriction is the use of Etherscans API. To illimenate the requirement of this API one could run a full ethereum node locally. While this will require some harware it will allow us to querry the entire blockcahin data locally rather than through the api. This will also allow multiple instances of the searcher program to run and check against the local Blockchain data. This is the first computational optimization to be looked into. 

The Second computational optimization is in regards to programatic optimization. It is known that python, an interpreted language is not as fast as a precompiled language low level language like C, Rust, or Go. Utilizing these hire languages and proccess optimization, large magnitudes of improvement could be made. 

## Domain Abstraction Optimizations

There is an interesting phenomena is that wallets with leading 0s actually have a lower transaction fee to the network. In short this is because an address with leading zeros takes up less memory on the distributed network. As a result there is a high demand for these "Vanity wallets" with leading zeros. If we can restrict the searched domain for wallets with leading zeross in their addresses this could potentially reduce the search space significantly. 

The second domain optimization is regarding the phenomena of two private keys sharing an address. I still need to look further into this. This could theoretically reduce the search space by a significant ammount. 


## Resources

- [Efficient Ethereum Addresses](https://medium.com/coinmonks/on-efficient-ethereum-addresses-3fef0596e263)
- [Vanity Wallet Generator JavaScript](https://github.com/MyEtherWallet/VanityEth)
- [Two private keys --> one address](https://ethereum.stackexchange.com/questions/10055/is-each-ethereum-address-shared-by-theoretically-2-96-private-keys)
- [Week random source for ECDSA](https://web.archive.org/web/20160308014317/http://www.nilsschneider.net/2013/01/28/recovering-bitcoin-private-keys.html)
- [Python Resource](https://www.arthurkoziel.com/generating-ethereum-addresses-in-python/)
