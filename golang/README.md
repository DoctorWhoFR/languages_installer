# Golang installer

There is 3 steps to install last golang version on your server.

### 1 - Get last golang version path

Go on : https://go.dev/dl/

Get the last linux link, like : https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
The path should be : go1.20.4.linux-amd64.tar.gz

### 2 - Execute golang_installer

(You should maybe have to execute the `chmod +x golang_installer` command to use it)

./golang_installer <version>
./golang_installer go1.20.4.linux-amd64.tar.gz

### 3 - Add golang to path and check install

- declare -x PATH=$PATH:/usr/local/go/bin
- go version

You should see something like : 
`go version go1.20.4 linux/amd64` 