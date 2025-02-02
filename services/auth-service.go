package services

import (
	"errors"
	"net/mail"
	"sync"
	"time"

	steamqsql "github.com/PetkoPetkov/streamq-backend/orm"
	"github.com/PetkoPetkov/streamq-backend/streamqsql/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var authServiceSession AuthService

func GetAuthService() AuthService {
	if authServiceSession == (AuthService{}) {
		authServiceSession = AuthService{}
	}
	return authServiceSession
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuthReq struct {
	UserAuth
	Username  string `json:"username"`
	CPassword string `json:"cpassword"`
	Profile   steamqsql.Profile
}

type AuthServiceInterface interface {
	Register(userPrototype UserAuthReq) (err error)
	Login(userPrototype UserAuthReq) (session string, err error)
}

type AuthService struct {
}

type SessionInterface interface {
	Session() string
}

type Session struct {
	value string
}

func (session Session) Session() string {
	return session.value
}

func (auth AuthService) Register(ctx *gin.Context, userPrototype UserAuthReq) error {
	var wg sync.WaitGroup
	var profile int64
	var profileErr error
	var hashErr error
	var hash []byte
	if userPrototype.Password != userPrototype.CPassword {
		return errors.New("passwords don't match")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		profile, profileErr = schemas.GetQueryCaller().CreateProfile(ctx, "ROLE_USER")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		hash, hashErr = bcrypt.GenerateFromPassword([]byte(userPrototype.Password), 11)
	}()

	if _, err := mail.ParseAddress(userPrototype.Email); err != nil {
		return err
	}

	count, qerr := schemas.GetQueryCaller().CheckIfEmailExists(ctx, userPrototype.Email)

	if qerr != nil {
		return qerr
	}

	if count > 0 {
		return errors.New("email already exists")
	}

	wg.Wait()
	if profileErr != nil {
		return profileErr
	}

	if hashErr != nil {
		return hashErr
	}

	userBuilder := steamqsql.CreateUserParams{
		Email:     userPrototype.Email,
		Username:  userPrototype.Username,
		Hash:      string(hash),
		Profileid: profile,
	}

	_, err := schemas.GetQueryCaller().CreateUser(ctx, userBuilder)

	if err != nil {
		return err
	}

	return nil
}

func (auth AuthService) checkPwd(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (auth AuthService) Login(ctx *gin.Context, userPrototype UserAuth) (session *Session, err error) {
	count, err := schemas.GetQueryCaller().CheckIfEmailExists(ctx, userPrototype.Email)
	if err != nil {
		return nil, err
	}

	if count != 1 {
		return nil, errors.New("user dosen't exist")
	}

	user, err := schemas.GetQueryCaller().FetchUserForSession(ctx, userPrototype.Email)

	if err != nil {
		return nil, err
	}

	match := auth.checkPwd(userPrototype.Password, user.Hash)

	if !match {
		return nil, errors.New("password is wrong")
	}

	sessionParams := steamqsql.InitializeSessionParams{
		Token:       uuid.NewString(),
		UserID:      user.ID,
		ProfileID:   user.Profileid,
		CreatedFrom: "",
		CreatedAt:   time.Now().Local(),
		ExpireDate:  time.Now().Local().Add(24 * time.Hour * time.Duration(3)),
	}

	sessionInstance, err := schemas.GetQueryCaller().InitializeSession(ctx, sessionParams)

	if err != nil {
		return nil, err
	}

	session = &Session{
		value: sessionInstance.Token,
	}

	return session, nil
}
