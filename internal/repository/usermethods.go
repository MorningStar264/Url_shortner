package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/MorningStar264/Url_shortner/internal/model"
	"github.com/MorningStar264/Url_shortner/internal/server"
	"github.com/jackc/pgx/v5"
)

type UserMethods struct {
	server *server.Server
}

func NewUserMethods(s *server.Server) *UserMethods {
	return &UserMethods{server: s}
}

func (m *UserMethods) CreateUser(u model.User) {
	conn := m.server.DB.Pool
	_, err := conn.Exec(context.Background(),
		"INSERT INTO users (username,email,password_hash) VALUES($1,$2,$3)", u.Username, u.Email, u.PasswordHash)
	if err != nil {
		log.Printf("Creation failed: %v", err)
	}
}


func (m *UserMethods) GetUser(u model.User) model.User {
	conn := m.server.DB.Pool
	rows, _ := conn.Query(context.Background(), "select * from users where username=$1 limit 1;", u.Username)
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[model.User])

	if err != nil {
		log.Printf("Fetch failed: %v", err)
	}
	return user

}

func (m *UserMethods) UpdateUser(u model.User) {
	conn := m.server.DB.Pool

	query := "UPDATE users SET "
	updates := []string{}
	args := []any{}
	argID := 1
	if u.Username != "" {
		updates = append(updates, fmt.Sprintf("username = $%d", argID))
		args = append(args, u.Username)
		argID++
	}

	if u.Email != "" {
		updates = append(updates, fmt.Sprintf("email = $%d", argID))
		args = append(args, u.Email)
		argID++
	}

	if u.PasswordHash != "" {
		updates = append(updates, fmt.Sprintf("password_hash = $%d", argID))
		args = append(args, u.PasswordHash)
		argID++
	}

	updates = append(updates, "updated_at = NOW()")
	query += strings.Join(updates, ", ")
	query += fmt.Sprintf(" WHERE username = $%d", argID)
	args = append(args, u.Username)
	_, err := conn.Exec(context.Background(), query, args...)

	if err != nil {
		log.Printf("Update failed: %v", err)
	}
}

func (m *UserMethods) DeleteUser(u model.User) {
	conn := m.server.DB.Pool
	_, err := conn.Exec(context.Background(), "DELETE from users where username=$1", u.Username)
	if err != nil {
		log.Printf("Delete failed: %v", err)
	}
}
