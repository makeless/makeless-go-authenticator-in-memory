package makeless_go_authenticator_in_memory

import (
	"sync"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/makeless/makeless-go/authenticator"
	"github.com/makeless/makeless-go/model"
	"github.com/makeless/makeless-go/struct"
)

type Authenticator struct {
	BaseAuthenticator makeless_go_authenticator.Authenticator
	Users             []*makeless_go_model.User

	*sync.RWMutex
}

func (authenticator *Authenticator) GetBaseAuthenticator() makeless_go_authenticator.Authenticator {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.BaseAuthenticator
}

func (authenticator *Authenticator) GetUsers() []*makeless_go_model.User {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.Users
}

func (authenticator *Authenticator) AuthenticatorHandler(c *gin.Context) (interface{}, error) {
	var login = &_struct.Login{
		RWMutex: new(sync.RWMutex),
	}

	if err := c.ShouldBind(login); err != nil {
		return nil, err
	}

	for _, user := range authenticator.GetUsers() {
		if *login.GetEmail() == *user.GetEmail() && *login.GetPassword() == *user.GetPassword() {
			return user, nil
		}
	}

	return nil, jwt.ErrFailedAuthentication
}

func (authenticator *Authenticator) CreateMiddleware() error {
	middlware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           authenticator.GetRealm(),
		Key:             authenticator.GetKey(),
		Timeout:         authenticator.GetTimeout(),
		MaxRefresh:      authenticator.GetMaxRefresh(),
		IdentityKey:     authenticator.GetIdentityKey(),
		PayloadFunc:     authenticator.PayloadHandler,
		IdentityHandler: authenticator.IdentityHandler,
		Authenticator:   authenticator.AuthenticatorHandler,
		Authorizator:    authenticator.AuthorizatorHandler,
		Unauthorized:    authenticator.UnauthorizedHandler,
		TimeFunc:        time.Now,
		SendCookie:      true,
		SecureCookie:    false, //non HTTPS dev environments
		CookieHTTPOnly:  true,
		CookieName:      "jwt",
		TokenLookup:     "header:Authorization,cookie:jwt",
	})

	if err != nil {
		return err
	}

	authenticator.SetMiddleware(middlware)
	return nil
}
