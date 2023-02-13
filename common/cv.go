package common

type CVStandardized struct {
	Name string  `json:"name"`
	A1   float64 `json:"a1"`
	A2   float64 `json:"a2"`
	A3   float64 `json:"a3"`
	A4   Gender  `json:"a4"`
	A5   [4]Skill `json:"a5"`
}

type CV struct {
	Name          string      `json:"name"`
	Education     Edu         `json:"education"`
	Language      EnglishCert `json:"language"`
	Exp           Experiment  `json:"exp"`
	Gender        Gender      `json:"gender"`
	PersonalSkill []Skill     `json:"personalskill"`
}

type CVString struct {
	Name          string   `db:"name"           json:"name"`
	Education     string   `db:"education"      json:"education"`
	LanguageCert  string   `db:"language"       json:"language"`
	LanguagePoint float64  `db:"language_point" json:"language_point"`
	Exp           float64  `db:"experience"     json:"exp"`
	Gender        Gender   `db:"gender"         json:"gender"`
	PersonalSkill []string `db:"person_skill"   json:"personalskill"`
}

func (cv *CV) ToCVString() CVString {
	var cvString CVString
	cvString.Name = cv.Name

	switch cv.Education {
	case Bachelor:
		cvString.Education = "Bachelor"
	case Engineer:
		cvString.Education = "Engineer"
	case Master:
		cvString.Education = "Master"
	case PhD:
		cvString.Education = "PhD"
	case Professor:
		cvString.Education = "Professor"
	default:
		cvString.Education = "Bachelor"
	}

	switch cv.Language.Cert {
	case IELTS:
		cvString.LanguageCert = "IELTS"
		cvString.LanguagePoint = cv.Language.Point
	case TOEIC:
		cvString.LanguageCert = "TOEIC"
		cvString.LanguagePoint = cv.Language.Point
	case TOEFL_ITP:
		cvString.LanguageCert = "TOEFL_ITP"
		cvString.LanguagePoint = cv.Language.Point
	case TOEFL_CBT:
		cvString.LanguageCert = "TOEFL_CBT"
		cvString.LanguagePoint = cv.Language.Point
	default:
		cvString.LanguageCert = "TOEFL_IBT"
		cvString.LanguagePoint = cv.Language.Point
	}

	cvString.Exp = float64(cv.Exp)
	cvString.Gender = cv.Gender
	var pSkill []string
	for i := 0; i < len(cv.PersonalSkill); i++ {
		tmpSkill := cv.PersonalSkill[i]
		pSkill = append(pSkill, tmpSkill.SkillString())
	}
	cvString.PersonalSkill = pSkill
	return cvString
}

func (cv *CV) ToCVStandardized() CVStandardized {
	var cvStandardized CVStandardized
	cvStandardized.Name = cv.Name

	switch cv.Education {
	case Bachelor:
		cvStandardized.A1 = 0.1348
	case Engineer:
		cvStandardized.A1 = 0.26967
	case Master:
		cvStandardized.A1 = 0.4045
	case PhD:
		cvStandardized.A1 = 0.5394
	case Professor:
		cvStandardized.A1 = 0.6742
	default:
		cvStandardized.A1 = 0.1348
	}

	cvStandardized.A2 = cv.Language.Standardize()

	cvStandardized.A3 = cv.Exp.Standardize()
	cvStandardized.A4 = cv.Gender
  for _, skill_id := range cv.PersonalSkill {
	  cvStandardized.A5[skill_id] = 1
  }
	return cvStandardized
}

func (cvString *CVString) ToCV() CV {
	var cv CV
	cvString.Name = cv.Name

	switch cvString.Education {
	case "Bachelor":
		cv.Education = Bachelor
	case "Engineer":
		cv.Education = Engineer
	case "Master":
		cv.Education = Master
	case "PhD":
		cv.Education = PhD
	case "Professor":
		cv.Education = Professor
	default:
		cv.Education = Bachelor
	}

	switch cvString.LanguageCert {
	case "IELTS":
		cv.Language.Cert = IELTS
		cv.Language.Point = cvString.LanguagePoint
	case "TOEIC":
		cv.Language.Cert = TOEIC
		cv.Language.Point = cvString.LanguagePoint
	case "TOEFL_ITP":
		cv.Language.Cert = TOEFL_ITP
		cv.Language.Point = cvString.LanguagePoint
	case "TOEFL_CBT":
		cv.Language.Cert = TOEFL_CBT
		cv.Language.Point = cvString.LanguagePoint
	default:
		cv.Language.Cert = TOEFL_IBT
		cv.Language.Point = cvString.LanguagePoint
	}

	cv.Exp = Experiment(cvString.Exp)
	cv.Gender = cvString.Gender
	var pSkill []Skill
	for _, skill := range cvString.PersonalSkill {
		standardSKill := StandardSkillString(skill)
		pSkill = append(pSkill, standardSKill)
	}
	cv.PersonalSkill = pSkill
	return cv
}
