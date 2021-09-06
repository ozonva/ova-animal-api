package animal

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"ova-animal-api/internal/config"
	"ova-animal-api/internal/domain"
	"time"
)

//go:generate mockgen -source=animal_repo.go -destination=animal_repo_mock.go -package=animal
type Repo interface {
	AddEntities(entities []domain.Animal) error
	ListEntities(limit, offset uint64) ([]domain.Animal, error)
	DescribeEntity(entityId uint64) (*domain.Animal, error)
	Delete(id uint64) error
}

func New(settings config.Db) (Repo, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		settings.Login,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Name,
	)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	db.SetConnMaxIdleTime(3 * time.Second)

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	r := repo{
		db: db,
	}
	return &r, nil
}

type repo struct {
	db *sql.DB
}

func (this *repo) AddEntities(entities []domain.Animal) error {
	ctx := context.Background()
	conn, err := this.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	commited := false
	tx, err := conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if !commited {
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO animal(user_id, name, type) VALUES ($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, e := range entities {
		_, err := stmt.Exec(e.UserId, e.Name, e.Type.String())
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	commited = true
	return err
}

func (this *repo) ListEntities(limit, offset uint64) ([]domain.Animal, error) {
	rows, err := this.db.Query("SELECT id, user_id, name, type FROM animal")
	if err != nil {
		return nil, err
	}

	result := make([]domain.Animal, 0)
	for rows.Next() {
		animal := domain.Animal{}
		if err = this.scanEntitiy(rows, &animal); err != nil {
			return nil, err
		}
		result = append(result, animal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, err
}

func (this *repo) DescribeEntity(entityId uint64) (*domain.Animal, error) {
	rows, err := this.db.Query("SELECT id, user_id, name, type FROM animal WHERE id=$1", entityId)
	if err != nil {
		return nil, err
	}

	animal := domain.Animal{}
	cnt := 0

	for rows.Next() {
		if cnt > 0 {
			return nil, fmt.Errorf("Selected more then 1 rows by id=%d", entityId)
		}
		cnt++

		if err = this.scanEntitiy(rows, &animal); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if cnt == 0 {
		return nil, fmt.Errorf("animal not found by id %d", entityId)
	}

	return &animal, nil
}

func (this *repo) Delete(id uint64) error {
	_, err := this.db.Exec("DELETE FROM animal WHERE id=$1", id)
	return err
}

func (this *repo) scanEntitiy(rows *sql.Rows, animal *domain.Animal) error {
	var (
		err     error
		typeStr string
	)
	if err := rows.Scan(&animal.Id, &animal.UserId, &animal.Name, &typeStr); err != nil {
		return err
	}
	animal.Type, err = domain.AnimalTypeString(typeStr)
	return err
}
