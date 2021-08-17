# HashCollision

Probabilistically exploring hash collisions in vanity wallets on Eth

I wrote a program to generate random private keys and the corresponding address from the public keys. 

The program then checks the wallet balance of the generate key pair through the Etherscan API.

Th key space is 2^160. The Etherscan API only excepts 5 querries per second with a maximum of three free API keys per account. This means a single account can querry 15 keys per second which is about 1,296,000 a day.
This barelly puts a dent in the keyspace as 1,296,000/2^160 is around 0.0000000000000000000000000000000000000000008, the probability that a hash collission will be found.

As of August 2021 there are a little under 1,000,000 active ethereum wallets with Eth or Eth based tokens in them. Some of these may be smart contracts, some of these may be lost wallets.

There is an interesting phenomena is that wallets with leading 0s actually have a lower transaction gas fee.


## Resources

- [Two private keys --> one address](https://ethereum.stackexchange.com/questions/10055/is-each-ethereum-address-shared-by-theoretically-2-96-private-keys)
- [Week random source for ECDSA](https://web.archive.org/web/20160308014317/http://www.nilsschneider.net/2013/01/28/recovering-bitcoin-private-keys.html)
- [Python Resource](https://www.arthurkoziel.com/generating-ethereum-addresses-in-python/)
