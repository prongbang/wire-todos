package todo

type UseCase interface {
	GetTodoList() []Todo
}

type useCase struct {
	Repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}

func (u *useCase) GetTodoList() []Todo {
	return u.Repo.FindAll()
}
