package firebase

import (
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/model"
	"github.com/morio-pg/stubtool/src/domain/repository"
)

type accountRepository struct{}

// NewAccountRepository return repository.AccountRepository
func NewAccountRepository() repository.AccountRepository {
	return &accountRepository{}
}

func (r *accountRepository) GetUID(c *gin.Context) (uid string, err error) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" || authorizationHeader[:7] != "Bearer " {
		return "", fmt.Errorf(model.ErrAuthorizationHeader, "idToken is invalid")
	}
	idToken := authorizationHeader[7:]

	app, err := firebase.NewApp(c, nil, opt)
	if err != nil {
		return "", fmt.Errorf(model.ErrInitializingApp, err)
	}

	client, err := app.Auth(c)
	if err != nil {
		return "", fmt.Errorf(model.ErrGettingAuthClient, err)
	}

	token, err := client.VerifyIDToken(c, idToken)
	if err != nil {
		return "", fmt.Errorf(model.ErrVerifyingIDToken, err)
	}

	return token.UID, nil
}

func (r *accountRepository) Delete(c *gin.Context, uid string) (err error) {
	app, err := firebase.NewApp(c, nil, opt)
	if err != nil {
		return fmt.Errorf(model.ErrInitializingApp, err)
	}

	client, err := app.Auth(c)
	if err != nil {
		return fmt.Errorf(model.ErrGettingAuthClient, err)
	}

	if err = client.DeleteUser(c, uid); err != nil {
		return fmt.Errorf(model.ErrDeletingUser, err)
	}

	return nil
}
