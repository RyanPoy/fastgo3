package main

import (
	"github.com/fastgo3/fastgo3"
	"flag"
)

var (
	ip	= flag.String("ip", "0.0.0.0", "IP or Host")
	port = flag.Int("port", 3030, "Listen Port")
)

func main() {
	flag.Parse()

	app := fastgo3.Default()
	app.Get("/ok", OkHandler)
	app.Get("/err", ErrHandler)
	app.Get("/finish", FinishHandler)
	app.Run(*ip, *port)
}

func OkHandler(ctx *fastgo3.Context) {
	user := make(map[string]string)
	user["name"] = "瑞安"
	user["password"] = "123we!@#QWE"
	ctx.Ok(user)
}

func ErrHandler(ctx *fastgo3.Context) {
	ctx.Err("This is an error message ！")
}

func FinishHandler(ctx *fastgo3.Context) {
	user := make(map[string]string)
	user["name"] = "瑞安"
	user["password"] = "123we!@#QWE"
	ctx.Finish(0, "获取用户信息成功", user)
}
