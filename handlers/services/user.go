package services

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Mkkysh/AvitoTest/models"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) UpdateSegment(id int, AddSegments []interface{}, RemoveSegments []interface{}) error {

	if len(AddSegments) != 0 {

		placeholders := make([]string, len(AddSegments))

		for i := range AddSegments {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}

		query := fmt.Sprintf(`SELECT id FROM "Segment" WHERE name IN (%s)`,
			strings.Join(placeholders, ", "))

		rows, err := u.db.Query(query, AddSegments...)

		if err != nil {
			log.Println(err)
			return err
		}

		defer rows.Close()

		var result []models.UserSegment

		for rows.Next() {
			var idSeg int
			err = rows.Scan(&idSeg)
			if err != nil {
				return err
			}
			result = append(result,
				models.UserSegment{
					IdUser:    id,
					IdSegment: idSeg,
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
		log.Println(query)
		_, err = u.db.Exec(query)

		if err != nil {
			return err
		}

	}

	if len(RemoveSegments) != 0 {
		placeholders := make([]string, len(RemoveSegments))

		for i := range RemoveSegments {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}

		query := fmt.Sprintf(`SELECT id FROM "Segment" WHERE name IN (%s)`,
			strings.Join(placeholders, ", "))

		rows, err := u.db.Query(query, RemoveSegments...)

		if err != nil {
			log.Println(err)
			return err
		}

		defer rows.Close()

		var result []interface{}
		result = append(result, id)

		for rows.Next() {
			var idSeg int
			err = rows.Scan(&idSeg)
			if err != nil {
				return err
			}
			result = append(result, idSeg)
		}

		placeholdersDelete := make([]string, len(result)-1)

		for i := 0; i < len(result)-1; i++ {
			placeholdersDelete[i] = fmt.Sprintf("$%d", i+2)
		}

		log.Println(strings.Join(placeholdersDelete, ", "))

		query = fmt.Sprintf(`DELETE FROM "UserSegment" WHERE id_user = $1 AND id_segment IN (%s)`,
			strings.Join(placeholdersDelete, ", "))

		log.Println(query)

		_, err = u.db.Exec(query, result...)
		if err != nil {
			return err
		}
	}

	return nil

}
