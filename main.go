package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
}

func main() {
	url := "https://api.github.com/users/Ez4aCE"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Error fetching data : ", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response : ", err)
	}

	var user GitHubUser
	err = json.Unmarshal(body, &user)

	if err != nil {
		fmt.Println("Error unmarshalling data : ", err)
		return
	}

	fmt.Println("👤 GitHub Profile Info")
	fmt.Printf("Username : %s\n", user.Login)
	fmt.Printf("Name : %s\n", user.Name)
	fmt.Printf("Bio : %s\n", user.Bio)
	fmt.Printf("Public repos : %d\n", user.PublicRepos)
	fmt.Printf("Followers : %d | Following : %d\n", user.Followers, user.Following)

}
