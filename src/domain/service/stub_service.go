package service

import (
	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/model"
	"github.com/morio-pg/stubtool/src/domain/repository"
	validator "gopkg.in/go-playground/validator.v8"
)

// StubService interface
type StubService interface {
	Save(c *gin.Context, stub *model.Stub) (err error)
}

type stubService struct {
	stubRepository repository.StubRepository
}

// NewStubService return StubService
func NewStubService(stubRepository repository.StubRepository) StubService {
	return &stubService{stubRepository}
}

func (s *stubService) Save(c *gin.Context, stub *model.Stub) (err error) {
	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	if err := validate.Struct(stub); err != nil {
		return err
	}

	if err := s.stubRepository.Save(c, stub); err != nil {
		return err
	}

	return nil
}
