package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

// Global Variables
var userName string
var conn net.Conn

func main() {
	// Global Variables

	//Wait group for Goroutines
	wg := new(sync.WaitGroup)
	wg.Add(1)

	//Asking and Validate UserName
	validateUserName()

	//Connecting to Server
	connectServer()

	//Sending the Username
	fmt.Fprintf(conn, userName)

	// Goroutine Functions
	go readMsg(wg, conn)
	go writeMessage(wg, conn)
	wg.Wait()
}

//Read Message Function
func readMsg(wg *sync.WaitGroup, conn net.Conn) {
	for true {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
	wg.Done()
}

// Write Message
func writeMessage(wg *sync.WaitGroup, conn net.Conn) {
	for true {
		var message string
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')
		for err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("MSG: %s", message)
		fmt.Fprintf(conn, message)
	}
	wg.Done()
}

func closeConnection(wg *sync.WaitGroup, conn net.Conn) {
	fmt.Println(conn.Close())
}

// Asking and Validating Name
func validateUserName() {
	fmt.Print("Enter NickName:- ")
	in := bufio.NewReader(os.Stdin)
	userName, _ = in.ReadString('\n')
	for len(userName) < 3 {
		fmt.Print("Enter Valid UserName:- ")
		in := bufio.NewReader(os.Stdin)
		userName, _ = in.ReadString('\n')
	}
	fmt.Printf("Sending UserName as %s", userName)
}

//Connecting to Server
func connectServer() {
	for true {
		var addr string
		fmt.Print("Enter Address:- ")
		in := bufio.NewReader(os.Stdin)
		addr, _ = in.ReadString('\n')
		addr = strings.TrimSpace(addr)
		conn, _ = net.Dial("tcp", addr)
		if conn == nil {
			fmt.Println("Error: Check address again or Server Offline")
		} else {
			fmt.Println("Connected to Server " + addr)
			break
		}
	}

}
