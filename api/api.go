package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rikikudohust/FilterCandidate/db"
)

type API struct {
  ItemDB *db.ItemDB
}

type Config struct {
	Version string
	Server  *gin.Engine
}

func NewAPI(setup Config) (*API, error) {
  a := &API {}

  v1 := setup.Server.Group("/v1")

  // v1.POST("/skills", a.postListSkill)
  v1.POST("/cvs", a.filterCandidate)
  v1.GET("/cvs", a.getListCV)


  return a, nil
}
