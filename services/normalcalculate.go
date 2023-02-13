package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"

	"github.com/rikikudohust/FilterCandidate/common"
)

func NormalCalculate(require common.CV, w common.Weight) []CVStringJson {
	jsonFile, err := os.Open("cvStandards.json")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cvRequire CVJson
	json.Unmarshal(byteValue, &cvRequire)

	requireStandard := require.ToCVStandardized()
	var value []APIValue
	for i, cv := range cvRequire.CVS {
		var W1, W2, W3, W4, W5 float64
		W5 = 0

		W1 = w.W1 * math.Pow((cv.A1-requireStandard.A1), 2)
		W2 = w.W2 * math.Pow((cv.A2-requireStandard.A2), 2)
		W3 = w.W3 * math.Pow((cv.A3-requireStandard.A3), 2)
		if requireStandard.A4 == 2 {
			W4 = 0
		} else {
			W4 = w.W4 * math.Exp(-(math.Pow((float64(cv.A4)-float64(requireStandard.A4)), 2) / 2))
		}
		for i, skill_value := range cv.A5 {
			if skill_value == 1 {
				if requireStandard.A5[i] == 1 {
					W5 += 1
				} else {
					W5 += 0.25
				}
			}
		}

		value = append(value, APIValue{
			Id:    i,
			Score: W1 + W2 + W3 + W4 + W5*w.W5,
		})
	}
	defer jsonFile.Close()
	sort.SliceStable(value, func(i, j int) bool {
		return value[i].Score > value[j].Score
	})

	jsonFile1, err := os.Open("cvStrings.json")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	byteValue1, _ := ioutil.ReadAll(jsonFile1)
	var cvData CVJsonData
  var cvs []CVStringJson
	json.Unmarshal(byteValue1, &cvData)
  for i:= 0; i < len(value) ; i++ {
    var values = cvData.CVS[value[i].Id]
    // var cv CVStringJson
    // cv.Name = values.Name
    // cv.Education = values.Education
    // cv.Exp = values.Exp
    // cv.Language.Cert = values.LanguageCert
    // cv.Language.Point = values.LanguagePoint
    // cv.PersonalSkill = values.PersonalSkill
    // cv.Gender = values.Gender
    // cv.Score = value[i].Score
    values.Score = value[i].Score

    cvs = append(cvs, values)
  }
	defer jsonFile.Close()

	return cvs
}

type CVJson struct {
	CVS []common.CVStandardized `json:"cvs" binding:"required"`
}

type CVJsonData struct {
	CVS []CVStringJson `json:"cvs" binding:"required"`
}

type APIValue struct {
	Id    int
	Score float64
}

type CVStringJson struct {
	Name          string                   `db:"name"           json:"name"`
	Education     string                   `db:"education"      json:"education"`
	Language      common.EnglishCertString `db:"language"       json:"language"`
	Exp           float64                  `db:"experience"     json:"exp"`
	Gender        common.Gender            `db:"gender"         json:"gender"`
	PersonalSkill []string                 `db:"person_skill"   json:"personalskill"`
  Score         float64
}
