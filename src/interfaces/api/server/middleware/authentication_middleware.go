package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/application/usecase"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/dto"
)

// AuthenticationMiddleware interface
type AuthenticationMiddleware interface {
	Check(c *gin.Context)
}

type authenticationMiddleware struct {
	authenticationUsecase usecase.AuthenticationUsecase
}

// NewAuthenticationMiddleware return AuthenticationMiddleware
func NewAuthenticationMiddleware(
	authenticationUsecase usecase.AuthenticationUsecase,
) AuthenticationMiddleware {
	return &authenticationMiddleware{authenticationUsecase}
}

func (m *authenticationMiddleware) Check(c *gin.Context) {
	uid, err := m.authenticationUsecase.Check(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, dto.Error{Message: http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}

	c.Set("uid", uid)
}
