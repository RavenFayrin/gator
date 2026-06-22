package main

import (
	"errors"
	"fmt"
	"log"

	"main/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("ERROR: Login expects username")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to s%\n", cmd.args[0])
	return nil
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("lydia")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
