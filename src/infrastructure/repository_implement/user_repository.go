package repository_implement

import (
	"amechan/src/domain/entity"
	"amechan/src/infrastructure"
)

type UserRepository struct {
	Database infrastructure.Mysql
}

func NewUserRepository (database infrastructure.Mysql)(userRepository *UserRepository){
	userRepository = new(UserRepository)
	userRepository.Database = database
	return
}

func (userRepo *UserRepository)Create(user *entity.User)(err error){
	db, err := userRepo.Database.Connect()
	if err != nil {
		return
	}

	query := `
INSERT INTO users(line_id, office_code, area_code)
VALUE (?, ?, ?);
`
	result := db.MustExec(query, user.LineId, user.OfficeCode, user.AreaCode)
	_, err = result.RowsAffected()
	return
}

