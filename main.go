package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main(){

	//=================GET method=================
	router := gin.Default() 
	// create a default middleware gin router that is for logger and recovery
	//when we come to the logger it will log the request and response
	//when we come to the recovery it will recover from panic
	// in case of server error, gin automatically set the response code 500 which is internal server error and return it to the JSON format

	router.GET("/getData", func(c *gin.Context){
		//here we are using the GET method to get the data from the server and execute the function
		//this function is a handler that contain gin Context, that contain response and request every information about 
		//the request and response included into single Context

		c.JSON(200, gin.H{
			"data": "Hi I am Gin Web Framework",
		})
		//c.Json means an object of response writer that simply return the response in JSON format(whatever in c.JSON)
		//200 is the response code
		//gin.H kind of a interface where we are saying, return this output form of JSON format with 200 response code

	})

	//====================POST method=================
	router.POST("/getDataPost", getDataPost)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)
	
	router.Run(":5000")
	//here gin framework will take default parameter 8080, if you want to give specific port you can define router.Run(":5000") like this

	//=========similar way that we can impliment above code here
	/*
	router := gin.Default()

	router.GET("/getData", getData)
	router.Run()

	func getData(c *gin.Context){
		c.JSON(200, gin.H{
			"data": "Hi I am Gin Web Framework",
		})
	}
		
	
	*/

	



}
//http:localhost:5000/getQueryString?name=deshan&age=23

func getQueryString(c*gin.Context){
	name:=c.Query("name")
	age:=c.Query("age")
	c.JSON(200, gin.H{
		"data": "hi i ma query string method",
		"name":name,
		"age":age,

	})
}

//http:localhost:5000/getUrlData/deshan/23

func getUrlData(c*gin.Context){
	name:=c.Param("name")
	age:=c.Param("age")
	c.JSON(200, gin.H{
		"data": "hi i am url data string method",
		"name":name,
		"age":age,

	})
}

func getDataPost(c *gin.Context){
	body:=c.Request.Body
	value, _  := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		
		"data": "this is a post request",
		"bodyData": string(value),
	})
}
