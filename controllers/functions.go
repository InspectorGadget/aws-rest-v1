package controllers

import (
	"net/http"

	"github.com/InspectorGadget/aws-rest/initializers"
	"github.com/InspectorGadget/aws-rest/structs"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		structs.APIResponse{"message": "Pong!"},
	)
}

func ListBuckets(c *gin.Context) {
	s3 := s3.New(initializers.AwsSession)
	results, err := s3.ListBuckets(nil)
	if err != nil {
		panic(err.Error())
	}

	buckets := []string{}
	for _, bucket := range results.Buckets {
		buckets = append(buckets, *bucket.Name)
	}

	c.JSON(
		http.StatusOK,
		structs.APIResponse{
			"total":   len(buckets),
			"buckets": buckets,
		},
	)
}
