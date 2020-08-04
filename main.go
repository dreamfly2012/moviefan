package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")

        c.Next()
    }
}

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "0.1",
		})
	})

	r.GET("/get", func(c *gin.Context) {
		//查询数据库,随机返回一条数据
		movie := Movie{}
		name := c.Query("name")
		err := movie.get(name)
		if err != nil {
			fmt.Println("查询失败", err)
		}
		fmt.Println(movie)
		//fmt.Println(len(movie))
		fmt.Printf("%v", movie)
		c.JSON(200, gin.H{
			"info": movie,
		})
	})

	r.GET("/random", func(c *gin.Context) {
		//查询数据库,随机返回一条数据
		movie := Movie{}

		movie = movie.random()

		c.JSON(200, gin.H{
			"info": movie,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
