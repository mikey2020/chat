package lib

import (
	"os"
	"fmt"
	"log"
	"net"
	"bufio"
)

const port = "8080"

func RunHost(ip string){
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ",listenErr)
	}
	fmt.Println("Listening on ", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
	fmt.Println("New connection accepted")
    for {
        handleHost(conn)
	}

}

func handleHost(conn net.Conn){
	reader := bufio.NewReader(conn)
	message,readError := reader.ReadString('\n')
	if readError != nil {
		log.Fatal("Error: ", readError)
	}
	fmt.Println("message received : ", message)
	
	fmt.Print("Enter message: ")
	replyReader := bufio.NewReader(os.Stdin)
	message, readErr := replyReader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Fprint(conn,message)
}


func RunGuest(ip string){
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp",ipAndPort)
	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}
	for {
		handleGuest(conn)
	}

}

func handleGuest(conn net.Conn){
	fmt.Print("Enter message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Fprint(conn,message)

	msgReader := bufio.NewReader(conn)
	message,readError := msgReader.ReadString('\n')
	if readError != nil {
		log.Fatal("Error: ", readError)
	}
	fmt.Println("message received : ", message)

}
