/*****************************************************************************
 * server.go
 * Nome: Henrique Lopes Lima, Gabriela Miranda Leal
 * Matrícula: 413031, 398624
 *****************************************************************************/

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const RecvBufferSize = 2048
const portUDP = ":12235"

/* TODO: server()
 * Abra socket e espere o cliente conectar
 * Imprima a mensagem recebida em stdout
 */
func server(serverPort string) {

	udpAddr, err := net.ResolveUDPAddr("udp", serverPort)
	checkErrorServer(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkErrorServer(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn)  {
	var buf [RecvBufferSize]byte
	for {
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return
		}
		if buf[0:] != nil {
			fmt.Print(string(buf[0:n]))
		}
		dayTime := time.Now().String()
		_, _ = conn.WriteToUDP([]byte(dayTime), addr)
	}
}

// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
	server(portUDP)
}

func checkErrorServer(err error){
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
		os.Exit(1)
	}
}