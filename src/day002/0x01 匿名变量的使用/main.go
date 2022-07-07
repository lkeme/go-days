package main

import "fmt"

func info(content string) (msg string, err error) {
	return content + "\r\n", nil
}

func main() {
	msg, _ := info("Hello world")
	fmt.Println(msg)
}
