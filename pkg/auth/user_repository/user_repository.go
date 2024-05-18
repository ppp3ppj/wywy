package user_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ppp3ppj/wywy/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
    Db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) models.UserRepository {
    return &userRepository{
        Db: db,
    }
}

func (us *userRepository) CreateUser(u models.User) error {
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

func (ur *userRepository) CheckEmail(email string) (models.User, error) {
    query := `SELECT id, email, password, username, first_name, last_name FROM users WHERE email = $1`
    var u models.User
    if err := ur.Db.QueryRow(query, email).Scan(&u.Id, &u.Email, &u.Password, &u.Username, &u.Firstname, &u.Lastname); err != nil {
        return u, fmt.Errorf("email not found")
    }
    return u, nil
}
