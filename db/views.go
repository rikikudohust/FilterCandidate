package db

type CVAPIView struct {
	Name          string   `db:"name"                 json:"name"`
	Education     string   `db:"education"            json:"education"`
	LanguageCert  string   `db:"language"             json:"language"`
	LanguagePoint float64  `db:"language_point"       json:"language_point"`
	Exp           float64  `db:"experience"           json:"exp"`
	Gender        int      `db:"gender"               json:"gender"`
	PersonalSkill []string `db:"person_skill_name"    json:"personalskill_name"`
}
