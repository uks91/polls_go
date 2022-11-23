package poll

// DTO - объект для транспортировки данных
// DAO - объект для сохранения данных в БД

type Option struct {
	Id         string `json:"id,omitempty" db:"options.id,omitempty"`
	Text       string `json:"text,omitempty" db:"options.text,omitempty"`
	QuestionId string `json:"-" db:"options.questionId,omitempty"`
	CreatedAt  string `json:"-" db:"createdAt"`
	UpdatedAt  string `json:"-" db:"updatedAt"`
}
type QuestionDTO struct {
	Text    string      `json:"name,omitempty"`
	Type    interface{} `json:"type,omitempty"`
	Options []string    `json:"options,omitempty"`
}

type Question struct {
	Id        string   `json:"id,int,omitempty" db:"id,omitempty"`
	Text      string   `json:"name,omitempty" db:"text,omitempty"`
	Type      string   `json:"type,omitempty" db:"type"`
	Options   []Option `json:"options,omitempty" db:"options,omitempty"`
	PollId    string   `json:"-" db:"pollId,omitempty"`
	CreatedAt string   `json:"-" db:"createdAt"`
	UpdatedAt string   `json:"-" db:"updatedAt"`
}

type PollDTO struct {
	Id          string        `json:"id,int,omitempty"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Questions   []QuestionDTO `json:"questions,omitempty"`
}

type Poll struct {
	Id          string     `json:"id,omitempty" db:"id"`
	Name        string     `json:"name,omitempty" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	Question    []Question `json:"question,omitempty" db:"-"`
	CreatedAt   string     `json:"-" db:"createdAt"`
	UpdatedAt   string     `json:"-" db:"updatedAt"`
}
