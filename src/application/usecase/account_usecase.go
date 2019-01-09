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
}

// NewAccountUsecase return AccountUsecase
func NewAccountUsecase(accountRepository repository.AccountRepository) AccountUsecase {
	return &accountUsecase{accountRepository}
}

func (u *accountUsecase) Delete(c *gin.Context, uid string) (err error) {
	return u.accountRepository.Delete(c, uid)
}
