package authenticator

import (
	"sync"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/authenticator"
	"github.com/go-saas/go-saas/model"
	"github.com/go-saas/go-saas/struct"
)

type Authenticator struct {
	BaseAuthenticator go_saas_authenticator.Authenticator
	Users             []*go_saas_model.User

	*sync.RWMutex
}

func (authenticator *Authenticator) GetBaseAuthenticator() go_saas_authenticator.Authenticator {
	authenticator.RLock()
	defer authenticator.RUnlock()

	return authenticator.BaseAuthenticator
}

func (authenticator *Authenticator) GetUsers() []*go_saas_model.User {
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
		TokenLookup:     "cookie:jwt",
	})

	if err != nil {
		return err
	}

	authenticator.SetMiddleware(middlware)
	return nil
}
