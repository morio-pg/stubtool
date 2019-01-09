package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/application/usecase"
	"github.com/morio-pg/stubtool/src/domain/model"
	"github.com/rs/xid"
)

// AdminHandler interface
type AdminHandler interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type adminHandler struct {
	adminUsecase usecase.AdminUsecase
}

// NewAdminHandler return AdminHandler
func NewAdminHandler(adminUsecase usecase.AdminUsecase) AdminHandler {
	return &adminHandler{adminUsecase}
}

func (h *adminHandler) Get(c *gin.Context) {
	stubID := c.Param("stubId")

	stub, err := h.adminUsecase.Get(c, stubID)
	if err != nil {
		h.badRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, stub)
}

func (h *adminHandler) GetAll(c *gin.Context) {
	uid := c.MustGet("uid").(string)

	stubs, err := h.adminUsecase.GetAll(c, uid)
	if err != nil {
		h.badRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, stubs)
}

func (h *adminHandler) Post(c *gin.Context) {
	stubID := xid.New().String()
	uid := c.MustGet("uid").(string)
	filename := c.PostForm("filename")
	description := c.PostForm("description")
	statusCode, _ := strconv.Atoi(c.PostForm("statusCode"))
	wait, _ := strconv.Atoi(c.PostForm("wait"))
	mimeTypeKey, _ := strconv.Atoi(c.PostForm("mimeTypeKey"))
	content := c.PostForm("content")

	stub, err := h.adminUsecase.Post(
		c, stubID, uid, filename, description, statusCode, wait, mimeTypeKey, content)
	if err != nil {
		h.badRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, stub)
}

func (h *adminHandler) Put(c *gin.Context) {
	stubID := c.Param("stubId")
	uid := c.MustGet("uid").(string)
	filename := c.PostForm("filename")
	description := c.PostForm("description")
	statusCode, _ := strconv.Atoi(c.PostForm("statusCode"))
	wait, _ := strconv.Atoi(c.PostForm("wait"))
	mimeTypeKey, _ := strconv.Atoi(c.PostForm("mimeTypeKey"))
	content := c.PostForm("content")

	stub, err := h.adminUsecase.Put(
		c, stubID, uid, filename, description, statusCode, wait, mimeTypeKey, content)
	if err != nil {
		h.badRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, stub)
}

func (h *adminHandler) Delete(c *gin.Context) {
	stubID := c.Param("stubId")
	uid := c.MustGet("uid").(string)

	err := h.adminUsecase.Delete(c, stubID, uid)
	if err != nil {
		h.badRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, model.Success{Status: "OK"})
}

func (h *adminHandler) badRequest(c *gin.Context, err error) {
	fmt.Println(err)
	c.JSON(http.StatusBadRequest, model.Error{Message: http.StatusText(http.StatusBadRequest)})
	c.Abort()
}
