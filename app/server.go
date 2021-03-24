/*****************************************************************************
 * server.go
 * Nome: Henrique Lopes Lima, Gabriela Miranda Leal
 * Matrícula: 413031, 398624
 *****************************************************************************/

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

func server(serverPort string) {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+serverPort)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	bufferRecv := make([]byte, BufferSize)
	var packetRecv packet

	_, addr, _ := conn.ReadFromUDP(bufferRecv)

	_ = binary.Read(bytes.NewReader(bufferRecv), binary.BigEndian, &packetRecv)

	fmt.Println("Pacote recebido do cliente")
	fmt.Println("header: ", packetRecv.Header)
	fmt.Println("message: ", string(packetRecv.Message[:]))

	if string(packetRecv.Message[:packetRecv.Header.PayloadLen]) == "Hello world" {
		// A2
		responsePack := response{
			Header: header{
				PayloadLen:    11,
				PSecret:       0,
				Step:          2,
				Matriculation: matriculation,
			},
			Len:     getNumber(),
			Num:     getNumber(),
			UdpPort: getNumber() + 5000,
			SecretA: getNumber(),
		}

		buffer := new(bytes.Buffer)

		err := binary.Write(buffer, binary.BigEndian, responsePack)
		checkError(err)

		_, _ = conn.WriteToUDP(buffer.Bytes(), addr)

		go server(strconv.Itoa(int(responsePack.UdpPort)))
	}
}

// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
	server(portUDP)
}
