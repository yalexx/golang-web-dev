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

			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]

			fmt.Println("***METHOD", m)
			fmt.Println("***URL", u)

			if m == "GET" && u == "/" {
				respond(c, "You are on GET /")
			}
			if m == "GET" && u == "/apply" {
				respond(c, "You are on GET /apply")
			}
			if m == "POST" && u == "/apply" {
				respond(c, "You are on POST /")
			}

		}

		i++

	}
}

func respond(c net.Conn, s string) {

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>"%v"</h1>
		</body>
		</html>
	`, s)

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}
