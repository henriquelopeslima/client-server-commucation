/*****************************************************************************
 * client.go
 * Nome: Henrique Lopes Lima, Gabriela Miranda Leal
 * Matrícula: 413031, 398624
 *****************************************************************************/

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
)

func client(serverIp string, serverPort string) {
	//UDPAddr
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+serverPort)
	checkError(err)

	//UDPConn
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	bufferSend := new(bytes.Buffer)

	// A1
	helloPack := packetA{
		Header: header{
			PayloadLen:    11,
			PSecret:       0,
			Step:          1,
			Matriculation: 624,
		},
		Message: [200]byte{'H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'},
	}

	err = binary.Write(bufferSend, binary.BigEndian, helloPack)
	checkError(err)

	_, err = conn.Write(bufferSend.Bytes())
	checkError(err)

	bufferRecv := make([]byte, BufferSize)
	var responseServer responseA

	_, err = conn.Read(bufferRecv)
	checkError(err)

	err = binary.Read(bytes.NewReader(bufferRecv), binary.BigEndian, &responseServer)
	checkError(err)

	conn.Close()

	fmt.Println("A2 resposta do servidor")
	printResponse(responseServer)

	//UDPAddr
	udpAddr, err = net.ResolveUDPAddr("udp", ":"+strconv.Itoa(int(responseServer.UdpPort)))
	checkError(err)

	//UDPConn
	conn, err = net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	packetBSend := packetB{
		Header: header{
			PayloadLen:    responseServer.Len,
			PSecret:       responseServer.SecretA,
			Step:          2,
			Matriculation: matriculation,
		},
		PacketId: 123,
		Payload:  1,
	}

	err = binary.Write(bufferSend, binary.BigEndian, packetBSend)
	checkError(err)

	_, err = conn.Write(bufferSend.Bytes())
	checkError(err)

	os.Exit(0)
}

// Main obtém argumentos da linha de comando e chama função client
func main() {
	client("127.0.0.1", portUDP)
}
