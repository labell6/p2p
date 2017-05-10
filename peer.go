package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
        "strings"
	"os"
)

// Instances of Msg store the transmited message and its source peer data 
type Msg struct {
	Message string
	From    Peer
}

// Peer stores screen name and ip address data
type Peer struct {
	Name    string
	Address string
}

// Peers provides a map of known Peers accessed by an Address
type Peers map[string]Peer

// PeerData implements the p2p control data for a peer
type PeerData struct {
	Self            Peer
	Peers           Peers
	addPeers        chan (Peer)
	exitPeer        chan (Peer)
	currentPeers    chan (Peers)
	getCurrentPeers chan (bool)
	receivedMsg     chan (Msg)
	userMsg         chan (Msg)
}

// Startup client logic
func main() {
        //Parse comandline argumments
	port := flag.String("p", "7500", "Port number")
	name := flag.String("n", "Surematics", "ScreenName")
	peer := flag.String("i", "", "IP address of peer to connect to.")
	flag.Parse()
        
        //initialize client control handler  
	peercontroller := NewPeerData(Peer{*name, getIp() + ":" + *port})
	peercontroller.start()
	
	if len(*peer) != 0 {
		peercontroller.addPeers <- Peer{"", *peer}
	}
        
        //scan stdin for input
	peercontroller.startStdinListener(peercontroller.Self)
}

// Get IP of this machine
func getIp() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        return "localhost"
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().String()
    idx := strings.LastIndex(localAddr, ":")

    return localAddr[0:idx]
}

// NewPeerData instantiate, initialize, and return PeerData control data
func NewPeerData(self Peer) *PeerData {
	peercontroller := new(PeerData)
	peercontroller.Self = self
	peercontroller.Peers = make(Peers)
	peercontroller.addPeers = make(chan (Peer))
	peercontroller.currentPeers = make(chan (Peers))
	peercontroller.getCurrentPeers = make(chan (bool))

	peercontroller.userMsg = make(chan (Msg))
	peercontroller.receivedMsg = make(chan (Msg))
	return peercontroller
}

// start api and command processing threads
func (peercontroller *PeerData) start() {
	go peercontroller.cmdLoop()
	go peercontroller.webListener()
	fmt.Printf("Peer node started on %s \n", peercontroller.Self.Address)
}

//Process detected events
func (peercontroller *PeerData) cmdLoop() {
	for {
		select {
		case peer := <-peercontroller.addPeers:
			if !peercontroller.existingPeer(peer) {
				peercontroller.Peers[peer.Address] = peer
				go peercontroller.sendAdd(peer)
				fmt.Printf("Connecting to: %s \n", peer.Address)
			}

		case <-peercontroller.getCurrentPeers:
			peercontroller.currentPeers <- peercontroller.Peers

		case peer := <-peercontroller.exitPeer:
			delete(peercontroller.Peers, peer.Address)

		case Msg := <-peercontroller.receivedMsg:
			fmt.Printf("%s: %s\n", Msg.From.Name, Msg.Message)

		case Msg := <-peercontroller.userMsg:
			for _, peer := range peercontroller.Peers {
				go peercontroller.sendTx(peer, Msg)
			}
		}
	}
}


func (peercontroller *PeerData) existingPeer(peer Peer) bool {
	if peer.Address == peercontroller.Self.Address {
		return true
	}
	_, existingPeer := peercontroller.Peers[peer.Address]
	return existingPeer
}

//connect to other peer nodes
func (peercontroller *PeerData) sendAdd(peer Peer) {
	URL := "http://" + peer.Address + "/add"

	j, _ := json.Marshal(peercontroller.Self)

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(j))
	if err != nil {
		peercontroller.exitPeer <- peer
		return
	}

	peercontroller.addPeers <- peer

	defer resp.Body.Close()
	otherPeers := Peers{}
	decode := json.NewDecoder(resp.Body)
	err = decode.Decode(&otherPeers)
	for _, peer := range otherPeers {
		peercontroller.addPeers <- peer
	}
}

//send message to peer
func (peercontroller *PeerData) sendTx(peer Peer, msg Msg) {
	URL := "http://" + peer.Address + "/tx"

	j, _ := json.Marshal(msg)

	_, err := http.Post(URL, "application/json", bytes.NewBuffer(j))
	
	if err != nil {
		peercontroller.exitPeer <- peer
		return
	}
}

//start api handlers to allow interaction with peers
func (peercontroller *PeerData) webListener() {
	http.HandleFunc("/tx", createTxHandler(peercontroller))
	http.HandleFunc("/add", createAddHandler(peercontroller))
	log.Fatal(http.ListenAndServe(peercontroller.Self.Address, nil))
}

//Implement add api method to update peercontroller peer data
func createAddHandler(peercontroller *PeerData) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		addingpeer := Peer{}
		decode := json.NewDecoder(r.Body)
		err := decode.Decode(&addingpeer)
		
		if err != nil {
			log.Printf("Add Peer Error: %v", err)
		}

		peercontroller.addPeers <- addingpeer
		peercontroller.getCurrentPeers <- true
		
		encode := json.NewEncoder(w)
		encode.Encode(<-peercontroller.currentPeers)
	}
}

//Implement Tx api method to update peercontroller msg data
func createTxHandler(peercontroller *PeerData) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mg := Msg{}
		decode := json.NewDecoder(r.Body)
		err := decode.Decode(&mg)
		if err != nil {
			log.Printf("Message Error: %v", err)
		}

		peercontroller.receivedMsg <- mg
	}
}

//Parse stdin for text and update peercontroller msg data
func (peercontroller *PeerData) startStdinListener(sender Peer) {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		message := line[:len(line)-1]
		peercontroller.userMsg <- Msg{message, sender}
	}
}
