package satu

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "dua  \n")
}