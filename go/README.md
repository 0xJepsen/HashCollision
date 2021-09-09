# Hash Collisions in Distributed ledgers

This is a Golang version of 0xJepsen's Python Hash Collision program.

## Running the program

Clone the repo locally:

`git clone https://github.com/0xJepsen/HashCollision`

Enter the 'go' folder:

`cd HashCollision/go`

Create a .env file:

`touch .env`

Add this to your newly created .env file and enter your Etherscan API key

`ETHERSCAN_KEY=<YOUR_API_KEY>`

Run the program:

`go run main.go`