package middleware

import (
	"net/http"

	"gorm.io/gorm"

	"bikefest/pkg/model"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		api_key := c.Query("api_key")
		if api_key != "peter12345" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Msg: "還敢偷看歐？",
			})
			return
		}
		c.Next()
	}
}
