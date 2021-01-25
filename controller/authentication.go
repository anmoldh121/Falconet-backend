package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/chatApp/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthRequest struct {
	Ph string `bson:"Ph,omitempty"`
}
type AuthResponse struct {
	Token  string
	PeerId interface{}
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
		return []byte("secretKey"), nil
	})
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func Register(c echo.Context, db *mongo.Database) (AuthResponse, error) {
	req := c.Request().Body
	var peerId interface{}
	json_map := make(map[string]string)
	err := json.NewDecoder(req).Decode(&json_map)
	if err != nil {
		return AuthResponse{}, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.Collection("peers")
	filter := bson.D{{"ph", json_map["Ph"]}}
	update := bson.D{{"$set", bson.D{{"ph", json_map["Ph"]}}}}
	opts := options.Update().SetUpsert(true)
	insertResult, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return AuthResponse{}, err
	}
	var results models.Peer
	peerId = insertResult.UpsertedID
	if insertResult.UpsertedID == nil {
		err := collection.FindOne(ctx, bson.D{{"ph", json_map["Ph"]}}).Decode(&results)
		if err != nil {
			return AuthResponse{}, err
		}
		peerId = results.PeerId
	}
	token, _ := CreateToken(json_map["Ph"])
	r := AuthResponse{Token: token, PeerId: peerId}
	return r, nil
}
