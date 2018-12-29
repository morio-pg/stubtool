package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/model"
)

// StubRepository interface
type StubRepository interface {
	Get(c *gin.Context, stubID string) (stub *model.Stub, err error)
	GetAll(c *gin.Context, uid string) (stubs []model.Stub, err error)
	Save(c *gin.Context, stub *model.Stub) (err error)
	Delete(c *gin.Context, stubID string) (err error)
}
