package fastgo3

import (
	"crypto/md5"
	"fmt"
	"time"
)

type MiddlewareFunc func() HandlerFunc

func SeqId() HandlerFunc {
	return func(context *Context) {
		seqId := fmt.Sprintf("%x", md5.New().Sum(nil))
    	context.SeqId = seqId
		context.Next()
	}
}

func Logger() HandlerFunc {
    return func(context *Context) {
    	bTime := time.Now().UnixNano() / int64(time.Millisecond)
		context.Next()

    	eTime := time.Now().UnixNano() / int64(time.Millisecond)
    	fmt.Printf(
    		"%s [%s] %q %d %q %d \n",
			context.RemoteIP(), // remote_addr
			context.Time(), // time_local
			context.RequestURI(), // request
			context.fastHttpRequestCtx.Response.StatusCode(), // status
			context.UserAgent(), // remote_user
			eTime - bTime, // deal time
		)
	}
}
