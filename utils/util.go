package utils

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

// Response ...
type Response struct {
	Code    int         `json:"-"`
	Message string      `json:"-"`
	Payload interface{} `json:"payload"`
}

// AddCors ...
func AddCors(router *mux.Router) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodDelete, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		MaxAge:           10,
		AllowCredentials: true,
	}).Handler(router)
}

// Payload ...
type Payload struct {
	ID  int64 `json:"id"`
	Exp int64 `json:"exp"`
}

const (
	salt = "sjakfslkaf23j213123kjklkjl"
)

// GeneratePasswordHash ...
func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// Send ...
func (res *Response) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)

	if res.Payload == nil && res.Code != http.StatusOK {
		res.Payload = struct {
			Error   bool   `json:"error,omitempty"`
			Message string `json:"message,omitempty"`
		}{
			Error:   true,
			Message: res.Message,
		}
	}
	if len(res.Message) == 0 {
		res.Message = http.StatusText(res.Code)
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("ERROR Sending response failed:", err)
	}
}
