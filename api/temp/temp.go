package temp

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getConnectionString() string {

	port := 5433
	username := "postgres"
	password := "postgresql5432"
	host_name := "localhost"
	database_name := "postgres"

	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", username, password, host_name, port, database_name)
	return connectionString
}

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Link struct {
	ID             int
	ShortCode      string
	LongURL        string
	ClicksCount    int
	LastAccessedAt time.Time
	CreatedBy      int
	CreatedAt      time.Time
	ExpiresAt      time.Time
	IsActive       bool
}

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connString := getConnectionString()
	ctx:=context.Background()
	Pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer Pool.Close()
	
}

// exec for updating ,inserting, deleting
func createUser(conn *pgxpool.Pool, u User) {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO users (username,email,password_hash) VALUES($1,$2,$3)", u.Username, u.Email, u.PasswordHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
}

func getUser(conn *pgxpool.Pool) {
	rows, _ := conn.Query(context.Background(), "select * from users;")
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])
	//RowTo[] single column
	//RowToStructByPos
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func updateUser(conn *pgxpool.Pool, u User) {
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
	query += fmt.Sprintf(" WHERE id = $%d", argID)
	args = append(args, u.ID)

	_, err := conn.Exec(context.Background(), query, args...)
	if err != nil {
		log.Printf("Update failed: %v", err)
	}
}

func deleteUser(conn *pgxpool.Pool, u User) {
	_, err := conn.Exec(context.Background(),
		"DELETE from users where id=$1", u.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
}

func createLink(conn *pgxpool.Pool, l Link) {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO links (short_code, long_url, created_by, expires_at) 
		VALUES ($1, $2, $3, NOW() + INTERVAL '30 days')`, l.ShortCode, l.LongURL, l.CreatedBy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
}

func getLink(conn *pgxpool.Pool, l Link) Link {
	rows, _ := conn.Query(context.Background(), "select * from users where short_code=$1;", l.ShortCode)
	links, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Link])
	if err != nil {
		log.Fatal(err)
	}
	var link Link
	for _, v := range links {
		link = v
	}
	return link
}
