package services

import "github.com/jmoiron/sqlx"

type User struct {
    Id int 
    Email string
    Password string
    Username string 
}

type UserService struct {
    User User
    Db *sqlx.DB
}

func NewUserService(u User, db *sqlx.DB) *UserService {
    return &UserService{
        User: u,
        Db: db,
    }
}

func (us *UserService) CreateUser(u User) error {
    return nil
}
