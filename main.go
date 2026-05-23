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
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanln(&username)
	url := "https://api.github.com/users/" + username

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Error fetching data : ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("❌ User not found. Check the username.")
		return
	} else if resp.StatusCode == 403 {
		fmt.Println("⚠️ Rate limit exceeded. Please wait before trying again.")
		return
	} else if resp.StatusCode != 200 {
		fmt.Printf("❌ Unexpected error: %s\n", resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("❌ Error reading response : ", err)
	}

	var user GitHubUser
	err = json.Unmarshal(body, &user)

	if err != nil {
		fmt.Println("❌ Error unmarshalling data : ", err)
		return
	}

	var ratio float64
	if user.Following > 0 {
		ratio = float64(user.Followers) / float64(user.Following)
	} else {
		ratio = float64(user.Followers)
	}

	var followerPerRepo float64
	if user.Followers > 0 {
		followerPerRepo = float64(user.Followers) / float64(user.PublicRepos)
	}

	fmt.Println("👤 GitHub Profile Info")
	fmt.Printf("Username : %s\n", user.Login)
	fmt.Printf("Name : %s\n", user.Name)
	fmt.Printf("Bio : %s\n", user.Bio)
	fmt.Printf("Public repos : %d\n", user.PublicRepos)
	fmt.Printf("Followers : %d | Following : %d\n", user.Followers, user.Following)
	fmt.Printf("Follower-to-Following Ratio: %.2f\n", ratio)
	fmt.Printf("Average Followers per Repo: %.4f\n", followerPerRepo)
}
