package makeless_go_authenticator_in_memory

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/makeless/makeless-go/security"
)

func (authenticator *Authenticator) SetMiddleware(middleware *jwt.GinJWTMiddleware) {
	authenticator.GetBaseAuthenticator().SetMiddleware(middleware)
}

func (authenticator *Authenticator) GetMiddleware() *jwt.GinJWTMiddleware {
	return authenticator.GetBaseAuthenticator().GetMiddleware()
}

func (authenticator *Authenticator) GetSecurity() makeless_go_security.Security {
	return nil
}
