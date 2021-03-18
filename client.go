/*****************************************************************************
 * client.go
 * Nome: Henrique Lopes Lima, Gabriela Miranda Leal
 * Matrícula: 413031, 398624
 *****************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const SendBufferSize = 2048
const serverUDP = ":12235"

func client(serverIp string, serverPort string) {
	//UDPAddr
	udpAddr, err := net.ResolveUDPAddr("udp", serverIp + serverPort)
	checkErrorClient(err)

	//UDPConn
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkErrorClient(err)
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, SendBufferSize)

	for {
		readTotal, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				checkErrorClient(err)
			}
			break
		}
		_, err = conn.Write(buf[:readTotal])
		checkErrorClient(err)
	}

	n, _ := conn.Read(buf[0:])
	fmt.Println(string(buf[0:n]))

	checkErrorClient(err)
	os.Exit(0)
}

// Main obtém argumentos da linha de comando e chama função client
func main() {
	client("127.0.0.1", serverUDP)
}

func checkErrorClient(err error){
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
		os.Exit(1)
	}
}