package main

import (
	"fmt"
	"main/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	cfg.SetUser("Lydia")
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Println(cfg)
}
