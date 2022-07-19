package entity

type User struct {
	Id        int64
	TeleId    int64
	UserName  string
	FirstName string
	LastName  string
	// TODO: don't forget for isBot checking!
}
