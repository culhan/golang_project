package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Index1(ctx *fasthttp.RequestCtx) {

    firstName, lastName, middlename := "John" , "1", "A'mal"

    fmt.Fprintf(ctx, "hello1, %s %s %s!\n", firstName, middlename, lastName )
	fmt.Fprintf(ctx, "comee123  %s \n", lastName)
 
}

func Hello1(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello1iniloh, %s!\n", ctx.UserValue("name"))
}