#Very simple P2P chat application

An example of a simple peer to peer chat application which uses
the net/http package to make Get, Post HTTP requests as the interprocess
communication procedure to broadcast messages between peers.

## Getting started

Currently this project has no dependencies on which it is reliant, so it can built
using `go`:

Recommend os : Ubuntu 16.04.2 LTS

Install GOlang 1.7

Login to the Ubuntu system and upgrade to apply latest security updates.

$ sudo apt-get update
$ sudo apt-get -y upgrade

Download golang:

$ wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz

Extract and install golang:

$ sudo tar -xvf go1.7.4.linux-amd64.tar.gz
$ sudo mv go /usr/local

configure the package GOROOT directory:

$ export GOROOT=/usr/local/go

configure the project GOROOT directory:

$ export GOPATH=$HOME/Projects

set path:

$ export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

Build Peer P2P Application

```
$ git clone https://github.com/labell6/p2p.git
$ cd p2p
$ go build -o peer
$ mv peer ../bin (optional)
```

## Starting a Peer node

To start a node for this network just run the binary:

```
$ ./peer
```

You can start a node with the following options:
```
$ ./peer -n testname

To specify the onscreen name broadcast to other peers

$ ./peer -p 8999 

To specify the network port used.

$ ./peer -i  159.8.181.148

To specify the ip address of a peer
```

Once two peers have started on different nodes. Text entered into stdin of one peer, will be broadcast to other peers.
