# fastgo3
a simple golang webframework base on [fasthttp](https://github.com/valyala/fasthttp)

simple usage:
``` Golang
import (
  "github.com/fastgo3/fastgo3"
)

func main() {
  app := fastgo3.NewApplication()
  app.Get("/helloworld", func (ctx *fastgo3.Context) {
    ctx.RenderString("hello, world!")
  })
  app.Run("0.0.0.0", 3030)
}
```

more code examples
* [Basic application](examples/basic_app)
* [Api application](examples/api_app)
* [render application](examples/render_app)
