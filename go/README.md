# Hash Collisions in Distributed ledgers

This is a Golang version of 0xJepsen's Python Hash Collision program.

## Running the program

`git clone github.com/0xJepsen/HashCollision`
`cd go`
`touch .env`
Add this to your newly created .env file and enter your Etherscan API key
`ETHERSCAN_KEY=<YOUR_API_KEY>`
`go run main.go`