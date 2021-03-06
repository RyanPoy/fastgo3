package main

import (
	"flag"
	"fmt"
	"github.com/RyanPoy/fastgo3"
)

var (
	ip	= flag.String("ip", "0.0.0.0", "IP or Host")
	port = flag.Int("port", 3030, "Listen Port")
)

func main() {
	flag.Parse()

	app := fastgo3.New()

	app.Get("/basic", basicAction)
	app.Get("/do-get", getHandler)
	app.Post("/do-post", postHandler)
	app.Get("/panic", panicHandler)
	app.Get("/redirect", redirctHandler)
	app.Run(*ip, *port)
}

func basicAction(ctx *fastgo3.Context) {
	fmt.Fprintf(ctx, "Basic Informations\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method)
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
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

func getHandler(ctx *fastgo3.Context) {
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "name=%s\n", ctx.StrParam("name", ""))
}

func postHandler(ctx *fastgo3.Context) {
	fmt.Fprintf(ctx, "Post string is %q\n", ctx.PostArgs())
	fmt.Fprintf(ctx, "password=%s\n", ctx.StrParam("password", ""))
}

func panicHandler(ctx *fastgo3.Context) {
	lst := make([]int, 0)
	lst[0] = 10
}

func redirctHandler(ctx *fastgo3.Context) {
	to := ctx.StrParam("to", "http://www.baidu.com/")
	ctx.Redirect(to)
}
