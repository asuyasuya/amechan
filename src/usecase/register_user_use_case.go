package usecase

import (
	"amechan/src/domain/factory_interface"
	"amechan/src/domain/repository_interface"
)

type RegisterUserUseCase struct {
	userFactory    factory_interface.UserFactory
	userRepository repository_interface.UserRepository
}

func NewRegisterUserUseCase(repository repository_interface.UserRepository, factory factory_interface.UserFactory) (registerUserUseCase *RegisterUserUseCase) {
	registerUserUseCase = new(RegisterUserUseCase)
	registerUserUseCase.userRepository = repository
	registerUserUseCase.userFactory = factory
	return
}

func (uc *RegisterUserUseCase) Invoke(lineId string, officeCode int, areaCode int) (err error) {
	//すでにそのLINEIDでユーザが登録されてあるかチェックし、オッケーだったら新しくuser (entityのもの)を作る
	user, err := uc.userFactory.Create(lineId, officeCode, areaCode)
	//作ったuserをDBに保存する
	err = uc.userRepository.Insert(user)
	return
}
