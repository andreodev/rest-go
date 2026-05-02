package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	userModels "rest-go/internal/models/users"
)

func main() {
	req := userModels.UserCreateRequest{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	b, err := json.Marshal(req)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewReader(b))

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		panic("erro to created user")
	}

	var responseApi userModels.UserCreateResponse

	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	fmt.Println("new user created", responseApi)
}
