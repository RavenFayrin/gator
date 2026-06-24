package main

import (
	"errors"
	"fmt"
	"log"
	"os"

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

func (c *commands) run(s *state, cmd command) error {
	f_cmd, ok := c.registeredCommands[cmd.name]
	if ok {
		return f_cmd(s, cmd)
	} else {
		return errors.New("ERROR: Command not found")
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("ERROR")
	}
	cmd := command{
		name: args[1],
		args: args[2:],
	}
	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
