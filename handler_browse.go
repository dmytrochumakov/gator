package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dmytrochumakov/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		i, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = i
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("Title: %s", post.Title)
	}

	return nil
}
