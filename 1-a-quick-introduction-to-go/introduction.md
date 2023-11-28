# 1 | A Quick Introduction to Go

### Install GO on Linux
- **Download newest go version**
    ```bash
    wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
    ```
- **Remove and extract**
    ```bash
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz && rm -rf go1.21.4.linux-amd64.tar.gz
    ```
- **Set permanent path**
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc && source ~/.bashrc
    ```
- **Verify version**
    ```bash
    go version
    ```

### Setup VScode for GO
1. Extensions > Go (Go team at Google)
2. Ctrl + Shift + P > GO: install tools > Check all and OK


### The go doc and godoc utilities
- **Install Go Doc**
    ```bash
    go install golang.org/x/tools/cmd/godoc@latest
    ```
- **Find information**
    ```bash
    go doc fmt.Printf
    go doc fmt

- **Set permanent path**
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc && source ~/.bashrc
    ```
    
- **Go Doc Server UI**
    ```bash
    godoc -http=:8001
    ```

