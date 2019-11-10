package todo

type Repository interface {
	FindAll() []Todo
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() []Todo {
	return []Todo{
		Todo{
			ID:        1,
			UserID:    1,
			Title:     "Todo 1",
			Completed: true,
		},
		Todo{
			ID:        2,
			UserID:    1,
			Title:     "Todo 2",
			Completed: false,
		},
	}
}
