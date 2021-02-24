package main

import (
	"flag"
	"fmt"
	"github.com/fastgo3/fastgo3"
)

var (
	ip	= flag.String("ip", "0.0.0.0", "IP or Host")
	port = flag.Int("port", 3030, "Listen Port")
)

func main() {
	flag.Parse()

	app := fastgo3.New()

	app.Get("/users/<id:i>", dynHandler)
	app.Run(*ip, *port)
}

func dynHandler(ctx *fastgo3.Context) {
	fmt.Fprintf(ctx, "Post string is %q\n", ctx.PostArgs())
	fmt.Fprintf(ctx, "password=%s\n", ctx.StrParam("password", ""))
}
