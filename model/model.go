package model

type Find struct {
	Substr string `json:"substr" binding:"required"`
}

type Check struct {
	Email string `json:"email" binding:"required"`
}

type User struct {
	ID        int    `json:"-"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type Hash struct {
	Input string `json:"input" binding:"required"`
}
