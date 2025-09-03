package models

type Bucket struct {
	UserId uint
	Pets   []Pet
}

type Favorite struct {
	UserId uint
	Pets   []Pet
}

type Buyer struct {
	User
}
