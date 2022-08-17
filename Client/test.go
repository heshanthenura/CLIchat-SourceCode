package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var nickName string
var con net.Conn

func main() {
	connectServer()
}

func validateNickName() {
	fmt.Print("Enter NickName:- ")
	in := bufio.NewReader(os.Stdin)
	nickName, _ = in.ReadString('\n')
	for len(nickName) < 3 {
		fmt.Print("Enter Valid NickName:- ")
		in := bufio.NewReader(os.Stdin)
		nickName, _ = in.ReadString('\n')
	}
	fmt.Printf("Sending NickName as %s", nickName)
}

func connectServer() {
	for true {
		var addr string
		fmt.Print("Enter Address:- ")
		in := bufio.NewReader(os.Stdin)
		addr, _ = in.ReadString('\n')
		addr = strings.TrimSpace(addr)
		con, _ = net.Dial("tcp", addr)
		if con == nil {
			fmt.Println("Error: Check address again or Server Offline")
		} else {
			fmt.Println("Connected to Server " + addr)
			break
		}
	}

}
