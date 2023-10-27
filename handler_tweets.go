package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/matthewsah/twitter_backend_clone/internal/database"
)

// To create a tweet, the request should contain the user's api key in the header
// The request should also contain Content in the JSON body
func (apiCfg *apiConfig) handlerCreateTweet(w http.ResponseWriter, r *http.Request, user database.User) {
	// parameters come from the JSON body
	type parameters struct {
		Content string `json:"content"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	fmt.Println("received params ", params)

	tweet, err := apiCfg.DB.CreateTweet(r.Context(), database.CreateTweetParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Content:   params.Content,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseTweetToTweet(tweet))
}
