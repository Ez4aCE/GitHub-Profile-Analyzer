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

	fmt.Println("1. View GitHub Profile")
	fmt.Println("2. Compare Two GitHub Profile")
	fmt.Print("Choose Option (1/2) :")
	var option int
	fmt.Scanln(&option)

	if option == 1 {
		var username string
		fmt.Print("Enter GitHub Username: ")
		fmt.Scanln(&username)
		user, err := fetchUser(username)
		if err != nil {
			fmt.Println("Error:", err)
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

	} else if option == 2 {
		var username1 string
		var username2 string
		fmt.Print("Enter GitHub Username One: ")
		fmt.Scanln(&username1)
		fmt.Print("Enter GitHub Username Two: ")
		fmt.Scanln(&username2)
		user1, err1 := fetchUser(username1)
		user2, err2 := fetchUser(username2)

		if err1 != nil || err2 != nil {
			fmt.Println("Error:", err1, err2)
			return
		}
		ratio1 := float64(user1.Followers)
		if user1.Following > 0 {
			ratio1 = float64(user1.Followers) / float64(user1.Following)
		}
		ratio2 := float64(user2.Followers)
		if user2.Following > 0 {
			ratio2 = float64(user2.Followers) / float64(user2.Following)
		}
		fmt.Println("📊 GitHub Profile Comparison")
		fmt.Printf("%-15s %-10s %-10s %-10s\n", "User", "Repos", "Followers", "Ratio")
		fmt.Printf("%-15s %-10d %-10d %-10.2f\n", user1.Login, user1.PublicRepos, user1.Followers, ratio1)
		fmt.Printf("%-15s %-10d %-10d %-10.2f\n", user2.Login, user2.PublicRepos, user2.Followers, ratio2)
	} else {
		fmt.Println("Invalid option")
	}

}
func fetchUser(username string) (GitHubUser, error) {
	url := "https://api.github.com/users/" + username

	resp, err := http.Get(url)
	if err != nil {
		return GitHubUser{}, fmt.Errorf("❌ Error fetching Profile : %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {

		return GitHubUser{}, fmt.Errorf("❌ User not found. Check the username")
	} else if resp.StatusCode == 403 {

		return GitHubUser{}, fmt.Errorf("⚠️ Rate limit exceeded. Please wait before trying again")
	} else if resp.StatusCode != 200 {

		return GitHubUser{}, fmt.Errorf("❌ Unexpected error: %s\n", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return GitHubUser{}, fmt.Errorf("❌ Error reading response : %s", err)
	}

	var user GitHubUser
	err = json.Unmarshal(body, &user)

	if err != nil {
		return GitHubUser{}, fmt.Errorf("❌ Error unmarshalling data : %s", err)
	}
	return user, nil
}
