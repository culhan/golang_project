package main

import (
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/go-redis/redis"

	satu "github.com/user/myProject/app/lib/satu"
	dua "github.com/user/myProject/app/lib/dua"
	token_jwt "github.com/user/myProject/app/lib/token_jwt"
)

func Index(ctx *fasthttp.RequestCtx) {

    firstName, lastName, middlename := "John" , "1", "A'mal"

    fmt.Fprintf(ctx, "hello1, %s %s %s!\n", firstName, middlename, lastName )
	fmt.Fprintf(ctx, "comee123  %s \n", lastName)

	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
	    fmt.Println(value)
	} else {
	    // fmt.Println(ctx,chicken)
	    fmt.Fprintf(ctx,"item is not exists \n")
	}

	dua.Index(ctx)
	satu.Index(ctx)

}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello1, %s!\n", ctx.UserValue("name"))
}

func setToken(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello1, %s!\n", ctx.UserValue("name"))
	token_jwt.SetToken(ctx)
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/w", Index1)
	router.GET("/wew/:name", Hello1)
	router.GET("/settoken", token_jwt.SetToken)
	
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
