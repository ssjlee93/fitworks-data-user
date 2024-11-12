package daos

import (
	"database/sql"
	"fmt"
	"github.com/ssjlee93/fitworks-data-user/dtos"
	"log"
)

const (
	readOneQuery = "SELECT * FROM roles WHERE id=?"
	readAllQuery = "SELECT * FROM roles"
)

type RoleDAOImpl struct {
	Db *sql.DB
}

func NewRoleDAOImpl(db *sql.DB) *RoleDAOImpl {
	return &RoleDAOImpl{Db: db}
}

func (this *RoleDAOImpl) ReadAll() ([]dtos.Role, error) {
	result := make([]dtos.Role, 0)
	// query
	rows, err := this.Db.Query(readAllQuery)
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
