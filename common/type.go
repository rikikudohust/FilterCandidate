package common

type Edu int
type English int
type Experiment float64
type Gender int
type Skill int

const (
	Bachelor Edu = iota
	Engineer
	Master
	PhD
	Professor
)

const (
	IELTS English = iota
	TOEIC
	TOEFL_ITP
	TOEFL_CBT
	TOEFL_IBT
)

// const (
// 	Intern Experiment = iota
// 	Fresher
// 	Junior
// 	Senior
// )

const (
	Male   Gender = 0
	Female Gender = 1
)

const (
	TeamWork Skill = iota
	OfficeInformation
	Analysis
	InformationProcessing
)

type EnglishCert struct {
	Cert  English
	Point float64
}

type EnglishCertString struct {
	Cert  string  `json:"Cert"`
	Point float64 `json:"Point"`
}

func (skill *Skill) SkillString() string {
	switch *skill {

	case TeamWork:
		return "TeamWork"
	case OfficeInformation:
		return "Office Information"
	case Analysis:
		return "Analysis"
	case InformationProcessing:
		return "InformationProcessing"
	default:
		return "Nothing"
	}
}

func (elc *EnglishCert) Standardize() float64 {
	point := elc.Point
	var standardPoint float64
	switch elc.Cert {
	case IELTS:
		if point < 2 {
			standardPoint = 0
		} else if point >= 2 && point < 3.5 {
			standardPoint = 0.1048
		} else if point >= 3.5 && point < 4.5 {
			standardPoint = 0.2097
		} else if point >= 4.5 && point < 5.5 {
			standardPoint = 0.3145
		} else if point >= 5.5 && point < 6.5 {
			standardPoint = 0.4193
		} else if point >= 6.5 && point < 7.5 {
			standardPoint = 0.5241
		} else {
			standardPoint = 0.6290
		}
	case TOEIC:
		if point < 255 {
			standardPoint = 0
		} else if point >= 255 && point < 400 {
			standardPoint = 0.1048
		} else if point >= 400 && point < 450 {
			standardPoint = 0.2097
		} else if point >= 450 && point < 600 {
			standardPoint = 0.3145
		} else if point >= 600 && point < 850 {
			standardPoint = 0.4193
		} else if point >= 850 && point < 910 {
			standardPoint = 0.5241
		} else {
			standardPoint = 0.6290
		}
	case TOEFL_ITP:
		if point < 347 {
			standardPoint = 0
		} else if point >= 347 && point < 400 {
			standardPoint = 0.1048
		} else if point >= 400 && point < 450 {
			standardPoint = 0.2097
		} else if point >= 450 && point < 500 {
			standardPoint = 0.3145
		} else if point >= 500 && point < 550 {
			standardPoint = 0.4193
		} else if point >= 550 && point < 600 {
			standardPoint = 0.5241
		} else {
			standardPoint = 0.6290
		}
	case TOEFL_CBT:
		if point < 60 {
			standardPoint = 0
		} else if point >= 60 && point < 96 {
			standardPoint = 0.1048
		} else if point >= 96 && point < 133 {
			standardPoint = 0.2097
		} else if point >= 133 && point < 173 {
			standardPoint = 0.3145
		} else if point >= 173 && point < 213 {
			standardPoint = 0.4193
		} else if point >= 213 && point < 250 {
			standardPoint = 0.5241
		} else {
			standardPoint = 0.6290
		}
	default:
		if point < 19 {
			standardPoint = 0
		} else if point >= 2 && point < 40 {
			standardPoint = 0.1048
		} else if point >= 3.5 && point < 45 {
			standardPoint = 0.2097
		} else if point >= 4.5 && point < 61 {
			standardPoint = 0.3145
		} else if point >= 5.5 && point < 80 {
			standardPoint = 0.4193
		} else if point >= 6.5 && point < 100 {
			standardPoint = 0.5241
		} else {
			standardPoint = 0.6290
		}
	}

	return standardPoint
}

func (exp Experiment) Standardize() float64 {
	if exp >= 0 && exp < 1 {
		return 0.0809
	} else if exp >= 1 && exp < 2 {
		return 0.3234
	} else if exp >= 2 && exp < 4 {
		return 0.4851
	} else {
		return 0.8085
	}
}

func StandardSkillString(name string) Skill {
	switch name {

	case "TeamWork":
		return TeamWork
	case "Office Information":
		return OfficeInformation
	case "Analysis":
		return Analysis
	case "Information Processing":
		return InformationProcessing
	default:
		return 0
	}
}

type Weight struct {
	W1 float64 `json:"w1"`
	W2 float64 `json:"w2"`
	W3 float64 `json:"w3"`
	W4 float64 `json:"w4"`
	W5 float64 `json:"w5"`
}
