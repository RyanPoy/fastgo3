package main

import (
	"github.com/fastgo3/fastgo3"
	"flag"
	"fmt"
)

var (
	ip	= flag.String("ip", "0.0.0.0", "IP or Host")
	port = flag.Int("port", 3030, "Listen Port")
)

func main() {
	flag.Parse()
	app := fastgo3.New()
	app.Get("/user.json", JsonHandler)
	app.Get("/user.html", HtmlHandler)
	app.Get("/user.string", StringHandler)
	app.Run(*ip, *port)
}

func JsonHandler(ctx *fastgo3.Context) {
	user := make(map[string]string)
	user["name"] = "瑞安"
	user["password"] = "123we!@#QWE"
	ctx.RenderJson(user)
}

func HtmlHandler(ctx *fastgo3.Context) {
	ctx.RenderHtml(`
<html>
  <head></head>
  <body>
    <ol>
	  <ot>name：</ot>
      <ot>瑞安</ot>
	</ol>
	<ol>
	  <ot>password：</ot>
      <ot>123qwe！@#QWE</ot>
	</ol>
  </body>
</html>
    `)
}

func StringHandler(ctx *fastgo3.Context) {
	ctx.RenderString(fmt.Sprintf("用户名：%s\n密码：%s", "瑞安", "123qwe！@#QWE"))
}
