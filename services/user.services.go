package services

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

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
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
    if err != nil {
        return err
    }
    query := `INSERT INTO users (email, password, username) VALUES ($1, $2, $3) RETURNING "id"`
    if err := us.Db.QueryRowContext(
        ctx,
        query,
        u.Email,
        string(hashPassword),
        u.Username,
    ).Scan(&u.Id); err != nil { 
      switch err.Error() {
            case "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)":
                return fmt.Errorf("username has been used")
            case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
                return fmt.Errorf("email has been used")
            default:
                return fmt.Errorf("insert user failed: %v", err)
        }
    }
    return err 
}
