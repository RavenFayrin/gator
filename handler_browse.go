package main

import (
	"context"
	"fmt"
	"main/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	var err error

	if len(cmd.Args) != 1 {
		limit = 2
	} else {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.PublishedAt.Time)
		fmt.Println("")
		fmt.Println(post.Description.String)
		fmt.Println("----------------------------------")
	}

	return nil
}
