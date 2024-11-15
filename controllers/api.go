package controllers

import (
	"RyuLdnWebsite/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Ldn struct {
	TotalGameCount     int `json:"total_game_count"`
	PrivateGameCount   int `json:"private_game_count"`
	PublicGameCount    int `json:"public_game_count"`
	InProgressCount    int `json:"in_progress_count"`
	MasterProxyCount   int `json:"master_proxy_count"`
	TotalPlayerCount   int `json:"total_player_count"`
	PrivatePlayerCount int `json:"private_player_count"`
	PublicPlayerCount  int `json:"public_player_count"`
}

type Game struct {
	ID             string   `json:"id"`
	IsPublic       bool     `json:"is_public"`
	PlayerCount    int      `json:"player_count"`
	MaxPlayerCount int      `json:"max_player_count"`
	GameName       string   `json:"game_name"`
	TitleID        string   `json:"title_id"`
	TitleVersion   string   `json:"title_version"`
	Mode           string   `json:"mode"`
	Status         string   `json:"status"`
	SceneID        int      `json:"scene_id"`
	Players        []string `json:"players"`
}

func GetLDNData(c *gin.Context) {
	var result Ldn
	res, err := services.RedisClient.JSONGet("ldn", ".")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = json.Unmarshal(res.([]byte), &result)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetPublicGames(c *gin.Context) {
	var result []Game
	res, err := services.RedisClient.JSONGet("games", ".")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = json.Unmarshal(res.([]byte), &result)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
