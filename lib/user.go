package go_playcraft_server

type User struct {
	userId string
	username string
	email string
	firstname string
	middlename string
}

//Basic CRUD functions
func (u *User) CreateUser (){}
func (u *User) DeleteUser (){}
func (u *User) EditUser (){}
func (u *User) UpdateUser (){}

//Utility functions
func (u *User) ChangeUserPassword (){}
func (u *User) ResetPassword (){}
