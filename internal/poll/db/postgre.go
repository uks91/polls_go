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

type opt struct {
	Id2   sql.NullString `db:"id,omitempty"`
	Text2 sql.NullString `db:"text,omitempty"`
}

type quest struct {
	Id   string `db:"id,omitempty"`
	Text string `db:"text,omitempty"`
	Type string `db:"type,omitempty"`
	//CreatedAt string `db:"createdAt,omitempty" json:"-"`
	//UpdatedAt string `db:"updatedAt,omitempty" json:"-"`
	PollId string `db:"pollId,omitempty"`
	Opt    opt    `db:"options"`
	//QuestionId string `db:"options.questionId,omitempty"`
}

//type opt2 struct {
//	Id2   sql.NullString `db:"id,omitempty"`
//	Text2 sql.NullString `db:"text,omitempty"`
//}
//
//type quest2 struct {
//	Id     string           `db:"id,omitempty"`
//	Text   string           `db:"text,omitempty"`
//	Type   string           `db:"type,omitempty"`
//	PollId string           `db:"pollId,omitempty"`
//	Ids    []sql.NullString `db:"options.id"`
//	Texts  []sql.NullString `db:"options.text"`
//}

func (p *pollStorage) GetQuestions(pollId string) ([]poll.Question, error) {
	queryTemplate := `SELECT "questions"."id", "questions"."text" , "questions"."type", "questions"."pollId" AS "pollId",
	"options"."id" AS "options.id", "options"."text" AS "options.text"
	FROM "questions" LEFT JOIN "options" ON "questions"."id"="options"."questionId" WHERE "pollId"=%s`
	query := fmt.Sprintf(queryTemplate, pollId)

	var questions []quest
	err := p.db.Select(&questions, query)
	if err != nil {
		fmt.Printf("select error: %v", err)
		return nil, err
	}

	fmt.Println(questions)
	//out := make([]poll.Question, 1)
	var out []poll.Question

	type createQuestion func(s *[]poll.Question, r quest) *poll.Question {

	}

	for i, row := range questions {
		var q poll.Question
		if i == 0 {
			q = poll.Question{
				Id:   row.Id,
				Text: row.Text,
				Type: row.Type,
			}
			out = append(out, q)
		} else {
			q = out[len(out)-1]
			if q.Id != row.Id {
				q = poll.Question{
					Id:   row.Id,
					Text: row.Text,
					Type: row.Type,
				}
				out = append(out, q)
			}
		}
		if row.Opt.Id2.Valid {
			q.Options = append(q.Options, poll.Option{
				Id:   row.Opt.Id2.String,
				Text: row.Opt.Text2.String,
			})
		}

	}

	//rows, err := p.db.Queryx(query)
	//if err != nil {
	//	fmt.Printf("Error during Getting: %v", err)
	//} else {
	//	//var opt poll.Option
	//	for rows.Next() {
	//		columns, err3 := rows.Columns()
	//		//err3 := rows.StructScan(&opt)
	//		if err3 != nil {
	//			fmt.Println("error during scanning row")
	//		} else {
	//			fmt.Println("Columns are: ", columns)
	//		}
	//		var q quest
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
	return poll.Poll{}, nil
}

func (p *pollStorage) GetQuestion(id string) (poll.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (p *pollStorage) CreatePoll(poll *poll.Poll) (int64, error) {
	//TODO implement me
	panic("implement me")
}

/*
func (u *pollStorage) GetUserByLogin(login string) (user.User, error) {
	var usr user.User
	query := fmt.Sprintf(`SELECT "id", "login", "password", "role" FROM users WHERE "login" = '%s'`, login)
	err := u.db.Get(&usr, query)
	if err != nil {
		return usr, fmt.Errorf("user with login=%s not found: %v", login, err)
	}
	return usr, nil
}

func (u *pollStorage) GetOne(uuid string) (user.User, error) {
	var usr user.User
	query := fmt.Sprintf(`SELECT "id", "login", "role", "password" from "users" WHERE "id"=%s`, uuid)
	err := u.db.Get(&usr, query)
	if err != nil {
		fmt.Printf("failed query id: %s", query)
		return usr, fmt.Errorf("user with id=%s not found: %v", uuid, err)
	}
	fmt.Println("usr == ", usr)
	return usr, nil
}

func (u *pollStorage) GetAll() ([]user.User, error) {
	var users []user.User
	err := u.db.Select(&users, `SELECT "id", "login", "role", "password" from "users"`)
	if err != nil {
		return nil, fmt.Errorf("unable to get all users: %v", err)
	}

	return users, nil
}

func (u *pollStorage) Create(usr user.User) (user.User, error) {
	out := usr
	query := fmt.Sprintf(`INSERT INTO "users" ("login", "password", "role", "createdAt", "updatedAt")  VALUES ('%s', '%s', '%s', '2022-11-22', '2022-11-22') RETURNING ("id", "login", "password", "role")`, usr.Username, usr.PasswordHash, usr.Role)
	fmt.Println(query)
	//row := u.db.QueryRowx(query)
	exec, err := u.db.NamedExec(`INSERT INTO "users" ("login", "password", "role", "createdAt", "updatedAt")  VALUES (:login, :password, :role, '2022-11-22', '2022-11-22') RETURNING ("id", "login", "password", "role")`, usr)
	if err != nil {
		return user.User{}, fmt.Errorf("unable to create new user: %v", err)
	}
	id, _ := exec.LastInsertId()
	fmt.Printf("inserted id is %s", string(id))
	out.ID = string(id)
	return out, nil
}

func (u *pollStorage) Delete(book *user.User) error {
	//TODO implement me
	panic("implement me")
}
*/

func NewStorage(db *sqlx.DB) poll.Storage {
	return &pollStorage{
		db: db,
	}
}
