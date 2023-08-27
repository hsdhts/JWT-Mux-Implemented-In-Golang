package authcontroller

import (
	"encoding/json"
	"github.com/jeypc/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json dari client
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Gagal mendecode json")
	}
	defer r.Body.Close()

	//hashing password with bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	//insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("gagal menyimpan data")
	}
	response, _ := json.Marshal(map[string]string{"message": "success"})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
