package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AuthRequest struct {
	Ph string
}
type AuthResponse struct {
	Token string
}

func CreateToken(peerId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"peer_id":    peerId,
	})
	jwtToken, err := token.SignedString([]byte("scretKey"))
	if err != nil {
		return "", nil
	}
	return jwtToken, nil
}

func ExtractToken(token string) string {
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VarifyToken(token string) (*jwt.Token, error) {
	tokenString := ExtractToken(token)
	tok, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error while parsing token")
		}
		return []byte(os.Getenv("Access_Secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func Register(c echo.Context) ([]byte, error) {
	req := c.Request().Body
	json_map := make(map[string]string)
	err := json.NewDecoder(req).Decode(&json_map)
	if err != nil {
		return nil, err
	}
	token, _ := CreateToken(json_map["Ph"])
	r := AuthResponse{Token: token}
	res, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return res, nil
}
