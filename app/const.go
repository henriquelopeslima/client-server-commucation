package main

import (
	"fmt"
	"math/rand"
	"os"
)

const BufferSize = 1024
const portUDP = "12235"
const matriculation = 624

type header struct {
	PayloadLen    uint32
	PSecret       uint32
	Step          uint16
	Matriculation uint32
}

type packetA struct {
	Header  header
	Message [200]byte
}

type packetB struct {
	Header   header
	PacketId uint32
	Payload  uint32
}

type packetD struct {
	Header  header
	Payload string
}

type responseA struct {
	Header  header
	Num     uint32
	Len     uint32
	UdpPort uint32
	SecretA uint32
}

type responseB struct {
	Header  header
	TcpPort uint32
	SecretB uint32
}

type responseC struct {
	Header  header
	Num2    uint32
	Len2    uint32
	SecretC uint32
	C       string
}

type responseD struct {
	Header  header
	SecretD uint32
}

type ack struct {
	Header        header
	AckedPacketId uint32
}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
		os.Exit(1)
	}
}

func printHeader(param header, previous string) {
	fmt.Println(previous + "Header")
	fmt.Println(previous+" - PayloadLen ", param.PayloadLen)
	fmt.Println(previous+" - PSecret ", param.PSecret)
	fmt.Println(previous+" - Step ", param.Step)
	fmt.Println(previous+" - Matriculation ", param.Matriculation)
}

func printResponseA(param responseA) {
	fmt.Println("Response A")
	printHeader(param.Header, " - ")
	fmt.Println(" - Len ", param.Len)
	fmt.Println(" - Num ", param.Num)
	fmt.Println(" - UdpPort ", param.UdpPort)
	fmt.Println(" - SecretA ", param.SecretA)
}

func printResponseB(param responseB) {
	fmt.Println("Response B")
	printHeader(param.Header, " - ")
	fmt.Println(" - TcpPort ", param.TcpPort)
	fmt.Println(" - SecretB ", param.SecretB)
}

func printResponseC(param responseC) {
	fmt.Println("Response C")
	printHeader(param.Header, " - ")
	fmt.Println(" - Num2 ", param.Num2)
	fmt.Println(" - Len2 ", param.Len2)
	fmt.Println(" - SecretC ", param.SecretC)
	fmt.Println(" - char ", string(param.C[:]))
}

func printResponseD(param responseD) {
	fmt.Println("Response D")
	printHeader(param.Header, " - ")
	fmt.Println(" - SecretD ", param.SecretD)
}

func printHello(param packetA) {
	fmt.Println("Request")
	printHeader(param.Header, " - ")
	fmt.Println(" - message ", string(param.Message[:]))
}

func printRequest(param packetB) {
	fmt.Println("Request")
	printHeader(param.Header, " - ")
	fmt.Println(" - PacketID ", param.PacketId)
	fmt.Println(" - Payload ", param.Payload)
}

func printPackD(param packetD) {
	fmt.Println("Request")
	printHeader(param.Header, " - ")
	fmt.Println(" - Payload ", string(param.Payload[:]))
}

func printAck(param ack) {
	fmt.Println("Request")
	printHeader(param.Header, " - ")
	fmt.Println(" - AckedPacketId ", param.AckedPacketId)
}

func getNumber(max int, min int) uint32 {
	return uint32(rand.Intn(max-min+1) + min)
}
