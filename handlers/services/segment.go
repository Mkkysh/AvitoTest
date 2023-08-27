package services

import (
	"database/sql"

	"github.com/Mkkysh/AvitoTest/models"
)

type SegmentService struct {
	db *sql.DB
}

func NewSegmentService(db *sql.DB) *SegmentService {
	return &SegmentService{
		db: db,
	}
}

func (s *SegmentService) Add(segment models.Segment) error {

	_, err := s.db.Exec(`INSERT INTO "Segment" (name) VALUES ($1)`, segment.Name)
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
