package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	m, u := request(conn)

	fmt.Println("Method and Url: ", m, u)
	// write response

	if u == "/hello" {
		respond(conn, `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello</strong></body></html>`)
	}
	{
		respond(conn, `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>KUR</strong></body></html>`)
	}

}

func request(conn net.Conn) (string, string) {
	i := 0
	method := ""
	url := ""

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)
		if i == 0 {
			method = strings.Fields(ln)[0]
			url = strings.Fields(ln)[1]
		}
		if ln == "" {
			break
		}
		i++
	}
	return method, url
}

func respond(conn net.Conn, body string) {

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
