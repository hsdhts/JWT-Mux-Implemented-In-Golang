package authcontroller

import (
	"encoding/json"
	"github.com/jeypc/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Gagal mendecode JSON")
	}

	defer r.Body.Close()

	// hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	//insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("Gagal menyimpan Data")
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
