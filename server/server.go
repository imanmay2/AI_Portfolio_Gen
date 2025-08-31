package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Welcome to the AI-Portfolio Generator.");


	
	app:=gin.Default();



	app.GET("/data",func(c *gin.Context){
		c.JSON(200,gin.H{"message":"Welcome to the AI Portfolio Generator."});
	})



	app.Run(":8080");
}