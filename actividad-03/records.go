package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Record struct {
	ID    int64  `json:"id"`
	Value string `json:"record_value"`
	Size  int    `json:"size"`
}

type recordStore struct {
	db *sql.DB
}

func newRecordStore(db *sql.DB) *recordStore {
	return &recordStore{db: db}
}

var (
	errNotFound   = errors.New("record not found")
	errBadRequest = errors.New("bad request")
)

func (s *recordStore) list() ([]Record, error) {
	rows, err := s.db.Query(`SELECT id, value, size FROM records ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Value, &r.Size); err != nil {
			return nil, err
		}
		records = append(records, r)
	}

	return records, rows.Err()
}

func (s *recordStore) get(id int64) (Record, error) {
	var r Record
	err := s.db.QueryRow(`SELECT id, value, size FROM records WHERE id = ?`, id).
		Scan(&r.ID, &r.Value, &r.Size)
	if errors.Is(err, sql.ErrNoRows) {
		return Record{}, errNotFound
	}
	return r, err
}

func (s *recordStore) create(value string) (Record, error) {
	size := len([]rune(value))
	if size == 0 || size > 100 {
		return Record{}, fmt.Errorf("%w: invalid input size must be >0 and <=100", errBadRequest)
	}

	result, err := s.db.Exec(`INSERT INTO records (value, size) VALUES (?, ?)`, value, size)
	if err != nil {
		return Record{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Record{}, err
	}

	return Record{ID: id, Value: value, Size: size}, nil
}

func (s *recordStore) delete(id int64) error {
	result, err := s.db.Exec(`DELETE FROM records WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errNotFound
	}

	return nil
}

func registerRecordRoutes(r *gin.Engine, store *recordStore) {
	r.GET("/records", func(c *gin.Context) {
		records, err := store.list()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if records == nil {
			records = []Record{}
		}
		c.JSON(http.StatusOK, records)
	})

	r.GET("/records/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		record, err := store.get(id)
		if errors.Is(err, errNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, record)
	})

	r.POST("/records", func(c *gin.Context) {
		var value string
		if err := c.ShouldBindJSON(&value); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "body must be a JSON string"})
			return
		}

		record, err := store.create(value)
		if err == nil {
			c.JSON(http.StatusCreated, record)
			return
		}

		if errors.Is(err, errBadRequest) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	})

	r.DELETE("/records/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		err = store.delete(id)
		if err == nil {
			c.Status(http.StatusNoContent)
			return
		}

		if errors.Is(err, errNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	})
}
