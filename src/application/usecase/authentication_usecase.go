package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/repository"
)

// AuthenticationUsecase interface
type AuthenticationUsecase interface {
	Check(c *gin.Context) (uid string, err error)
}

type authenticationUsecase struct {
	accountRepository repository.AccountRepository
}

// NewAuthenticationUsecase return AuthenticationUsecase
func NewAuthenticationUsecase(accountRepository repository.AccountRepository) AuthenticationUsecase {
	return &authenticationUsecase{accountRepository}
}

func (u *authenticationUsecase) Check(c *gin.Context) (uid string, err error) {
	return u.accountRepository.GetUID(c)
}
