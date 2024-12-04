package daos

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ssjlee93/fitworks-data-user/dtos"
)

const (
	readOneRole  = "SELECT * FROM roles WHERE role_id=$1"
	readAllRoles = "SELECT * FROM roles"
	createRole   = "INSERT INTO roles (role) VALUES ($1) RETURNING *"
	updateRole   = "UPDATE roles SET role=$1, updated=current_timestamp WHERE role_id=$2 RETURNING *"
	deleteRole   = "DELETE FROM roles WHERE role_id=$1 RETURNING *"
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
	rows, err := roleDao.d.Query(readAllRoles)
	if err != nil {
		return nil, fmt.Errorf("roles : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var role dtos.Role
		if err := rows.Scan(
			&role.RoleID,
			&role.Role,
			&role.Created,
			&role.Updated,
		); err != nil {
			return nil, fmt.Errorf("roles : %v", err)
		}
		result = append(result, role)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("roles : %v", err)
	}
	return result, nil
}

func (roleDao *RoleDAOImpl) ReadOne(id int64) (*dtos.Role, error) {
	// query
	row := roleDao.d.QueryRow(readOneRole, id)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Delete: %v", err)
	}
	return result, nil
}

func (roleDao *RoleDAOImpl) Create(role dtos.Role) (*dtos.Role, error) {
	// query
	row := roleDao.d.QueryRow(createRole, role.Role)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Delete: %v", err)
	}
	return result, nil
}

func (roleDao *RoleDAOImpl) Update(role dtos.Role) (*dtos.Role, error) {
	// query
	row := roleDao.d.QueryRow(updateRole, role.Role, role.RoleID)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Update: %v", err)
	}
	return result, nil
}

func (roleDao *RoleDAOImpl) Delete(id int64) (*dtos.Role, error) {
	// query
	row := roleDao.d.QueryRow(deleteRole, id)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Delete: %v", err)
	}
	return result, nil
}

func scanRole(row *sql.Row) (*dtos.Role, error) {
	// read each row and load it into the entity
	var result dtos.Role
	if err := row.Scan(
		&result.RoleID,
		&result.Role,
		&result.Created,
		&result.Updated,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("scanning row: no row found %v", err)
		}
		return nil, fmt.Errorf("scanning: %v", err)
	}
	return &result, nil
}
