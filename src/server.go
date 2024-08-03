package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomPassword(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    var password []byte
    for i := 0; i < length; i++ {
        password = append(password, charset[rng.Intn(len(charset))])
    }
    return string(password)
}

type PageVariables struct {
    Password string
    Hash     string
    Match    bool
}

func handler(w http.ResponseWriter, r *http.Request) {
    password := generateRandomPassword(12) 

    hash, err := HashPassword(password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    match := CheckPasswordHash(password, hash)

    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
        return
    }

    data := PageVariables{
        Password: password,
        Hash:     hash,
        Match:    match,
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
    }
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    port := ":8080"

    http.HandleFunc("/", handler)

    fmt.Println("The server was started on http://localhost:8080")

    http.ListenAndServe(port, nil)
}
