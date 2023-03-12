package sys

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Uid uint
	jwt.RegisteredClaims
}
