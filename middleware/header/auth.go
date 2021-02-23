package header

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/models"
)

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}

func AuthString() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := AuthHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}

		ah := h.Authorization
		if len(ah) > 5 && strings.ToUpper(ah[0:5]) == "TOKEN" {
			token := ah[6:]
			if token != "" {
				user := models.FindOneByAuthKey(token)
				if user != nil {
					c.Set("User", user)
				}
			}
		}
		c.Next()
	}
}
