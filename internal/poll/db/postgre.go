package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal/poll"
	//"github.com/uks91/polls_go/internal/user"
)

var strg poll.Storage = &pollStorage{}

type pollStorage struct {
	db *sqlx.DB
}

type optionRow struct {
	Id   sql.NullString `db:"id,omitempty"`
	Text sql.NullString `db:"text,omitempty"`
}

type questionRow struct {
	Id     string    `db:"id,omitempty"`
	Text   string    `db:"text,omitempty"`
	Type   string    `db:"type,omitempty"`
	PollId string    `db:"pollId,omitempty"`
	Opt    optionRow `db:"options"`
}

func (p *pollStorage) GetQuestions(pollId string) ([]poll.Question, error) {
	queryTemplate := `SELECT "questions"."id", "questions"."text" , "questions"."type", "questions"."pollId" AS "pollId",
	"options"."id" AS "options.id", "options"."text" AS "options.text"
	FROM "questions" LEFT JOIN "options" ON "questions"."id"="options"."questionId" WHERE "pollId"=%s`
	query := fmt.Sprintf(queryTemplate, pollId)

	var questions []questionRow
	err := p.db.Select(&questions, query)
	if err != nil {
		fmt.Printf("select error: %v", err)
		return nil, err
	}

	//fmt.Println(questions)
	var out []poll.Question

	createQuestion := func(row *questionRow) *poll.Question {
		q := poll.Question{
			Id:   row.Id,
			Text: row.Text,
			Type: row.Type,
		}
		out = append(out, q)
		return &q
	}

	for i, row := range questions {
		var q *poll.Question
		if i == 0 {
			q = createQuestion(&row)
		} else {
			q = &out[len(out)-1]
			if q.Id != row.Id {
				q = createQuestion(&row)
			}
		}
		if row.Opt.Id.Valid {
			q.Options = append(q.Options, poll.Option{
				Id:   row.Opt.Id.String,
				Text: row.Opt.Text.String,
			})
		}

	}

	//rows, err := p.db.Queryx(query)
	//if err != nil {
	//	fmt.Printf("Error during Getting: %v", err)
	//} else {
	//	//var optionRow poll.Option
	//	for rows.Next() {
	//		columns, err3 := rows.Columns()
	//		//err3 := rows.StructScan(&optionRow)
	//		if err3 != nil {
	//			fmt.Println("error during scanning row")
	//		} else {
	//			fmt.Println("Columns are: ", columns)
	//		}
	//		var q questionRow
	//		err4 := rows.StructScan(&q)
	//		if err4 != nil {
	//			fmt.Printf("scanning error: %v", err4)
	//		} else {
	//			fmt.Println(q)
	//		}
	//	}
	//}
	//fmt.Printf("Query: %s", query)

	//	queryTemplate2 := `select "questions"."id", "questions"."text", "questions"."type", "questions"."pollId" as "pollId",
	//array_agg("options"."id") as "options.id", array_agg("options"."text") as "options.text"
	//from "questions" left join "options" on "questions"."id"="options"."questionId" where "pollId"=%s group by "questions"."id", "questions"."text", "questions"."type", "pollId"`
	//	query2 := fmt.Sprintf(queryTemplate2, pollId)
	//	var q2 []quest2
	//	err = p.db.Select(&q2, query2)
	//	p.db.QueryRow(query2).Scan(pq.Array())
	//	if err != nil {
	//		fmt.Printf("error in query2: %v", err)
	//	} else {
	//		fmt.Printf("q2 = ", q2)
	//	}
	return out, nil
}

func (p *pollStorage) GetPollsList() ([]poll.Poll, error) {
	var polls []poll.Poll
	err := p.db.Select(&polls, `SELECT * from "polls"`)
	if err != nil {
		return nil, fmt.Errorf("unable to get all polls: %v", err)
	}
	return polls, nil
}

func (p *pollStorage) GetPoll(id string) (poll.Poll, error) {
	var pollObj poll.Poll
	err := p.db.Get(&pollObj, `SELECT * from "polls" WHERE "id" = $1`, id)
	if err != nil {
		return pollObj, fmt.Errorf("unable to get poll with id=%s: %v", id, err)
	}
	questions, err := p.GetQuestions(id)
	if err != nil {
		return pollObj, fmt.Errorf("unable to get questions with pollId=%s: %v", id, err)
	}
	pollObj.Question = questions
	return pollObj, nil
}

func (p *pollStorage) GetQuestion(id string) (poll.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (p *pollStorage) CreatePoll(pollPtr *poll.Poll) (string, error) {
	const pollTemplate = `INSERT INTO "polls" ("name", "description", "createdAt", "updatedAt") 
	VALUES ('%s', '%s', '2022-11-22', '2022-11-22') RETURNING ("id")`

	const questionTemplate = `INSERT INTO "questions" ("text", "type", "createdAt", "updatedAt", "pollId") 
	VALUES ('%s', '%s', '2022-11-22', '2022-11-22', '%s') RETURNING ("id")`

	const optionTemplate = `INSERT INTO "options" ("text", "createdAt", "updatedAt", "questionId") 
	VALUES ('%s', '2022-11-22', '2022-11-22', '%s') RETURNING ("id")`

	//var pollObj poll.Poll
	pollQuery := fmt.Sprintf(pollTemplate, pollPtr.Name, pollPtr.Description)
	rows, err := p.db.Queryx(pollQuery)
	//pollExec, err := p.db.NamedExec(pollTemplate, *pollPtr)
	fmt.Printf(pollQuery)
	if err != nil {
		return "-1", fmt.Errorf("unable to create a new poll: %v", err)
	}

	readId := func(r *sqlx.Rows) (string, error) {
		if r.Next() {
			type idStruct struct {
				Id string `db:"id"`
			}
			var newId idStruct
			e := r.StructScan(&newId)
			if e != nil {
				return "-1", fmt.Errorf("unable to read id: %v", e)
			}
			return newId.Id, nil
		}
		return "-1", fmt.Errorf("unable to read id: rows is empty")
	}

	pollId, err := readId(rows)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, question := range pollPtr.Question {
		//questExec, err := p.db.NamedExec(questionTemplate, question)
		questQuery := fmt.Sprintf(questionTemplate, question.Text, question.Type, pollId)
		fmt.Println(questQuery)
		rows, err = p.db.Queryx(questQuery)
		if err != nil {
			return "-1", fmt.Errorf("unable to create a new question: %v", err)
		}
		questId, err := readId(rows)
		if err != nil {
			fmt.Printf("%v", err)
		}
		for _, option := range question.Options {
			optionQuery := fmt.Sprintf(optionTemplate, option.Text, questId)
			fmt.Printf(optionQuery)
			_, err := p.db.Queryx(optionQuery)
			//_, err := p.db.NamedExec(optionTemplate, option)
			if err != nil {
				return "-1", fmt.Errorf("unable to create a new option: %v", err)
			}
		}
	}
	return pollId, nil
}

func NewStorage(db *sqlx.DB) poll.Storage {
	return &pollStorage{
		db: db,
	}
}
