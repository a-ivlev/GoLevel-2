package main

import (
	"fmt"

	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

const version = "2.0.2"

func main() {
	fmt.Printf("Project version %s", version)
}
