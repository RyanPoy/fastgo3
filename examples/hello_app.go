package main

import (
	"flag"
	"fmt"
	"fastgo3"
)

var (
	ip	= flag.String("ip", "0.0.0.0", "IP or Host")
	port = flag.Int("port", 3031, "Listen Port")
)

func main() {
	flag.Parse()

	app := fastgo3.NewApplication()
	app.Amount([]fastgo3.Route{
		fastgo3.Post("/hello", helloAction),
	})
	app.Run(*ip, *port)
}

func helloAction(ctx *fastgo3.Context) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", ctx.Request())

	// Set arbitrary headers
	ctx.SetHeader("Hello-Header", "World-Value")

	// Set cookies
	ctx.SetCookie("HelloCookie", "World !")
}

