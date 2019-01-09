package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/model"
	"github.com/morio-pg/stubtool/src/domain/repository"
)

// StubUsecase interface
type StubUsecase interface {
	Get(c *gin.Context, stubID string, filename string) (stub *model.Stub, err error)
}

type stubUsecase struct {
	stubRepository repository.StubRepository
}

// NewStubUsecase return StubUsecase
func NewStubUsecase(stubRepository repository.StubRepository) StubUsecase {
	return &stubUsecase{stubRepository}
}

func (u *stubUsecase) Get(c *gin.Context, stubID string, filename string) (stub *model.Stub, err error) {
	stub, err = u.stubRepository.Get(c, stubID)
	if err != nil {
		return nil, err
	}
	if stub.Filename != filename {
		return nil, fmt.Errorf(model.ErrMismatchFilename, stub.Filename, filename)
	}

	return stub, nil
}
