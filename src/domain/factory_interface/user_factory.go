package factory_interface

import "amechan/src/domain/entity"

type UserFactory interface {
	Create(string, int, int)(*entity.User, error)
}

