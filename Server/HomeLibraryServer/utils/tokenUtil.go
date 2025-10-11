package utils

import (
	"context"
	"os"
	"time"
	"errors"

	"github.com/Sandeshpd/home-library/Server/HomeLibraryServer/database"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SignedDetails struct {
	Email  string
	Name   string
	Role   string
	UserID string
	jwt.RegisteredClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")
var SECRET_REFRESH_KEY string = os.Getenv("SECRET_REFRESH_KEY")

// Generate Token and Refresh Token
func GenerateAllTokens(email, name, role, userId string) (string, string, error) {
	claims := &SignedDetails{
		Email:  email,
		Name:   name,
		Role:   role,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "HomeLibrary",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", "", err
	}

	refreshClaims := &SignedDetails{
		Email:  email,
		Name:   name,
		Role:   role,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "HomeLibrary",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(SECRET_REFRESH_KEY))

	if err != nil {
		return "", "", err
	}

	return signedToken, signedRefreshToken, nil
}

// Update token if expired
func UpdateAllTokens(userId, token, refreshToken string, client *mongo.Client) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateData := bson.M{
		"$set": bson.M{
			"token": token,
			"refresh_token": refreshToken,
			"update_at": updateAt,
		},
	}
	var userCollection *mongo.Collection = database.OpenCollection("users", client)
	_, err = userCollection.UpdateOne(ctx, bson.M{"user_id": userId}, updateData)

	if err != nil {
		return err
	}

	return nil
}

// Extract the token for authorization purposes
func GetAccessToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorzation header is required.")
	}

	tokenString := authHeader[len("Bearer "):]
	if tokenString == "" {
		return "", errors.New("Bearer token is required.")
	}

	return tokenString, nil
}

// Validate the token if it is right or wrong
func ValidateToken(tokenString string) (*SignedDetails, error) {
	claims := &SignedDetails{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("Token has expired.")
	}

	return claims, nil
}