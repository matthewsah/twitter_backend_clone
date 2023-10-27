package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/matthewsah/twitter_backend_clone/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		UserToFollowID uuid.UUID `json:"user_to_follow_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	fmt.Println("got params ", params)

	follow, err := apiCfg.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		UserID:         user.ID,
		UserToFollowID: params.UserToFollowID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFollowToFollow(follow))
}

func (apiCfg *apiConfig) handlerGetFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	follows, err := apiCfg.DB.GetFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFollowsToFollows(follows))
}

func (apiCfg *apiConfig) handlerDeleteFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	followIDStr := chi.URLParam(r, "followID")
	followID, err := uuid.Parse(followIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	apiCfg.DB.DeleteFollow(r.Context(), database.DeleteFollowParams{
		UserID:         user.ID,
		UserToFollowID: followID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
