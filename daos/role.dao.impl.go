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
	readAllQuery = "SELECT * FROM roles"
	createQuery  = "INSERT INTO roles (role) VALUES ($1) RETURNING *"
	updateQuery  = "UPDATE roles SET role=$1, updated=$2 WHERE role_id=$3 RETURNING *"
	deleteQuery  = "DELETE FROM roles WHERE role_id=$1 RETURNING *"
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
			log.Fatalf("roles : %v", err)
		}
		result = append(result, role)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("roles : %v", err)
	}
	return result, nil
}

func (roleDao *RoleDAOImpl) ReadOne(id int64) (*dtos.Role, error) {
	var result dtos.Role
	// query
	row := roleDao.d.QueryRow(readOneQuery, id)
	if err := row.Scan(&result.RoleID, &result.Role, &result.Created, &result.Updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("ReadOne: 0 rows found %v", err)
		}
		return nil, fmt.Errorf("ReadOne: %v", err)
	}
	return &result, nil
}

func (roleDao *RoleDAOImpl) Create(role dtos.Role) (*dtos.Role, error) {
	var result dtos.Role
	row := roleDao.d.QueryRow(createQuery, role.Role)
	if err := row.Scan(&result.RoleID, &result.Role, &result.Created, &result.Updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("create: no row created %v", err)
		}
		return nil, fmt.Errorf("create: %v", err)
	}
	return &result, nil
}

func (roleDao *RoleDAOImpl) Update(role dtos.Role) (*dtos.Role, error) {
	var result dtos.Role
	row := roleDao.d.QueryRow(updateQuery, role.Role, role.Updated, role.RoleID)
	if err := row.Scan(&result.RoleID, &result.Role, &result.Created, &result.Updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("update: no row updated %v", err)
		}
		return nil, fmt.Errorf("update: %v", err)
	}
	return &result, nil
}

func (roleDao *RoleDAOImpl) Delete(id int64) (*dtos.Role, error) {
	var result dtos.Role
	row := roleDao.d.QueryRow(deleteQuery, id)
	if err := row.Scan(&result.RoleID, &result.Role, &result.Created, &result.Updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("delete: no row deleted %v", err)
		}
		return nil, fmt.Errorf("delete: %v", err)
	}
	return &result, nil
}
