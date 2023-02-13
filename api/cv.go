package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/rikikudohust/FilterCandidate/api/parsers"
	"github.com/rikikudohust/FilterCandidate/common"
	"github.com/rikikudohust/FilterCandidate/services"
)

func (a *API) getACV(c *gin.Context) {
}

func (a *API) postCV(c *gin.Context) {
	type okResponse struct {
		Success string `json:"success"`
	}
	// Return OK
	c.JSON(http.StatusOK, &okResponse{
		Success: "OK",
	})
}

func (a *API) getListCV(c *gin.Context) {
	jsonFile, err := os.Open("cvStandards.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cvRequire CVJson
	json.Unmarshal(byteValue, &cvRequire)
	c.JSON(http.StatusOK, cvRequire)
	defer jsonFile.Close()
}

func (a *API) filterCandidate(c *gin.Context) {

	var require requireCV
	err := c.ShouldBindJSON(&require)
	if err != nil {
		return
	}
  fmt.Println(require)
	commonCVString := requireAPIToCommon(&require)
	data := services.NormalCalculate(commonCVString.ToCV(), require.W)

	c.JSON(http.StatusOK, data)
}

type requireCV struct {
	Education     string        `json:"education"`
	LanguageCert  string        `json:"language"`
	LanguagePoint float64       `json:"language_point"`
	Exp           float64       `json:"exp"`
	Gender        int           `json:"gender"`
	PersonalSkill []string      `json:"personalskill"`
	W             common.Weight `json:"weight"`
}

func requireAPIToCommon(require *requireCV) *common.CVString {
	return &common.CVString{
		Name:          "",
		Education:     require.Education,
		LanguageCert:  require.LanguageCert,
		LanguagePoint: require.LanguagePoint,
		Exp:           require.Exp,
		Gender:        common.Gender(require.Gender),
		PersonalSkill: require.PersonalSkill,
	}
}

// type receiveCV struct {
// 	Name          string         `json:"name"`
// 	Education     string         `json:"education"`
// 	LanguageCert  string         `json:"language"`
// 	LanguagePoint float64        `json:"language_point"`
// 	Exp           float64        `json:"exp"`
// 	Gender        int            `json:"gender"`
// 	PersonalSkill []common.Skill `json:"personalskill"`
// }
//
// func cvCreationAPIToCommon(cvAPI *receiveCV) *common.CVString {
// 	return &common.CVString{
// 		Name:          cvAPI.Name,
// 		Education:     cvAPI.Education,
// 		LanguageCert:  cvAPI.LanguageCert,
// 		LanguagePoint: cvAPI.LanguagePoint,
// 		Exp:           cvAPI.Exp,
// 		Gender:        common.Gender(cvAPI.Gender),
// 		PersonalSkill: cvAPI.PersonalSkill,
// 	}
// }

type CVJson struct {
	CVS []common.CVStandardized `json:"cvs" binding:"required"`
}
