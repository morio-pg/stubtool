package repository

import "github.com/gin-gonic/gin"

// AccountRepository interface
type AccountRepository interface {
	GetUID(c *gin.Context) (uid string, err error)
	Delete(c *gin.Context, uid string) (err error)
}
