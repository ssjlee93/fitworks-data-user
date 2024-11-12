package daos

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ssjlee93/fitworks-data-user/dtos"
	"log"
)

const (
	readOneQuery = "SELECT * FROM roles WHERE role_id=$1"
	readAllQuery = "SELECT * FROM roles;"
)

type RoleDAOImpl struct {
	d *sql.DB
}

func NewRoleDAOImpl(db *sql.DB) *RoleDAOImpl {
	return &RoleDAOImpl{d: db}
}

func (roleDao *RoleDAOImpl) ReadAll() ([]dtos.Role, error) {
	result := make([]dtos.Role, 0)
	// query
	rows, err := roleDao.d.Query(readAllQuery)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("roles : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var role dtos.Role
		if err := rows.Scan(&role.RoleID, &role.Role, &role.Created, &role.Updated); err != nil {
			log.Fatal(err)
			return nil, fmt.Errorf("roles : %v", err)
		}
		result = append(result, role)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("roles : %v", err)
	}
	return result, nil
}

func (roleDao *RoleDAOImpl) ReadOne(id uint) (dtos.Role, error) {
	var result dtos.Role
	// query
	row := roleDao.d.QueryRow(readOneQuery, id)
	if err := row.Scan(&result.RoleID, &result.Role, &result.Created, &result.Updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
	return result, nil
}
