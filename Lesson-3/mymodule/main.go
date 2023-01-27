package main

import (
	"fmt"

	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

const version="v1.0.1"

func main() {
	fmt.Printf("Project version %s", version)
}
