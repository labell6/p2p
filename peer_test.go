//Implementation of a simple peer to peer chat application using
//the net/http package for communication.

package main

import (
	"net/http"
	"reflect"
	"testing"
        "os/exec"
        "log"
        "strings"
)

type TestPeer struct {
        Name    string
        Address string
}


func TestGetIp(t *testing.T) {

  ipAddr := GetIp()
   
  out, _ := exec.Command("ping",  ipAddr, "-c 5", "-i 3", "-w 10").Output()
  if strings.Contains(string(out), "Destination Host Unreachable") {
    t.Errorf("GetIP failed, IP address not found  !")
  }

  log.Println("IP Address %s detected.",ipAddr)
}

func TestData(t *testing.T) {
        port := "7500"
        name := "Surematics"
        peer := "127.0.0.1"

        //initialize client control handler
        pdata := Peer{name, peer+":"+port}
        log.Println("pdata.Name detected: ",pdata.Name)

        if pdata.Name != "Surematics" {
           t.Errorf("Access Denied Peer.Name")
	}

        if pdata.Address != "127.0.0.1:7500" {
           t.Errorf("Access Denied Peer.Address")
        }


        testpeercontroller := NewPeerData(pdata)
        testpeercontroller.start()

        if (&testpeercontroller == nil) {
           t.Errorf("Error Initializing server")
        }
        
        log.Println("server detected: ",&testpeercontroller)
}

func TestNewPeerData(t *testing.T) {
	type args struct {
		self Peer
	}
	tests := []struct {
		name string
		args args
		want *PeerData
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPeerData(tt.args.self); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPeerData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeerData_start(t *testing.T) {
	type fields struct {
		Self            Peer
		Peers           Peers
		addPeers        chan (Peer)
		exitPeer        chan (Peer)
		currentPeers    chan (Peers)
		getCurrentPeers chan (bool)
		receivedMsg     chan (Msg)
		userMsg         chan (Msg)
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peercontroller := &PeerData{
				Self:            tt.fields.Self,
				Peers:           tt.fields.Peers,
				addPeers:        tt.fields.addPeers,
				exitPeer:        tt.fields.exitPeer,
				currentPeers:    tt.fields.currentPeers,
				getCurrentPeers: tt.fields.getCurrentPeers,
				receivedMsg:     tt.fields.receivedMsg,
				userMsg:         tt.fields.userMsg,
			}
			peercontroller.start()
		})
	}
}

func TestPeerData_cmdLoop(t *testing.T) {
	type fields struct {
		Self            Peer
		Peers           Peers
		addPeers        chan (Peer)
		exitPeer        chan (Peer)
		currentPeers    chan (Peers)
		getCurrentPeers chan (bool)
		receivedMsg     chan (Msg)
		userMsg         chan (Msg)
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peercontroller := &PeerData{
				Self:            tt.fields.Self,
				Peers:           tt.fields.Peers,
				addPeers:        tt.fields.addPeers,
				exitPeer:        tt.fields.exitPeer,
				currentPeers:    tt.fields.currentPeers,
				getCurrentPeers: tt.fields.getCurrentPeers,
				receivedMsg:     tt.fields.receivedMsg,
				userMsg:         tt.fields.userMsg,
			}
			peercontroller.cmdLoop()
		})
	}
}

func TestPeerData_existingPeer(t *testing.T) {
	type fields struct {
		Self            Peer
		Peers           Peers
		addPeers        chan (Peer)
		exitPeer        chan (Peer)
		currentPeers    chan (Peers)
		getCurrentPeers chan (bool)
		receivedMsg     chan (Msg)
		userMsg         chan (Msg)
	}
	type args struct {
		peer Peer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peercontroller := &PeerData{
				Self:            tt.fields.Self,
				Peers:           tt.fields.Peers,
				addPeers:        tt.fields.addPeers,
				exitPeer:        tt.fields.exitPeer,
				currentPeers:    tt.fields.currentPeers,
				getCurrentPeers: tt.fields.getCurrentPeers,
				receivedMsg:     tt.fields.receivedMsg,
				userMsg:         tt.fields.userMsg,
			}
			if got := peercontroller.existingPeer(tt.args.peer); got != tt.want {
				t.Errorf("PeerData.existingPeer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeerData_sendAdd(t *testing.T) {
	type fields struct {
		Self            Peer
		Peers           Peers
		addPeers        chan (Peer)
		exitPeer        chan (Peer)
		currentPeers    chan (Peers)
		getCurrentPeers chan (bool)
		receivedMsg     chan (Msg)
		userMsg         chan (Msg)
	}
	type args struct {
		peer Peer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peercontroller := &PeerData{
				Self:            tt.fields.Self,
				Peers:           tt.fields.Peers,
				addPeers:        tt.fields.addPeers,
				exitPeer:        tt.fields.exitPeer,
				currentPeers:    tt.fields.currentPeers,
				getCurrentPeers: tt.fields.getCurrentPeers,
				receivedMsg:     tt.fields.receivedMsg,
				userMsg:         tt.fields.userMsg,
			}
			peercontroller.sendAdd(tt.args.peer)
		})
	}
}

func TestPeerData_sendTx(t *testing.T) {
	type fields struct {
		Self            Peer
		Peers           Peers
		addPeers        chan (Peer)
		exitPeer        chan (Peer)
		currentPeers    chan (Peers)
		getCurrentPeers chan (bool)
		receivedMsg     chan (Msg)
		userMsg         chan (Msg)
	}
	type args struct {
		peer Peer
		msg  Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peercontroller := &PeerData{
				Self:            tt.fields.Self,
				Peers:           tt.fields.Peers,
				addPeers:        tt.fields.addPeers,
				exitPeer:        tt.fields.exitPeer,
				currentPeers:    tt.fields.currentPeers,
				getCurrentPeers: tt.fields.getCurrentPeers,
				receivedMsg:     tt.fields.receivedMsg,
				userMsg:         tt.fields.userMsg,
			}
			peercontroller.sendTx(tt.args.peer, tt.args.msg)
		})
	}
}

func TestPeerData_webListener(t *testing.T) {
	type fields struct {
		Self            Peer
		Peers           Peers
		addPeers        chan (Peer)
		exitPeer        chan (Peer)
		currentPeers    chan (Peers)
		getCurrentPeers chan (bool)
		receivedMsg     chan (Msg)
		userMsg         chan (Msg)
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peercontroller := &PeerData{
				Self:            tt.fields.Self,
				Peers:           tt.fields.Peers,
				addPeers:        tt.fields.addPeers,
				exitPeer:        tt.fields.exitPeer,
				currentPeers:    tt.fields.currentPeers,
				getCurrentPeers: tt.fields.getCurrentPeers,
				receivedMsg:     tt.fields.receivedMsg,
				userMsg:         tt.fields.userMsg,
			}
			peercontroller.webListener()
		})
	}
}

func Test_createAddHandler(t *testing.T) {
	type args struct {
		peercontroller *PeerData
	}
	tests := []struct {
		name string
		args args
		want func(http.ResponseWriter, *http.Request)
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createAddHandler(tt.args.peercontroller); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createAddHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createTxHandler(t *testing.T) {
	type args struct {
		peercontroller *PeerData
	}
	tests := []struct {
		name string
		args args
		want func(http.ResponseWriter, *http.Request)
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTxHandler(tt.args.peercontroller); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTxHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeerData_startStdinListener(t *testing.T) {
	type fields struct {
		Self            Peer
		Peers           Peers
		addPeers        chan (Peer)
		exitPeer        chan (Peer)
		currentPeers    chan (Peers)
		getCurrentPeers chan (bool)
		receivedMsg     chan (Msg)
		userMsg         chan (Msg)
	}
	type args struct {
		sender Peer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			peercontroller := &PeerData{
				Self:            tt.fields.Self,
				Peers:           tt.fields.Peers,
				addPeers:        tt.fields.addPeers,
				exitPeer:        tt.fields.exitPeer,
				currentPeers:    tt.fields.currentPeers,
				getCurrentPeers: tt.fields.getCurrentPeers,
				receivedMsg:     tt.fields.receivedMsg,
				userMsg:         tt.fields.userMsg,
			}
			peercontroller.startStdinListener(tt.args.sender)
		})
	}
}
