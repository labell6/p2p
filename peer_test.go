package main

import (
  "os/exec"
  "testing"
  "log"
  "strings"
)


type TestMsg struct {
        Message string
        From    Peer
}

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






     
