package main

import "github.com/joho/godotenv"

func configLoad() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
