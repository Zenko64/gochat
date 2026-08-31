# Gochat
---
Gochat is a simple network utility that allows chatting between computers in a network by broadcasting UDP packages.
Each message has a max size of 1024 characters, and each chatter is identified by the IP Address.

How to use:
Open the binary in multiple computers in the same network, then send a message. all the computers will receive this message.

Tested on Windows.


## Compiling
> Make sure you have the latest version of Golang installed.  

```bat
git clone https://github.com/Zenko64/gochat.git
cd gochat
```

To run:
```bat
go run .
```

To build:
```bat
go build .
```
