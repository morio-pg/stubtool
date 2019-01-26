package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/application/usecase"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/dto"
)

// AccountHandler interface
type AccountHandler interface {
	Delete(c *gin.Context)
}

type accountHandler struct {
	accountUsecase usecase.AccountUsecase
}

// NewAccountHandler return AccountHandler
func NewAccountHandler(accountUsecase usecase.AccountUsecase) AccountHandler {
	return &accountHandler{accountUsecase}
}

func (h *accountHandler) Delete(c *gin.Context) {
	uid := c.MustGet("uid").(string)

	err := h.accountUsecase.Delete(c, uid)
	if err != nil {
		h.badRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Success{Status: "OK"})
}

func (h *accountHandler) badRequest(c *gin.Context, err error) {
	fmt.Println(err)
	c.JSON(http.StatusBadRequest, dto.Error{Message: http.StatusText(http.StatusBadRequest)})
	c.Abort()
}
