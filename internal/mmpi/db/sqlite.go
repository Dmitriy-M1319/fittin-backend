package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/services"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteTestResultRepository struct {
	db *sql.DB
}

func NewSQLiteTestResultRepository(db *sql.DB) services.TestResultRepository {
	return &sqliteTestResultRepository{db: db}
}

func (r *sqliteTestResultRepository) Init(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS test_results (
		uuid TEXT PRIMARY KEY,
		scales TEXT NOT NULL,
		information TEXT NOT NULL
	);
	`
	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *sqliteTestResultRepository) Create(ctx context.Context, result *models.TestResult) error {
	scalesJSON, err := json.Marshal(result.Scales)
	if err != nil {
		return fmt.Errorf("failed to marshal scales: %w", err)
	}

	query := `INSERT INTO test_results (uuid, scales, information) VALUES (?, ?, ?)`
	_, err = r.db.ExecContext(ctx, query, result.Uuid, scalesJSON, result.Info)
	if err != nil {
		return fmt.Errorf("failed to create test result: %w", err)
	}

	return nil
}

func (r *sqliteTestResultRepository) GetByUUID(ctx context.Context, uuid string) (*models.TestResult, error) {
	query := `SELECT uuid, scales, information FROM test_results WHERE uuid = ?`
	row := r.db.QueryRowContext(ctx, query, uuid)

	var result models.TestResult
	var scalesJSON string

	err := row.Scan(&result.Uuid, &scalesJSON, &result.Info)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("test result not found")
		}
		return nil, fmt.Errorf("failed to get test result: %w", err)
	}

	err = json.Unmarshal([]byte(scalesJSON), &result.Scales)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal scales: %w", err)
	}

	return &result, nil
}

func (r *sqliteTestResultRepository) GetAll(ctx context.Context) ([]*models.TestResult, error) {
	query := `SELECT uuid, scales, information FROM test_results`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get test results: %w", err)
	}
	defer rows.Close()

	var results []*models.TestResult
	for rows.Next() {
		var result models.TestResult
		var scalesJSON string

		err := rows.Scan(&result.Uuid, &scalesJSON, &result.Info)
		if err != nil {
			return nil, fmt.Errorf("failed to scan test result: %w", err)
		}

		err = json.Unmarshal([]byte(scalesJSON), &result.Scales)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal scales: %w", err)
		}

		results = append(results, &result)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return results, nil
}

func (r *sqliteTestResultRepository) Delete(ctx context.Context, uuid string) error {
	query := `DELETE FROM test_results WHERE uuid = ?`
	res, err := r.db.ExecContext(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete test result: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("test result not found")
	}

	return nil
}
