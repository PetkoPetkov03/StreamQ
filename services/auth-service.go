package services

import (
	"errors"
	"net/mail"

	steamqsql "github.com/PetkoPetkov/streamq-backend/orm"
	"github.com/PetkoPetkov/streamq-backend/streamqsql/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var authServiceSession AuthService

func GetAuthService() AuthService {
	if authServiceSession == (AuthService{}) {
		authServiceSession = AuthService{
			session: "",
		}
	}
	return authServiceSession
}

type UserAuthReq struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CPassword string `json:"cpassword"`
	Profile   steamqsql.Profile
}

type AuthServiceInterface interface {
	Register(userPrototype UserAuthReq) (err error)
}

type AuthService struct {
	session string
}

func (auth AuthService) Register(ctx *gin.Context, userPrototype UserAuthReq) error {
	if userPrototype.Password != userPrototype.CPassword {
		return errors.New("passwords don't match")
	}

	profile, err := schemas.GetQueryCaller().CreateProfile(ctx, "ROLE_USER")

	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userPrototype.Password), 11)

	if err != nil {
		return err
	}

	if _, err := mail.ParseAddress(userPrototype.Email); err != nil {
		return err
	}

	userBuilder := steamqsql.CreateUserParams{
		Email:     userPrototype.Email,
		Username:  userPrototype.Username,
		Hash:      string(hash),
		Profileid: profile.ID.(int64),
	}

	_, err = schemas.GetQueryCaller().CreateUser(ctx, userBuilder)

	if err != nil {
		return err
	}

	return nil
}
