package config

import "github.com/joho/godotenv"

func Init() error {
	err := godotenv.Load("../../app.env")
	if err != nil {
		return err
	}
	return nil
}
