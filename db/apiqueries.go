package db

import (
	"github.com/rikikudohust/FilterCandidate/common"
)

func(itemDb *ItemDB) AddListSkillAPI(skills []common.PersonalSkill) error {
  return itemDb.AddListSkill(skills)
}

func(itemDb *ItemDB) AddACVAPI(cv common.CVString) error {
  return itemDb.AddACV(cv)
}

func (itemDb *ItemDB) GetACVAPI(id int) (*CVAPIView, error) {
  return nil, nil

}

func (itemDb *ItemDB) GetListSkillAPI() ([]common.PersonalSkill, error) {
  return nil, nil
}

