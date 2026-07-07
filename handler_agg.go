package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time>", cmd.Name)
	}
	time_between_reqs := cmd.Args[0]

	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}

	fmt.Printf("Collecting Feeds every %v\n", timeBetweenRequests.String())
	fmt.Println("------------------")

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println(err)
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println(err)
	}

	rss_feed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Println(err)
	}

	for _, item := range rss_feed.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println("")
	}
}
