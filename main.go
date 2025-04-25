package main

import (
	"github.com/InspectorGadget/aws-rest/controllers"
	"github.com/InspectorGadget/aws-rest/initializers"
	"github.com/InspectorGadget/aws-rest/middlewares"
	"github.com/gin-gonic/gin"
)

const AWS_REGION string = "ap-southeast-1"

func init() {
	initializers.InitializeSession(AWS_REGION) // Initialize AWS Session
}

func main() {
	r := gin.Default()
	needsAuth := r.Group("/").Use(middlewares.UseAuth)

	r.GET("/", controllers.Index)
	needsAuth.GET("/buckets", controllers.ListBuckets)

	r.Run(":3000")
}
