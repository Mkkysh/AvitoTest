package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Mkkysh/AvitoTest/dto"
	"github.com/Mkkysh/AvitoTest/models"
)

type LogService struct {
	db *sql.DB
}

func NewLogService(db *sql.DB) *LogService {
	return &LogService{
		db: db,
	}
}

func (s *LogService) Add(logs []models.Log) error {

	values := []string{}
	for _, log := range logs {
		values = append(values, fmt.Sprintf(`(%d, %d, '%s', '%s')`,
			log.IdUser,
			log.IdSegment,
			log.Operation,
			log.Timestamp.Format("2006-01-02 15:04:05")))
	}

	valuesStr := strings.Join(values, ", ")

	query := fmt.Sprintf(`INSERT INTO "Log" 
						(id_user, id_segment, operation, timestamp) 
						VALUES %s`, valuesStr)

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *LogService) Get(date string) ([]dto.LogResponse, error) {

	year, month := strings.Split(date, "-")[0], strings.Split(date, "-")[1]

	query := `SELECT id_user, name, operation, timestamp FROM "Log" l 
	LEFT JOIN "Segment" s on l.id_segment = s.id
	WHERE EXTRACT(YEAR FROM timestamp) = $1 
	AND EXTRACT(MONTH FROM timestamp) = $2`

	rows, err := s.db.Query(query, year, month)

	if err != nil {
		return nil, err
	}

	var logs []dto.LogResponse
	for rows.Next() {
		var log dto.LogResponse
		err = rows.Scan(&log.IdUser, &log.NameSegment,
			&log.Operation, &log.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}
