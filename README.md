# go-saas-authenticator-in-memory
Go SaaS Framework - Golang In-Memory Authenticator Implementation

[![Build Status](https://ci.loeffel.io/api/badges/go-saas/go-saas-authenticator-in-memory/status.svg)](https://ci.loeffel.io/go-saas/go-saas-authenticator-in-memory)

```go
package main

import (
    "os"
    "sync"
    "github.com/go-saas/go-saas-authenticator-in-memory"
    "github.com/go-saas/go-saas/authenticator/basic"
    "github.com/go-saas/go-saas/model"
)

func main() {
    email := os.Getenv("AUTH_EMAIL")
    password := os.Getenv("AUTH_PASSWORD")

    authenticator := &go_saas_authenticator_in_memory.Authenticator{
        BaseAuthenticator: &go_saas_authenticator_basic.Authenticator{
            Realm:       "auth",
            Key:         os.Getenv("JWT_KEY"),
            Timeout:     time.Hour,
            MaxRefresh:  time.Hour,
            IdentityKey: "id",
            RWMutex:     new(sync.RWMutex),
        },
        Users: []*go_saas_model.User{
            {
                Model:    go_saas_model.Model{Id: 1},
                Email:    &email,
                Password: &password,
                RWMutex:  new(sync.RWMutex),
            },
        },
        RWMutex: new(sync.RWMutex),
    }
}
```