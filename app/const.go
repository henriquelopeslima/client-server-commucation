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
	Matriculation int16
}

type packet struct {
	Header  header
	Message [200]byte
}

type response struct {
	Header  header
	Num     uint16
	Len     uint16
	UdpPort uint16
	SecretA uint16
}

type ack struct {
	Header        header
	AckedPacketId uint16
}

type request struct {
	Header   header
	PacketId uint16
	Payload  [200]byte
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

func printResponse(param response) {
	fmt.Println("Response ")
	printHeader(param.Header, " - ")
	fmt.Println(" - Len ", param.Len)
	fmt.Println(" - Num ", param.Num)
	fmt.Println(" - UdpPort ", param.UdpPort)
	fmt.Println(" - SecretA ", param.SecretA)
}

func getNumber() uint16 {
	min := 10
	max := 300
	return uint16(rand.Intn(max-min+1) + min)
}
