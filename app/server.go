/*****************************************************************************
 * server.go
 * Nome: Henrique Lopes Lima, Gabriela Miranda Leal
 * Matrícula: 413031, 398624
 *****************************************************************************/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
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

func serverB(serverPort string, pack responseA) {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+serverPort)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClientB(conn, pack)
	}
}

func serverC(serverPort string, pack responseB) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+serverPort)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		checkError(err)
		handleClientC(conn, pack)
	}
}

func handleClientA(conn *net.UDPConn) {
	fmt.Println("Conexão A")

	var packetRecv packetA
	bufferRecv := make([]byte, BufferSize)
	_, addr, _ := conn.ReadFromUDP(bufferRecv)
	dec := gob.NewDecoder(bytes.NewReader(bufferRecv))
	err := dec.Decode(&packetRecv)
	checkError(err)

	fmt.Println("Pacote recebido do cliente")
	printHello(packetRecv)

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
		enc := gob.NewEncoder(buffer)
		err := enc.Encode(&responsePack)
		checkError(err)

		_, _ = conn.WriteToUDP(buffer.Bytes(), addr)

		go serverB(strconv.Itoa(int(responsePack.UdpPort)), responsePack)
	}
}

func handleClientB(conn *net.UDPConn, pack responseA) {
	fmt.Println("Conexão B")

	ackRecv := ack{
		Header: pack.Header,
	}

	var address *net.UDPAddr

	for i := 1; i <= int(pack.Num); i++ {
		var requestB packetB
		bufferRecv := make([]byte, BufferSize)

		_, addr, _ := conn.ReadFromUDP(bufferRecv)
		dec := gob.NewDecoder(bytes.NewReader(bufferRecv))
		err := dec.Decode(&requestB)
		checkError(err)

		address = addr

		printRequest(requestB)

		ackRecv.AckedPacketId = requestB.PacketId

		//B2
		buffer := new(bytes.Buffer)

		enc := gob.NewEncoder(buffer)
		err = enc.Encode(&ackRecv)
		checkError(err)

		_, _ = conn.WriteToUDP(buffer.Bytes(), addr)
	}

	var _responseB = responseB{
		Header:  pack.Header,
		TcpPort: getNumber() + 5000,
		SecretB: getNumber(),
	}

	buffer := new(bytes.Buffer)

	enc := gob.NewEncoder(buffer)
	err := enc.Encode(&_responseB)
	checkError(err)

	_, _ = conn.WriteToUDP(buffer.Bytes(), address)

	conn.Close()

	serverC(strconv.Itoa(int(_responseB.TcpPort)), _responseB)
}

func handleClientC(conn net.Conn, pack responseB) {
	fmt.Println("Conexão C")

	randomChar := string('a' + rune(rand.Intn(26)))
	_responseC := responseC{
		Header:  pack.Header,
		Len2:    getNumber(),
		Num2:    getNumber(),
		SecretC: getNumber(),
		C:       randomChar,
	}

	var bufferSend = new(bytes.Buffer)

	bufferSend = new(bytes.Buffer)
	enc := gob.NewEncoder(bufferSend)
	err := enc.Encode(&_responseC)
	checkError(err)

	_, err = conn.Write(bufferSend.Bytes())

	// D1 receiver
	for i := 1; i <= int(_responseC.Num2); i++ {
		var _packD packetD
		bufferRecv := make([]byte, BufferSize)
		_, err = conn.Read(bufferRecv)
		checkError(err)

		dec := gob.NewDecoder(bytes.NewReader(bufferRecv))
		err = dec.Decode(&_packD)
		checkError(err)
		printPackD(_packD)
		fmt.Println(_responseC.Num2, " - ", i)
	}

	// D2 send
	//_responseD := responseD{
	//	Header: pack.Header,
	//	SecretD: getNumber(),
	//}
	//
	//var bufferSend2 = new(bytes.Buffer)
	//
	//bufferSend = new(bytes.Buffer)
	//enc = gob.NewEncoder(bufferSend2)
	//err = enc.Encode(&_responseD)
	//checkError(err)
}

// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
	serverA(portUDP)
}
