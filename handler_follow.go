package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dmytrochumakov/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	printFeedNameAndCurrentUsername(feed, user)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feedFollowed := range feedsFollowed {
		fmt.Printf("* Name:          %s\n", feedFollowed.FeedName)
	}
	return nil
}

func printFeedNameAndCurrentUsername(feed database.Feed, user database.User) {
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* User:       	 %s\n", user.Name)
}
