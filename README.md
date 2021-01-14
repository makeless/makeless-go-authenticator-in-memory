# Makeless - Authenticator In-Memory

Makeless - SaaS Framework - Golang In-Memory Authenticator Implementation

[![Build Status](https://ci.loeffel.io/api/badges/makeless/makeless-go-authenticator-in-memory/status.svg)](https://ci.loeffel.io/makeless/makeless-go-authenticator-in-memory)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

```go
package main

import (
    "time"
    "os"
    "sync"
    "github.com/makeless/makeless-go-authenticator-in-memory"
    "github.com/makeless/makeless-go/authenticator/basic"
    "github.com/makeless/makeless-go/model"
)

func main() {
    email := os.Getenv("AUTH_EMAIL")
    password := os.Getenv("AUTH_PASSWORD")

    authenticator := &makeless_go_authenticator_in_memory.Authenticator{
        BaseAuthenticator: &makeless_go_authenticator_basic.Authenticator{
            Realm:       "auth",
            Key:         os.Getenv("JWT_KEY"),
            Timeout:     time.Hour,
            MaxRefresh:  time.Hour,
            IdentityKey: "id",
            RWMutex:     new(sync.RWMutex),
        },
        Users: []*makeless_go_model.User{
            {
                Model:    makeless_go_model.Model{Id: 1},
                Email:    &email,
                Password: &password,
                RWMutex:  new(sync.RWMutex),
            },
        },
        RWMutex: new(sync.RWMutex),
    }
}
```