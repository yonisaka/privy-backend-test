package middleware

import (
	"errors"
	"net/http"
	"privy-backend-test/internal/exceptions"
	"privy-backend-test/internal/helpers"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check authorization is not empty
		if c.Request.Header["Authorization"] == nil {
			err := errors.New("Not Authorized")
			errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
			c.SecureJSON(errorState.Code, errorState)
			c.Abort()
			return
		}

		// check token container
		authorizationHeader := c.Request.Header["Authorization"]
		if !strings.Contains(authorizationHeader[0], "Bearer") {
			err := errors.New("Invalid Token")
			errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
			c.SecureJSON(errorState.Code, errorState)
			c.Abort()
			return
		}

		// check uuid is not empty and not revoked
		tokenString := strings.Replace(authorizationHeader[0], "Bearer ", "", -1)
		_, claims, err := helpers.RequestTokenJwt(tokenString)
		if err != nil {
			err := errors.New("Not Authorized")
			errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
			c.SecureJSON(errorState.Code, errorState)
			c.Abort()
			return
		}

		if len(claims) == 0 {
			err := errors.New("Not Authorized")
			errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
			c.SecureJSON(errorState.Code, errorState)
			c.Abort()
			return
		}

		tm, err := time.Parse(time.RFC3339, claims["expiary_time"].(string))
		if err != nil {
			errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
			c.SecureJSON(errorState.Code, errorState)
			c.Abort()
			return
		}

		if tm.Before(time.Now()) {
			err := errors.New("Expired Token")
			errorState := exceptions.ErrorException(http.StatusNotAcceptable, err.Error())
			c.SecureJSON(errorState.Code, errorState)
			c.Abort()
			return
		}
	}
}
