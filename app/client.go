/*****************************************************************************
 * client.go
 * Nome: Henrique Lopes Lima, Gabriela Miranda Leal
 * Matrícula: 413031, 398624
 *****************************************************************************/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func client(serverIp string, serverPort string) {
	//UDPAddr
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+serverPort)
	checkError(err)

	//UDPConn
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	var bufferSend = new(bytes.Buffer)

	// Aqui estaremos enviando o pacote Hello
	helloPack := packetA{
		Header: header{
			PayloadLen:    11,
			PSecret:       0,
			Step:          1,
			Matriculation: 624,
		},
		Message: [200]byte{'H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'},
	}

	enc := gob.NewEncoder(bufferSend)
	err = enc.Encode(helloPack)
	_, err = conn.Write(bufferSend.Bytes())
	checkError(err)

	bufferRecv := make([]byte, BufferSize)
	var responseServer responseA

	_, err = conn.Read(bufferRecv)
	checkError(err)

	dec := gob.NewDecoder(bytes.NewReader(bufferRecv))
	err = dec.Decode(&responseServer)
	checkError(err)

	conn.Close()

	fmt.Println("A2 resposta do servidor")
	printResponseA(responseServer)

	//UDPAddr
	udpAddr, err = net.ResolveUDPAddr("udp", ":"+strconv.Itoa(int(responseServer.UdpPort)))
	checkError(err)

	//UDPConn
	conn, err = net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	for i := 1; i <= int(responseServer.Num); i++ {
		packetBSend := packetB{
			Header: header{
				PayloadLen:    responseServer.Len,
				PSecret:       responseServer.SecretA,
				Step:          2,
				Matriculation: matriculation,
			},
			PacketId: uint32(i),
			Payload:  1,
		}

		bufferSend = new(bytes.Buffer)
		enc = gob.NewEncoder(bufferSend)
		err = enc.Encode(&packetBSend)
		checkError(err)

		time.Sleep(500 * time.Millisecond)
		_, err = conn.Write(bufferSend.Bytes())
		checkError(err)

		// Receiver
		bufferRecv = make([]byte, BufferSize)
		var ackRecv ack
		_, err = conn.Read(bufferRecv)
		checkError(err)

		dec = gob.NewDecoder(bytes.NewReader(bufferRecv))
		err = dec.Decode(&ackRecv)
		checkError(err)

		printAck(ackRecv)
	}

	bufferRecv = make([]byte, BufferSize)
	var _responseServerB responseB
	_, err = conn.Read(bufferRecv)
	dec = gob.NewDecoder(bytes.NewReader(bufferRecv))
	err = dec.Decode(&_responseServerB)
	checkError(err)
	fmt.Println("Agora vamos iniciar o novo passo : )\nRumo ao Hexa (Finalizando o passo B2)")
	printResponseB(_responseServerB)

	//TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(int(_responseServerB.TcpPort)))
	checkError(err)

	//TCPConn
	connTCP, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	connTCP.LocalAddr()

	// C2 receiver
	var _responseC responseC
	bufferRe := make([]byte, BufferSize)
	_, err = connTCP.Read(bufferRe)
	checkError(err)

	dec = gob.NewDecoder(bytes.NewReader(bufferRe))
	err = dec.Decode(&_responseC)
	checkError(err)
	printResponseC(_responseC)

	// D1 send
	for i := 1; i <= int(_responseC.Num2); i++ {
		packetCSend := packetD{
			Header: header{
				PayloadLen:    responseServer.Len,
				PSecret:       responseServer.SecretA,
				Step:          2,
				Matriculation: matriculation,
			},
			Payload: _responseC.C,
		}

		bufferSend := new(bytes.Buffer)
		enc = gob.NewEncoder(bufferSend)
		err = enc.Encode(&packetCSend)
		checkError(err)

		time.Sleep(100 * time.Millisecond)
		_, err = connTCP.Write(bufferSend.Bytes())
		checkError(err)
	}

	// D2 receiver
	var _responseD responseD
	bufferRecv = make([]byte, BufferSize)
	_, err = connTCP.Read(bufferRecv)
	checkError(err)
	dec = gob.NewDecoder(bytes.NewReader(bufferRecv))
	err = dec.Decode(&_responseD)
	printResponseD(_responseD)

	os.Exit(0)
}

// Main obtém argumentos da linha de comando e chama função client
func main() {
	client("127.0.0.1", portUDP)
}
