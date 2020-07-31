package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	r := gin.Default()
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

		err := movie.get(bson.ObjectIdHex("5f2283b27b046edc980d517d"))
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
