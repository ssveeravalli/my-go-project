package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PicnicLocation struct {
	Location     string
	MaxOccupancy float64
	HasMusic     bool
}

// albums slice to seed record album data.
var picnicLocations = []PicnicLocation{
	{Location: "Monroe Park", MaxOccupancy: 6, HasMusic: true},
	{Location: "Golden Gate Bridge", MaxOccupancy: 5, HasMusic: true},
	{Location: "Monroe Park", MaxOccupancy: 4, HasMusic: true},
	{Location: "Monroe Park", MaxOccupancy: 3, HasMusic: false},
	{Location: "Monroe Park", MaxOccupancy: 2, HasMusic: true},
}

// getPicnicLocations responds with the list of all PicnicLocations as JSON.
func getPicnicLocations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, picnicLocations)
}
