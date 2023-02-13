package api

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/rikikudohust/FilterCandidate/common"
)

func (a *API) getListSkill(c *gin.Context) {
}

func (a *API) postListSkill(c *gin.Context) {

	type okResponse struct {
		Success string `json:"success"`
	}
	// Return OK
	c.JSON(http.StatusOK, &okResponse{
		Success: "OK",
	})

}

// type receiveSkil struct {
// 	id         int    `json:"id" binding:"required"`
// 	skill_name string `json:"skill_name binding:"required"`
// }
//
// type receiveListSkill struct {
//   skills []receiveSkil `binding:"required"`
// }
//
// func skillAPIToCommon(skillData *receiveSkil) *common.PersonalSkill {
//   return &common.PersonalSkill{
//     Id: skillData.id,
//     Skill_name: skillData.skill_name,
//   }
// }
