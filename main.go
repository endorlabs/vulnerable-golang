package main

import (
	"fmt"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/owncast/owncast/logging"
        _ "github.com/grafana/grafana-plugin-sdk-go"
)

func main() {
	fmt.Println("Hello world")
}
