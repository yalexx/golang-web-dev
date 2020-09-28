package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(c net.Conn) {

	i := 0
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			// the REQUEST method

			fmt.Println("***METHOD", strings.Fields(ln)[0])

			//REQUEST URI

			fmt.Println("***URL", strings.Fields(ln)[1])
		}

		i++

		fmt.Fprintf(c, "Requesr lines: : %s\n", ln)

		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}

		body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
		io.WriteString(c, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(c, "Content-Type: text/plain\r\n")
		io.WriteString(c, "\r\n")
		io.WriteString(c, body)

	}
}
