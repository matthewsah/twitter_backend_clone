package main

import (
	"fmt"
	"net/http"

	"github.com/matthewsah/twitter_backend_clone/internal/database"
)

// Gets the user's feed
// feed should consist of posts from accounts that the user follows
// in order of descending time
func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	tweets, err := apiCfg.DB.GetFeed(r.Context(), user.ID)
	type tweetDisplay struct {
		Username string
		PostTime string
		Content  string
	}
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get follows: %v", err))
		return
	}

	tweetDisplays := []tweetDisplay{}
	fmt.Printf("@%s's feed\n\n", user.Username)

	for _, tweet := range tweets {
		postTime := tweet.CreatedAt.Time.Format("3:04 PM Jan 02, 2006")
		newTweet := fmt.Sprintf("@%s %s\n%s\n", tweet.Username.String, postTime, tweet.Content.String)
		newTweetDisplay := tweetDisplay{
			Username: tweet.Username.String,
			PostTime: postTime,
			Content:  tweet.Content.String,
		}
		tweetDisplays = append(tweetDisplays, newTweetDisplay)
		fmt.Println(newTweet)
	}

	respondWithJSON(w, 200, tweetDisplays)
}
