package poll

// DTO - объект для транспортировки данных
// DAO - объект для сохранения данных в БД

type QuestionDTO struct {
	Text    string
	Type    int
	Options []string
}

type Question struct {
	Text string
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
