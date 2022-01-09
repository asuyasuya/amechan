package entity

type User struct {
	Id         int
	LineId     string
	OfficeCode int
	AreaCode   int
}

func NewUser(lineId string, officeCode int, areaCode int)(user *User){
	user = new(User)
	user.LineId = lineId
	user.OfficeCode = officeCode
	user.AreaCode = areaCode
	return
}
