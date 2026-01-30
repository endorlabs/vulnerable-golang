package main

import (
	"fmt"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/google/uuid"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/unrolled/secure"
	_ "google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world")
}
