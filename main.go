package main

import (
	"fmt"
	"net/http"
	"webframework"
)

func errorRecover() webframework.HandlerFunc {
	return func(c *webframework.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}

func main() {
	fmt.Println("Understand a website")
	r := webframework.New()
	r.Use(errorRecover())
	r.GET("/", func(c *webframework.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *webframework.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *webframework.Context) {
		c.JSON(http.StatusOK, webframework.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/panic", func(c *webframework.Context) {
		panic("Go panic")
	})

	r.Run(":1234")
}
