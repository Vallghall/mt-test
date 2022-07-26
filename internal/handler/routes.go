package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	g := gin.Default()
	// handler for adding social facts
	g.POST("/add-fact", h.addFact)
	// handler for retrieving graph os relationships
	g.GET("/relations", h.getRelationsGraph)
	return g
}
