package main

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

const (
	ResponseTemplate = `HTTP/1.1 %d OK
Date: Thu, 09 Dec 2004 12:07:48 GMT
Server: Golang Scuff
Content-type: text/plain

%s`
)

func main() {
	lgr, err := zap.NewDevelopment()

	port := "8090"
	lgr.Info("listening", zap.String("port", port))

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		lgr.Error("failed to listen to port")
		return
	}

	for {
		con, err := lis.Accept()
		if err != nil {
			lgr.Error("failed to accept connection")
			lis.Close()
			return
		}

		lgr.Info("incoming connection")

		buf := make([]byte, 1024)
		_, err = con.Read(buf)
		if err != nil {
			lgr.Error("failed to read bytes")
			continue
		}
		println(string(buf))
		status := 200
		message := "wowie"

		len, err := con.Write([]byte(fmt.Sprintf(ResponseTemplate, status, message)))
		if err != nil {
			lgr.Error("failed to write")
			continue
		}
		lgr.Info("writen", zap.Int("len", len))

		lgr.Info("closing connection")
		err = con.Close()
		if err != nil {
			lgr.Error("failed to close connection")
			continue
		}

	}
}
