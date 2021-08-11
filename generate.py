from eth_account import Account
import secrets
from etherscan import Etherscan
eth = Etherscan("API")
import json
import pprint
import time

count = 0;
while True:
    priv = secrets.token_hex(32)
    private_key = "0x" + priv
    acct = Account.from_key(private_key)
    if (count%5==0):
        time.sleep(1)
    balance = eth.get_eth_balance(acct.address)
    print("balance of wallet {} is {}".format(acct.address, balance))
    if (balance != '0'):
        print("FOUND ETH IS WALLET with pub key: {}".format(acct.address))
        break
    count +=1 

print("checked: {} random addresses".format(count))
