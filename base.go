package authenticator

import (
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-saas/go-saas/security"
)

func (authenticator *Authenticator) SetMiddleware(middleware *jwt.GinJWTMiddleware) {
	authenticator.GetBaseAuthenticator().SetMiddleware(middleware)
}

func (authenticator *Authenticator) GetMiddleware() *jwt.GinJWTMiddleware {
	return authenticator.GetBaseAuthenticator().GetMiddleware()
}

func (authenticator *Authenticator) GetSecurity() go_saas_security.Security {
	return nil
}

func (authenticator *Authenticator) GetRealm() string {
	return authenticator.GetBaseAuthenticator().GetRealm()
}

func (authenticator *Authenticator) GetKey() []byte {
	return authenticator.GetBaseAuthenticator().GetKey()
}

func (authenticator *Authenticator) GetTimeout() time.Duration {
	return authenticator.GetBaseAuthenticator().GetTimeout()
}

func (authenticator *Authenticator) GetMaxRefresh() time.Duration {
	return authenticator.GetBaseAuthenticator().GetMaxRefresh()
}

func (authenticator *Authenticator) GetIdentityKey() string {
	return authenticator.GetBaseAuthenticator().GetIdentityKey()
}

func (authenticator *Authenticator) PayloadHandler(data interface{}) jwt.MapClaims {
	return authenticator.GetBaseAuthenticator().PayloadHandler(data)
}

func (authenticator *Authenticator) IdentityHandler(c *gin.Context) interface{} {
	return authenticator.GetBaseAuthenticator().IdentityHandler(c)
}

func (authenticator *Authenticator) AuthorizatorHandler(data interface{}, c *gin.Context) bool {
	return authenticator.GetBaseAuthenticator().AuthorizatorHandler(data, c)
}

func (authenticator *Authenticator) UnauthorizedHandler(c *gin.Context, code int, message string) {
	authenticator.GetBaseAuthenticator().UnauthorizedHandler(c, code, message)
}

func (authenticator *Authenticator) GetAuthUserId(c *gin.Context) uint {
	return authenticator.GetBaseAuthenticator().GetAuthUserId(c)
}
