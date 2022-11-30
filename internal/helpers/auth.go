package helpers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    time.Time
	RtExpires    time.Time
}

func CreateTokenJWT(user_id interface{}) *TokenDetails {
	// create ssid
	suuid, err := uuid.NewV4()
	if err != nil {
		return nil
	}
	rsuuid, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	// create token
	expiary_duration, _ := time.ParseDuration(GoDotEnvVariable("JWT_TOKEN_EXPIRY"))
	token_expiary := time.Now().Add(expiary_duration)
	atClaims := jwt.MapClaims{
		"authorized":   true,
		"dtid":         suuid.String(),
		"expiary_time": token_expiary,
		"user_id":      user_id,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(GoDotEnvVariable("JWT_KEY")))

	// create refresh token
	// get refresh duration
	r_expiary_duration, _ := time.ParseDuration(GoDotEnvVariable("JWT_REFRESH_TOKEN_EXPIRY"))
	refresh_token_expiary := time.Now().Add(r_expiary_duration)
	ratClaims := jwt.MapClaims{
		"authorized":           true,
		"dtid":                 rsuuid.String(),
		"refresh_expiary_time": refresh_token_expiary,
		"user_id":              user_id,
	}
	rat := jwt.NewWithClaims(jwt.SigningMethodHS256, ratClaims)
	refresh_token, _ := rat.SignedString([]byte(GoDotEnvVariable("JWT_KEY")))

	tokenDetails := &TokenDetails{
		AccessToken:  token,
		RefreshToken: refresh_token,
		AccessUuid:   suuid.String(),
		RefreshUuid:  rsuuid.String(),
		AtExpires:    token_expiary,
		RtExpires:    refresh_token_expiary,
	}

	return tokenDetails
}

func RequestTokenJwt(authorizationHeader string) (interface{}, jwt.MapClaims, error) {
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, jwt.MapClaims{}, err
	}

	// get claim token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		suuid := fmt.Sprintf("%v", claims["dtid"])

		// validate token
		if !ok && !token.Valid {
			return nil, jwt.MapClaims{}, err
		}

		return suuid, claims, err
	}

	return nil, jwt.MapClaims{}, errors.New("token is not valid")
}
