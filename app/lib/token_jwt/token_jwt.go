package token_jwt

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type Key int

const MyKey Key = 123123213

// JWT schema of the data it will store
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// create a JWT and put in the clients cookie
func SetToken(ctx *fasthttp.RequestCtx) {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	// expireCookie := time.Now().Add(time.Hour * 1)

	claims := Claims{
		"myusername",
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	// cookie := fasthttp.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	// fasthttp.SetCookie(res, &cookie)

	// http.Redirect(res, req, "/profile", 307)
	fmt.Fprintf(ctx, "%s", signedToken)
	validate(ctx)
}

// middleware to protect private pages
func validate(ctx *fasthttp.RequestCtx) {

	// ctx.Request.Header.Peek("Authorization")
	header := ctx.Request.Header.Peek("Authorization") 
	fmt.Fprintf(ctx, "%s", header)
	// return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	// 	cookie, err := req.Cookie("Auth")
	// 	if err != nil {
	// 		http.NotFound(res, req)
	// 		return
	// 	}

	// 	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
	// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 			return nil, fmt.Errorf("Unexpected signing method")
	// 		}
	// 		return []byte("secret"), nil
	// 	})
	// 	if err != nil {
	// 		http.NotFound(res, req)
	// 		return
	// 	}

	// 	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
	// 		ctx := context.WithValue(req.Context(), MyKey, *claims)
	// 		page(res, req.WithContext(ctx))
	// 	} else {
	// 		http.NotFound(res, req)
	// 		return
	// 	}
	// })
}