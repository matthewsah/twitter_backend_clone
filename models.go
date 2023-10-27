package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/matthewsah/twitter_backend_clone/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	APIKey    string    `json:"api_key"`
}

type Tweet struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		APIKey:    dbUser.ApiKey,
	}
}

func databaseTweetToTweet(dbTweet database.Tweet) Tweet {
	return Tweet{
		ID:        dbTweet.ID,
		CreatedAt: dbTweet.CreatedAt,
		UpdatedAt: dbTweet.UpdatedAt,
		Content:   dbTweet.Content,
		UserID:    dbTweet.UserID,
	}
}
