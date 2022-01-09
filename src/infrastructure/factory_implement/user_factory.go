package repository_implement

import (
	"amechan/src/domain/entity"
	"amechan/src/infrastructure"
	"errors"
)

type UserFactory struct {
	Database infrastructure.Mysql
}

func NewUserFactory (database infrastructure.Mysql)(userFactory *UserFactory){
	userFactory = new(UserFactory)
	userFactory.Database = database
	return
}

type userMapper struct {
	Id int `db:"id"`
}


// すでにそのLineIdでユーザが登録されてあるかチェックし、登録されていなかったら新しくentity.userを作る
func (userFactory *UserFactory)Create(lineId string, officeCode int, areaCode int)(user *entity.User, err error){
	db, err := userFactory.Database.Connect()
	if err != nil {
		return
	}

	query := `
SELECT id
FROM users
WHERE line_id = ?
`
	var userMapperVar []userMapper
	err = db.Select(&userMapperVar, query, lineId)
	if err != nil {
		return
	}

	if len(userMapperVar) != 0 {
		// ユーザが登録済み
		err = errors.New("this user already exists")
		return
	}

	user = entity.NewUser(lineId, officeCode, areaCode)
	return
}
