package queue

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type queueMysql struct {
	db *sql.DB
}

func NewQueue(db *sql.DB) Queue {
	return &queueMysql{db: db}
}

func (m *queueMysql) Count(ctx context.Context) int64 {
	rows, err := m.db.Query("SELECT COUNT(*) FROM job")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	return int64(count)
}

func (m *queueMysql) Push(ctx context.Context, job *Job) error {
	sqlStatement := fmt.Sprintf(`INSERT INTO job (name) VALUES ('%v')`, job.Message)
	_, err := m.db.Exec(sqlStatement)

	return err
}

func (m *queueMysql) Pop(ctx context.Context) (*Job, error) {
	var job Job
	err := m.db.QueryRowContext(ctx, `SELECT message FROM job LIMIT 1`).Scan(&job.Message)

	return &job, err
}
