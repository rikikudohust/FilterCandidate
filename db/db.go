package db

import (
	// "postgres/common"
	"context"
	"fmt"

	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rikikudohust/FilterCandidate/common"
)

type ItemDB struct {
	dbWrite *sqlx.DB
}

func NewItemDB(
	dbWrite *sqlx.DB,
) *ItemDB {
	return &ItemDB{
		dbWrite: dbWrite,
	}
}

func (itemDB *ItemDB) DB() *sqlx.DB {
	return itemDB.dbWrite
}

func (itemDb *ItemDB) AddACV(cv common.CVString) error {
	tx, err := itemDb.dbWrite.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	queryInsertCV := `INSERT INTO cv (name, education, language, language_point, experience, gender) VALUES ($1, $2, $3, $4, $5, $6)`
	queryInsertSkill := `INSERT INTO skill(cv_id, skill_id)`

	result, err := itemDb.dbWrite.Exec(queryInsertCV, cv.Name, cv.Education, cv.LanguageCert, cv.LanguagePoint, cv.Exp, cv.Gender)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)
	for _, skill := range cv.PersonalSkill {
		_, err = itemDb.dbWrite.Exec(queryInsertSkill, 2, skill)
		if err != nil {
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (itemDb *ItemDB) AddListSkill(skills []common.PersonalSkill) error {
	tx, err := itemDb.dbWrite.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	query := `INSERT INTO personal_skill (id, skill_name)
  VALUES(:id, :skill_name) ON CONFLICT (id) DO NOTHING;`

	_, err = sqlx.NamedExec(itemDb.dbWrite, query, skills)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


// func (itemDB *ItemDB) AddNewQuestion(question common.Question) error {
//   _, err := itemDB.dbWrite.Exec(
//     `INSERT INTO "Question" (id, content, amount_answer, url ,vote)
//     VALUES ($1, $2, $3, $4, $5);`,
//     question.ID, question.Content, question.AmountAnswer, question.URL, question.Vote,
//     )
//
//   return err
// }
//
// func (itemDB *ItemDB) AddNewAnswer(answer common.Answer) error {
//   _, err := itemDB.dbWrite.Exec(
//     `INSERT INTO "Answer" (id, content, vote, question_id)
//     VALUES ($1, $2, $3, $4)
//     ON CONFLICT (id) DO NOTHING;`,
//     answer.ID, answer.Content, answer.Vote, answer.Question_ID,
//     )
//
//   return err
// }
//
// func (itemDB *ItemDB) AddManyQuestion(questions []common.Question) error {
//   _, err := sqlx.NamedExec(itemDB.dbWrite,
//     `INSERT INTO "Question" (id, content, amount_answer, url ,vote)
//     VALUES (:id, :content, :amount_answer, :url, :vote)
//     ON CONFLICT (id) DO NOTHING`,
//     questions,
//     )
//   return err
// }
//
// func (itemDB *ItemDB) AddManyAnswer(answers []common.Answer) error {
//   _, err := sqlx.NamedExec(itemDB.dbWrite,
//     `INSERT INTO "Answer" (id, content, vote, question_id)
//     VALUES (:id, :content, :vote, :question_id)
//     ON CONFLICT (id) DO NOTHING;`,
//     answers,
//     )
//   return err
// }
//
// func (itemDB *ItemDB) getAQuestion(id int) (*common.Question, error) {
//   var question common.Question
//   row := itemDB.dbWrite.QueryRowx(`SELECT * FROM "Question" WHERE id = $1 LIMIT 1`, id)
//   err := row.StructScan(&question)
//   return &question, err
// }
//
// func (itemDB *ItemDB) GetListQuestion(limit int, page int) ([]common.Question, error) {
//   offset := limit * (page - 1)
//   SQL := `SELECT * FROM "Question" ORDER BY "id" ASC LIMIT $2 OFFSET $1`
//   rows, err := itemDB.dbWrite.Queryx(SQL, offset, limit)
//   if err != nil {
//     log.Println(err)
//     return nil, err
//   }
//
//   var questions []common.Question
//
//   for rows.Next() {
//     var question common.Question
//     err = rows.StructScan(&question)
//     if err != nil {
//       log.Println(err)
//       return nil, err
//     }
//     questions = append(questions, question)
//   }
//   return questions, nil
// }
//
// func (itemDB *ItemDB) GetListAnswer(limit int, page int) ([]common.Answer, error) {
//   offset := limit * (page - 1)
//   SQL := `SELECT * FROM "Answer" ORDER BY "id" ASC LIMIT $2 OFFSET $1`
//   rows, err := itemDB.dbWrite.Queryx(SQL, offset, limit)
//   if err != nil {
//     log.Println(err)
//     return nil, err
//   }
//
//   var answers []common.Answer
//
//   for rows.Next() {
//     var answer common.Answer
//     err = rows.StructScan(&answer)
//     if err != nil {
//       log.Println(err)
//       return nil, err
//     }
//     answers = append(answers, answer)
//   }
//   return answers, nil
// }
//
// func (itemDB *ItemDB) getAAnswer(id int) (*common.Answer, error) {
//   var answer common.Answer
//   row := itemDB.dbWrite.QueryRowx(`SELECT * FROM "Answer" WHERE id = $1 LIMIT 1`, id)
//   err := row.StructScan(&answer)
//   return &answer, err
// }
//
// func (itemDB *ItemDB) AddManyDocument(documents []common.DocumentV2) error {
//   _, err := sqlx.NamedExec(itemDB.dbWrite,
//     `INSERT INTO "Document" (id, content, type ,vote, question_id)
//     VALUES (:id, :content, :type, :vote, :question_id)
//     ON CONFLICT (id) DO NOTHING;`,
//     documents,
//     )
//   return err
// }
//
