package dao

import (
	"context"
	"fmt"
	"main/dbconfig"
	"main/dto"
)

func CreateGame(game dto.Game) error {
	collection := dbconfig.DATABASE.Collection("Games")

	_, err := collection.InsertOne(context.Background(), game)
	if err != nil {
		fmt.Println("dao error")
		return nil
	}

	return nil
}
