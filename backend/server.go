package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"path/filepath"
	"runtime"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
    tmpl     *template.Template
    tmplOnce sync.Once
)

type Request struct {
    Length int `json:"length"`
}

type PageVariables struct {
    Password string
    Hash     string
    Match    bool
}

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
    tmpl, err = template.ParseFiles(filepath.Join("../frontend", "index.html"))
    if err != nil {
        panic("Error loading template: " + err.Error())
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method == http.MethodPost {
        var req Request
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            http.Error(w, "Error decoding request", http.StatusBadRequest)
            return
        }

        tmplOnce.Do(initTemplates)

        password := generateRandomPassword(req.Length)
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

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    } else {
        tmplOnce.Do(initTemplates)

        http.FileServer(http.Dir("../frontend")).ServeHTTP(w, r)
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