# root-unix-system
Go script to root yourself on the system

### Requirement
A unix-like system. Won't work on windows. Works on macOS/Linux distro's

# Installation 
Install Go ``minimum Go version: 1.19``

Navigate to ``https://go.dev/dl/`` install the one you need compabitle with your OS.

```
# Installation
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz 

// you may need to run the above command as sudo, if you do change alse the $HOME/.profile of root with source $HOME/.profile

export PATH=$PATH:/usr/local/go/bin
source $HOME/.profile

go version // if you get output it works. Do not forget to follow the last 2 steps on differents users on your machine
```

# Usage
One-time run:
```go run .```

After that run one-time:
``go build`` to compile everything so you can have an executable file.

Now your installation is done, just run the executable and there u go
