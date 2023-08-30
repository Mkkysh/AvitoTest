package services

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Mkkysh/AvitoTest/models"
)

type SegmentService struct {
	db     *sql.DB
	logger *LogService
}

func NewSegmentService(db *sql.DB, logger *LogService) *SegmentService {
	return &SegmentService{
		db:     db,
		logger: logger,
	}
}

func (s *SegmentService) Add(segment models.Segment, partAuto string) error {

	stmt, err := s.db.Prepare(`INSERT INTO "Segment" (name) VALUES ($1) RETURNING id`)
	if err != nil {
		return err
	}

	var id int
	err = stmt.QueryRow(segment.Name).Scan(&id)
	if err != nil {
		return err
	}

	if partAuto == "" {
		return nil
	}

	partAutoInt, err := strconv.ParseFloat(partAuto, 64)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`SELECT *
						FROM "User"
						ORDER BY RANDOM()
						LIMIT (SELECT COUNT(*) * %f FROM "User")`, partAutoInt/100)

	rows, err := s.db.Query(query)
	//_, err = s.db.Exec(query)

	if err != nil {
		return err
	}

	defer rows.Close()

	var result []models.UserSegment
	for rows.Next() {
		var idUser int
		err = rows.Scan(&idUser)
		if err != nil {
			return err
		}
		result = append(result,
			models.UserSegment{
				IdUser:    idUser,
				IdSegment: id,
			},
		)
	}

	values := []string{}
	for _, userSegment := range result {
		values = append(values, fmt.Sprintf("(%d, %d)",
			userSegment.IdUser,
			userSegment.IdSegment))
	}

	valuesStr := strings.Join(values, ", ")

	query = fmt.Sprintf(`INSERT INTO "UserSegment" (id_user, id_segment) VALUES %s`, valuesStr)

	_, err = s.db.Exec(query)

	if err != nil {
		return err
	}

	var logs []models.Log
	for _, userSegment := range result {
		logs = append(logs, models.Log{
			IdUser:    userSegment.IdUser,
			IdSegment: userSegment.IdSegment,
			Operation: "add",
			Timestamp: time.Now(),
		})
	}

	err = s.logger.Add(logs)
	if err != nil {
		return err
	}

	return nil
}

func (s *SegmentService) Delete(name string) error {

	_, err := s.db.Exec(`DELETE FROM "Segment" WHERE name = $1`, name)
	if err != nil {
		return err
	}
	return nil
}
