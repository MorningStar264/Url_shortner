package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MorningStar264/Url_shortner/internal/model"
	"github.com/MorningStar264/Url_shortner/internal/server"
	"github.com/jackc/pgx/v5"
)

type LinkMethods struct {
	server *server.Server
}

func NewLinkMethods(s *server.Server) *LinkMethods {
	return &LinkMethods{server: s}
}

func (m *LinkMethods) CreateLink(l model.Link) {
	conn := m.server.DB.Pool
	_, err := conn.Exec(context.Background(),
		"INSERT INTO links (short_code,long_url,created_by,created_at,expires_at,last_accessed_at) VALUES($1,$2,$3,NOW(),NOW() + INTERVAL '1 day',NOW())", l.ShortCode, l.LongURL, l.CreatedBy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
}

func (m *LinkMethods) GetLink(l model.Link) (model.Link) {
	conn := m.server.DB.Pool
	rows, _ := conn.Query(context.Background(),"SELECT * from links where short_code=$1 limit 1", l.ShortCode)
	link, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[model.Link])
	if err != nil {
		log.Printf("Fetch failed: %v", err)
	}
	fmt.Println(link)
	return link
}

func (m *LinkMethods) DeleteLink(l model.Link) {
	conn := m.server.DB.Pool
	_, err := conn.Exec(context.Background(), "DELETE from links where short_code=$1", l.ShortCode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	fmt.Println("Link deleted: ",l.ShortCode)
}

