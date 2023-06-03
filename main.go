package main

import (
	"bufio"
	"crypto/rand"
	"math/big"
	"net"
)

var conns map[string]net.Conn

func main() {
	chat, _ := net.Listen("tcp", ":5000")
	conns = make(map[string]net.Conn)
	for {
		conn, _ := chat.Accept()
		nickInt, _ := rand.Int(rand.Reader, big.NewInt(27))
		conns[nickInt.String()] = conn
		stream := bufio.NewReader(conn)
		go func() {
			for {
				line, err := stream.ReadBytes('\n')
				if err != nil {
					break
				}
				for connNick, conn := range conns {
					if connNick != nickInt.String() {
						conn.Write([]uint8(nickInt.String() + ": " + string(line[:len(line)-1]) + "\r\n"))
					}
				}
			}
			conn.Close()
		}()
	}
}
