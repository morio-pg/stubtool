package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/repository"
)

// AccountUsecase interface
type AccountUsecase interface {
	Delete(c *gin.Context, uid string) (err error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
	stubRepository    repository.StubRepository
}

// NewAccountUsecase return AccountUsecase
func NewAccountUsecase(
	accountRepository repository.AccountRepository,
	stubRepository repository.StubRepository,
) AccountUsecase {
	return &accountUsecase{accountRepository, stubRepository}
}

func (u *accountUsecase) Delete(c *gin.Context, uid string) (err error) {
	if err := u.stubRepository.DeleteAll(c, uid); err != nil {
		return err
	}

	return u.accountRepository.Delete(c, uid)
}
