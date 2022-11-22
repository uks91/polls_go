package poll

// DTO - объект для транспортировки данных
// DAO - объект для сохранения данных в БД

type QuestionDTO struct {
	Text    string
	Type    int
	Options []string
}

type Option struct {
	Id         string `db:"options.id,omitempty"`
	Text       string `db:"options.text,omitempty"`
	QuestionId string `db:"options.questionId,omitempty"`
	//CreatedAt  string `db:"createdAt,omitempty" json:"-"`
	//UpdatedAt  string `db:"updatedAt,omitempty" json:"-"`
}

type Question struct {
	Id      string   `db:"id,omitempty"`
	Text    string   `db:"text,omitempty"`
	Type    string   `db:"type"`
	Options []Option `db:"options,omitempty"`
	//CreatedAt string `db:"createdAt,omitempty" json:"-"`
	//UpdatedAt string `db:"updatedAt,omitempty" json:"-"`
}

type PollDTO struct {
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions,omitempty"`
}

type Poll struct {
	Id          string     `json:"id,omitempty" db:"id"`
	Name        string     `json:"name,omitempty" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	Question    []Question `json:"question,omitempty" db:"-"`
	CreatedAt   string     `json:"createdAt,omitempty" db:"createdAt"`
	UpdatedAt   string     `json:"updatedAt,omitempty" db:"updatedAt"`
}
