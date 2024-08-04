package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"runtime"
	"sync"

	"golang.org/x/crypto/bcrypt"
)
 
var (
    tmpl     *template.Template
    tmplOnce sync.Once
)

func generateRandomPassword(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, length)
    for i := range result {
        num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            panic(err)
        }
        result[i] = charset[num.Int64()]
    }
    return string(result)
}

func initTemplates() {
    var err error
    tmpl, err = template.ParseFiles("index.html")
    if err != nil {
        panic("Error loading template: " + err.Error())
    }
}

type PageVariables struct {
    Password string
    Hash     string
    Match    bool
}

func handler(w http.ResponseWriter, r *http.Request) {
    tmplOnce.Do(initTemplates)
    
    password := generateRandomPassword(14)

    hash, err := HashPassword(password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    match := CheckPasswordHash(password, hash)

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
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    port := ":8080"

    http.HandleFunc("/", handler)

    fmt.Println("The server was started on http://localhost:8080")
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Println("Failed to start the server:", err)
    }
}
