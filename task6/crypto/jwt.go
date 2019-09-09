package crypto

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"../config"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth struct {
	Key          []byte
	Lifetime     time.Duration
	ClientSecret string
	Issuer       string
	Audience     string
}

type JWTClaim struct {
	Purpose string `json:"purpose"`
	jwt.StandardClaims
}

var auth *JWTAuth

func initJwtKey() {
	auth = &JWTAuth{
		Key:          []byte(config.MustGetString("auth.jwt_key")),
		Lifetime:     time.Duration(config.MustGetInt("auth.jwt_lifetime")) * time.Minute,
		ClientSecret: config.MustGetString("auth.client_secret"),
		Issuer:       config.MustGetString("auth.jwt_issuer"),
		Audience:     config.MustGetString("auth.jwt_audience"),
	}
}

func (j *JWTAuth) New(subject string, args ...string) (token string, expiredAt int64, err error) {
	// Get purpose if exist
	var purpose string
	if len(args) > 0 {
		purpose = args[0]
	}
	// Get lifetime if exist
	var lifetime time.Duration
	if len(args) > 1 {
		tmp, err := strconv.Atoi(args[1])
		if err != nil {
			// Set lifetime to two weeks
			tmp = 60
		}
		lifetime = time.Duration(tmp) * time.Minute
	} else {
		lifetime = j.Lifetime
	}
	// Initiate current timestamp and expire time
	t := time.Now().UTC()
	expiredAt = t.Add(lifetime).Unix()
	// Initiate access token claims
	claims := &JWTClaim{
		Purpose: purpose,
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			Issuer:    j.Issuer,
			Audience:  j.Audience,
			IssuedAt:  t.Unix(),
			ExpiresAt: expiredAt,
		},
	}
	// claims token
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = payload.SignedString(j.Key)
	if err != nil {
		log.Error(err)
		return "", 0, err
	}
	// Return access token and expire time
	return token, expiredAt, nil
}

func (j *JWTAuth) Verify(input string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(input, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("auth: unexpected signing method: %v", token.Header["alg"])
		}
		return j.Key, nil
	})
	// Check parsing err
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// Check claim
	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("auth: invalid token claims")
	}
	return claims, nil
}

func NewJWT(subject string, args ...string) (token string, expiredAt int64, err error) {
	return auth.New(subject, args...)
}

// NewJWTPurpose generates jwt token that must have subject, purpose and lifetime
func NewJWTPurpose(subject string, purpose string, lifetime int) (token string, expiredAt int64, err error) {
	return auth.New(subject, purpose, strconv.Itoa(lifetime))
}

// NewJWTAnonymous generates jwt token without subject and must have purpose and lifetime
func NewJWTAnonymous(purpose string, lifetime int) (token string, expiredAt int64, err error) {
	return auth.New("", purpose, strconv.Itoa(lifetime))
}

func VerifyJWT(input string) (*JWTClaim, error) {
	return auth.Verify(input)
}

// @todo replace with more secure client secret validation
func ValidateSecret(clientSecret string) bool {
	if clientSecret != auth.ClientSecret {
		return false
	}
	return true
}
