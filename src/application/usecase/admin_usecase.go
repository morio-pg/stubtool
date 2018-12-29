package usecase

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/model"
	"github.com/morio-pg/stubtool/src/domain/repository"
	"github.com/morio-pg/stubtool/src/domain/service"
)

// AdminUsecase interface
type AdminUsecase interface {
	Get(c *gin.Context, stubID string) (stub *model.Stub, err error)
	GetAll(c *gin.Context, uid string) (stub []model.Stub, err error)
	Post(c *gin.Context, stubID string, uid string, filename string, description string,
		statusCode int, wait int, MimeTypeKey int, Content string) (stub *model.Stub, err error)
	Put(c *gin.Context, stubID string, uid string, filename string, description string,
		statusCode int, wait int, MimeTypeKey int, Content string) (stub *model.Stub, err error)
	Delete(c *gin.Context, stubID string, uid string) (err error)
}

type adminUsecase struct {
	stubRepository repository.StubRepository
	stubService    service.StubService
}

// NewAdminUsecase return AdminUsecase
func NewAdminUsecase(stubRepository repository.StubRepository, stubService service.StubService) AdminUsecase {
	return &adminUsecase{stubRepository, stubService}
}

func (u *adminUsecase) Get(c *gin.Context, stubID string) (stub *model.Stub, err error) {
	return u.stubRepository.Get(c, stubID)
}

func (u *adminUsecase) GetAll(c *gin.Context, uid string) (stub []model.Stub, err error) {
	return u.stubRepository.GetAll(c, uid)
}

func (u *adminUsecase) Post(c *gin.Context, stubID string, uid string, filename string, description string,
	statusCode int, wait int, mimeTypeKey int, content string) (stub *model.Stub, err error) {
	stub = &model.Stub{
		StubID:      stubID,
		UID:         uid,
		Filename:    filename,
		Description: description,
		StatusCode:  statusCode,
		Wait:        wait,
		MimeTypeKey: mimeTypeKey,
		Content:     content,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err = u.stubService.Save(c, stub); err != nil {
		return nil, err
	}

	return stub, nil
}

func (u *adminUsecase) Put(c *gin.Context, stubID string, uid string, filename string, description string,
	statusCode int, wait int, mimeTypeKey int, content string) (stub *model.Stub, err error) {
	stub, err = u.stubRepository.Get(c, stubID)
	if err != nil {
		return nil, err
	}
	if stub.UID != uid {
		return nil, fmt.Errorf(model.ErrMismatchUID, stub.UID, uid)
	}

	stub.Filename = filename
	stub.Description = description
	stub.StatusCode = statusCode
	stub.Wait = wait
	stub.MimeTypeKey = mimeTypeKey
	stub.Content = content
	stub.UpdatedAt = time.Now()

	if err = u.stubService.Save(c, stub); err != nil {
		return nil, err
	}

	return stub, nil
}

func (u *adminUsecase) Delete(c *gin.Context, stubID string, uid string) (err error) {
	stub, err := u.stubRepository.Get(c, stubID)
	if err != nil {
		return err
	}
	if stub.UID != uid {
		return fmt.Errorf(model.ErrMismatchUID, stub.UID, uid)
	}

	if err = u.stubRepository.Delete(c, stubID); err != nil {
		return err
	}

	return nil
}
