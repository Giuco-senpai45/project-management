package controllers

import (
	"go-project/helpers"
	"go-project/middleware"
	"go-project/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	all, err := userModel.GetAllUsers()
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, all)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	var updatedUserBody models.User
	err = helpers.ReadJSON(w, r, &updatedUserBody)
	updatedUserBody.ID = uint(id)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	updatedUser, err := userModel.UpdateUser(&updatedUserBody)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, updatedUser)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUserBody models.User
	err := helpers.ReadJSON(w, r, &newUserBody)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserBody.Password), 10)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	newUserBody.Password = string(hashedPassword)

	user, err := userModel.CreateUser(&newUserBody)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	var newUserBody models.User
	err := helpers.ReadJSON(w, r, &newUserBody)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	password := newUserBody.Password

	user, err := userModel.LoginUser(email)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600 * 24 * 30, // 30 days
	}

	// Set the cookie in the response
	http.SetCookie(w, cookie)
	helpers.WriteJSON(w, http.StatusOK, user)
}

func GetAllUserProjects(w http.ResponseWriter, r *http.Request) {
	user, err := middleware.RequireAuth(w, r)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	projects, err := userModel.GetAllUserProjects(user)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, projects)
}
