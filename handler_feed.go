package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"main/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:			%v\n", feed.ID)
	fmt.Printf(" * Created At:	%v\n", feed.CreatedAt)
	fmt.Printf(" * Updated At:	%v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:		%v\n", feed.Name)
	fmt.Printf(" * Url:			%v\n", feed.Url)
	fmt.Printf(" * UserID:		%v\n", feed.UserID)
}
