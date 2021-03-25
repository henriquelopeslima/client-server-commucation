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

func serverA(serverPort string) {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+serverPort)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClientA(conn)
	}
}

func serverB(serverPort string) {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+serverPort)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClientB(conn)
	}
}

func handleClientA(conn *net.UDPConn) {
	bufferRecv := make([]byte, BufferSize)
	var packetRecv packetA

	_, addr, _ := conn.ReadFromUDP(bufferRecv)

	_ = binary.Read(bytes.NewReader(bufferRecv), binary.BigEndian, &packetRecv)

	fmt.Println("Pacote recebido do cliente")
	fmt.Println("header: ", packetRecv.Header)
	fmt.Println("message: ", string(packetRecv.Message[:]))

	if string(packetRecv.Message[:packetRecv.Header.PayloadLen]) == "Hello world" {
		// A2
		responsePack := responseA{
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

		go serverB(strconv.Itoa(int(responsePack.UdpPort)))
	}
}

func handleClientB(conn *net.UDPConn) {
	fmt.Println("Segunda conexão")
	bufferRecv := make([]byte, BufferSize)
	var requestB packetB

	_, addr, _ := conn.ReadFromUDP(bufferRecv)

	_ = binary.Read(bytes.NewReader(bufferRecv), binary.BigEndian, &requestB)

	fmt.Println("Pacote recebido do cliente - agora estamos no passo B")
	fmt.Println("header: ", requestB.Header)
	fmt.Println("packetId: ", requestB.PacketId)
	fmt.Println("payload: ", requestB.Payload)

	// B2
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, requestB)
	checkError(err)

	_, _ = conn.WriteToUDP(buffer.Bytes(), addr)
}

// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
	serverA(portUDP)
}
