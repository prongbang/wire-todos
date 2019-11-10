//+build wireinject

package todo

import "github.com/google/wire"

func Initial() (UseCase, error) {
	wire.Build(NewUseCase, NewRepository)
	return nil, nil
}
