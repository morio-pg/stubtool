package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/application/usecase"
	"github.com/morio-pg/stubtool/src/domain/model"
)

// StubsHandler interface
type StubsHandler interface {
	Get(c *gin.Context)
}

type stubsHandler struct {
	stubUsecase usecase.StubUsecase
}

// NewStubsHandler return StubsHandler
func NewStubsHandler(stubUsecase usecase.StubUsecase) StubsHandler {
	return &stubsHandler{stubUsecase}
}

func (h *stubsHandler) Get(c *gin.Context) {
	startTime := time.Now().UnixNano()
	stubID := c.Param("stubId")
	filename := c.Param("filename")

	stub, err := h.stubUsecase.Get(c, stubID, filename)
	if err != nil {
		h.notFound(c, err)
		return
	}

	// 掛った時間をナノ秒からミリ秒に変換
	wait := stub.Wait - int(((time.Now().UnixNano() - startTime) / 1000000))
	if wait < 0 {
		wait = 0
	}
	fmt.Println(wait)
	time.Sleep(time.Duration(wait) * time.Millisecond)

	c.Data(stub.StatusCode, stub.MimeType(), []byte(stub.Content))
}

func (h *stubsHandler) notFound(c *gin.Context, err error) {
	fmt.Println(err)
	c.JSON(http.StatusNotFound, model.Error{Message: http.StatusText(http.StatusNotFound)})
	c.Abort()
}
