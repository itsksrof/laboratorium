package main

import (
	"crypto"
	"errors"

	"github.com/golang-jwt/jwt"
)

func main() {
    println("working...")
}

// createToken creates a new token object, specifying signing method and the
// claims you would like it to contain.
func createToken(alg jwt.SigningMethod) *jwt.Token {
    token := jwt.NewWithClaims(alg, jwt.MapClaims{
        "foo": "bar",
    })

    return token
}

// signToken creates and returns a complete and signed JWT. The token is signed
// using the SigningMethod specified in the token.
func signToken(token *jwt.Token, key crypto.PrivateKey) string  {
    tokenStr, err := token.SignedString(key)
    if err != nil {
        return err.Error()
    }

    return tokenStr
}

// parseToken verifies the signature and returns the parsed token. keyFunc will
// receive the parsed token and should return the cryptographic key for verifying
// the signature.
func parseToken(
    tokenStr string, 
    alg jwt.SigningMethod,
    key crypto.PublicKey,
) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        switch alg.(type) {
        case *jwt.SigningMethodRSA:
            if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
                return nil, errors.New("expected *jwt.SigningMethodRSA")
            }
        case *jwt.SigningMethodECDSA:
            if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
                return nil, errors.New("expected *jwt.SigningMethodECDSA")
            }
        case *jwt.SigningMethodEd25519:
            if _, ok := t.Method.(*jwt.SigningMethodEd25519); !ok {
                return nil, errors.New("expected *jwt.SigningMethodEd25519")
            }
        default:
            return nil, errors.New("unexpected signing method")
        }

        return key, nil 
    })
    if err != nil {
        return nil, err
    }

    return token, nil
}

// validateToken 
func validateToken(token *jwt.Token) error {
    if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
        return errors.New("unable to validate token")
    }

    return nil
}
