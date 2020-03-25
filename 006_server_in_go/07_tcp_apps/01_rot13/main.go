package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		fs:= strings.Fields(ln)
		var value int
		if len(fs) == 1 {
			value = 13
		}else{
			var err error
			value,err = strconv.Atoi(fs[1])
			if err != nil {
				value= 13
			}
		}


		fmt.Println(value)
		fmt.Println(fs)
		r := rot13(fs[0], value)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

func rot13(str string, key int) string {
	runes := []rune(str)

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, char := range runes {
		index := strings.Index(alphabet, string(char))
		if index == -1 {
			return ""
		}
		newIndex := (index + key) % 26
		runes[i] = rune(alphabet[newIndex])
	}
	return string(runes)
}
