package main

import (
	"fmt"
	"math/rand"
	"os"
)

const BufferSize = 2048
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

type responseA struct {
	Header  header
	Num     uint32
	Len     uint32
	UdpPort uint32
	SecretA uint32
}

type ack struct {
	Header        header
	AckedPacketId uint32
}

type packetB struct {
	Header   header
	PacketId uint32
	Payload  uint32
}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
		os.Exit(1)
	}
}

func printHeader(param header, previous string) {
	fmt.Println(previous + "Header ")
	fmt.Println(previous+" - PayloadLen ", param.PayloadLen)
	fmt.Println(previous+" - PSecret ", param.PSecret)
	fmt.Println(previous+" - Step ", param.Step)
	fmt.Println(previous+" - Matriculation ", param.Matriculation)
}

func printResponse(param responseA) {
	fmt.Println("Response ")
	printHeader(param.Header, " - ")
	fmt.Println(" - Len ", param.Len)
	fmt.Println(" - Num ", param.Num)
	fmt.Println(" - UdpPort ", param.UdpPort)
	fmt.Println(" - SecretA ", param.SecretA)
}

func printRequest(param packetB) {
	fmt.Println("Request ")
	printHeader(param.Header, " - ")
	fmt.Println(" - PacketID ", param.PacketId)
	fmt.Println(" - Payload ", param.Payload)
}

func printAck(param ack) {
	fmt.Println("Request ")
	printHeader(param.Header, " - ")
	fmt.Println(" - AckedPacketId ", param.AckedPacketId)
}

func getNumber() uint32 {
	min := 10
	max := 300
	return uint32(rand.Intn(max-min+1) + min)
}
