package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	cfg := Config{}

	filepath, err := get_home_dir()
	if err != nil {
		return cfg, err
	}
	file, err := os.Open(filepath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func get_home_dir() (string, error) {
	url, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	full_url := filepath.Join(url, ".gatorconfig.json")
	return full_url, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func write(cfg Config) error {
	filepath, err := get_home_dir()
	if err != nil {
		return err
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&cfg)
	if err != nil {
		return err
	}
	return nil
}
