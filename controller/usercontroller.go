package controller

import (
	"encoding/json"
	"github.com/WytQuant/bunframework/cookie"
	"github.com/WytQuant/bunframework/models"
	"github.com/uptrace/bunrouter"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type userData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register(w http.ResponseWriter, r bunrouter.Request) error {
	var data userData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": err.Error(),
		})
	}

	password, bcErr := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if bcErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": bcErr.Error(),
		})
	}

	user := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  string(password),
	}

	if err := models.CreateUser(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": err.Error(),
		})
	}

	return bunrouter.JSON(w, bunrouter.H{
		"message": "register",
	})
}

func Login(w http.ResponseWriter, r bunrouter.Request) error {
	var data userData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": err.Error(),
		})
	}

	var user models.User
	if !models.CheckEmail(data.Email, &user) {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "not found user",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "not authorized",
		})
	}

	sess, sessErr := cookie.Store.Get(r.Request, "cookie")
	if sessErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": sessErr.Error(),
		})
	}

	sess.Values[cookie.AuthKey] = true
	sess.Values[cookie.UserId] = user.ID
	sess.Save(r.Request, w)

	return bunrouter.JSON(w, bunrouter.H{
		"message": "Logged in",
	})
}

func Logout(w http.ResponseWriter, r bunrouter.Request) error {
	sess, err := cookie.Store.Get(r.Request, "cookie")
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "Logged out with no session",
		})
	}

	sess.Options.MaxAge = -1
	sess.Save(r.Request, w)
	return bunrouter.JSON(w, bunrouter.H{
		"message": "Logged out successfully",
	})
}

func GetUser(w http.ResponseWriter, r bunrouter.Request) error {
	sess, err := cookie.Store.Get(r.Request, "cookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "not authorized",
		})
	}

	if sess.Values[cookie.AuthKey] == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "not authorized, authkey equal nil",
		})
	}

	userId := sess.Values[cookie.UserId]
	if userId == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "not authorized",
		})
	}

	var user *models.User
	user, err = models.GetUser(userId.(int))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "not authorized",
		})
	}

	user.Password = "Credential"

	return bunrouter.JSON(w, user)
}
