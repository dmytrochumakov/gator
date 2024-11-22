package main

import (
	"context"
	"fmt"
)

func handlerAggregator(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed) // Prints: {Name:Alice Age:30}
	return nil
}
