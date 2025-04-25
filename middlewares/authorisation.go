package middlewares

import (
	"net/http"

	"github.com/InspectorGadget/aws-rest/structs"
	"github.com/gin-gonic/gin"
)

const API_KEY string = "thisisatestkey"

func UseAuth(c *gin.Context) {
	authorizationValue := c.Request.Header.Get("Authorization")
	if authorizationValue != API_KEY {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			structs.APIResponse{
				"message": "You are not authenticated!",
			},
		)
	}

	// Set header (Just for fun!)
	c.Header("authenticated", "true")

	// Pass to the next function
	c.Next()
}
