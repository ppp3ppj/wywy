package models

type User struct {
    Id string
    Email string
    Password string
    Username string
    Firstname string
    Lastname string
}

type UserRepository interface {
    CreateUser(u User) error
    CheckEmail(email string) (User, error)
}
