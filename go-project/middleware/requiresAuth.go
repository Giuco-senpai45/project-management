package middleware

import (
	"fmt"
	"go-project/helpers"
	"go-project/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		return nil, err
	}

	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			helpers.ErrorJSON(w, err, http.StatusUnauthorized)
			return nil, err

		} else {
			var userModelInterface models.User
			user, err := userModelInterface.GetUserByID(int64(claims["sub"].(float64)))
			if err != nil {
				helpers.ErrorJSON(w, err, http.StatusUnauthorized)
				return nil, err
			}

			return user, nil
		}
	} else {
		helpers.ErrorJSON(w, err, http.StatusUnauthorized)
		return nil, err
	}
}
