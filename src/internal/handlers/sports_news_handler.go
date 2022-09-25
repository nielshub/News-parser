package handlers

import (
	"incrowd/src/internal/model"
	"incrowd/src/internal/ports"
	"incrowd/src/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SportNewsHandler struct {
	router           *gin.RouterGroup
	sportNewsService ports.SportNewsService
}

func NewSportNewsHandler(app *gin.RouterGroup, sportNewsService ports.SportNewsService) *SportNewsHandler {
	newsAPI := &SportNewsHandler{sportNewsService: sportNewsService}

	newsRooter := app.Group("/v1/teams/t94")
	newsRooter.GET("/news", newsAPI.getNews)
	newsRooter.GET("/news/:id", newsAPI.getNewsByID)

	newsAPI.router = newsRooter

	return newsAPI
}

func (snh *SportNewsHandler) getNews(c *gin.Context) {
	news, err := snh.sportNewsService.GetNews(c)
	if err != nil {
		log.Logger.Error().Msgf("Error getting news. Error: %s", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error getting news"})
		return
	}

	response := model.NewsResponse{
		Status: "success",
		Data:   news,
		Metadata: model.NewsResponseMetadata{
			CreatedAt:  time.Now(),
			TotalItems: len(news),
			Sort:       "",
		},
	}

	c.JSON(http.StatusOK, response)
}

func (snh *SportNewsHandler) getNewsByID(c *gin.Context) {
	id := c.Query("id")

	newsByID, err := snh.sportNewsService.GetNewsWithID(c, id)
	if err != nil {
		log.Logger.Error().Msgf("Error getting news with ID: "+id+". Error: %s", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error getting news by ID"})
		return
	}

	response := model.NewsByIDResponse{
		Status: "success",
		Data:   *newsByID,
		Metadata: model.NewsResponseMetadata{
			CreatedAt: time.Now(),
		},
	}

	c.JSON(http.StatusOK, response)
}
