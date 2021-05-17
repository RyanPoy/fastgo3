# fastgo3

A simple Golang web framework base on [fasthttp](https://github.com/valyala/fasthttp). 

*fastgo3* does meaning that *Fast Go Go Go*.


## simple usage:
``` Golang
import (
  "github.com/RyanPoy/fastgo3"
)

func main() {
  app := fastgo3.New()
  app.Get("/helloworld", func (ctx *fastgo3.Context) {
    ctx.RenderString("hello, world!")
  })
  app.Run("0.0.0.0", 3030)
}
```

## more examples
* [Basic application](examples/basic_app)
* [Api application](examples/api_app)
* [render application](examples/render_app)
