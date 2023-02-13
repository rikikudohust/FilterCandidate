package parsers

import (
	"github.com/gin-gonic/gin"
)

type CVFilter struct {
  CVIndex int `uri:"cvIndex" binding:"required"`
}

func ParserCVFilter(c *gin.Context) (int, error) {
  var cvFilter CVFilter
  if err := c.ShouldBindUri(&cvFilter); err != nil {
    return 0, err
  }

  return  cvFilter.CVIndex, nil
}
