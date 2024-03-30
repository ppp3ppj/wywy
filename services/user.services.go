package services

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    Id string 
    Email string
    Password string
    Username string 
    Firstname string
    Lastname string
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
    query := `INSERT INTO users (email, password, username, first_name, last_name) VALUES ($1, $2, $3, $4, $5) RETURNING "id"`
    if err := us.Db.QueryRowContext(
        ctx,
        query,
        u.Email,
        string(hashPassword),
        u.Username,
        u.Firstname,
        u.Lastname,
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

func (us *UserService) CheckEmail(email string) (User, error) {
    query := `SELECT id, email, password, username, first_name, last_name FROM users WHERE email = $1`
    var u User
    if err := us.Db.QueryRow(query, email).Scan(&u.Id, &u.Email, &u.Password, &u.Username, &u.Firstname, &u.Lastname); err != nil {
        return u, fmt.Errorf("email not found")
    }
    return u, nil
}

func (us *UserService) ValidateEmail(email string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    query := `SELECT email FROM users WHERE email = $1`
    var e string
    if err := us.Db.QueryRowContext(ctx, query, email).Scan(&e); err != nil {
        return fmt.Errorf("email found")
    }
    return fmt.Errorf("email not found")
}

func (us *UserService) GetUserProfile(id string) (User, error) {
    query := `SELECT id, email, username, first_name, last_name FROM users WHERE id = $1`
    var u User
    if err := us.Db.QueryRow(query, id).Scan(&u.Id, &u.Email, &u.Username, &u.Firstname, &u.Lastname); err != nil {
        return u, fmt.Errorf("user not found")
    }
    return u, nil
}

func (us *UserService) GetUserById(id string) (User, error) {
    query := `SELECT id, email, username, first_name, last_name FROM users WHERE id = $1`
    var u User
    if err := us.Db.QueryRow(query, id).Scan(&u.Id, &u.Email, &u.Username, &u.Firstname, &u.Lastname); err != nil {
        return u, fmt.Errorf("user not found")
    }
    return u, nil
}
