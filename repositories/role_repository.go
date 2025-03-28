package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ssjlee93/fitworks-data-user/models"
)

const (
	readOneRole  = "SELECT * FROM roles WHERE role_id=$1"
	readAllRoles = "SELECT * FROM roles"
	createRole   = "INSERT INTO roles (role) VALUES ($1) RETURNING *"
	updateRole   = "UPDATE roles SET role=$1, updated=current_timestamp WHERE role_id=$2 RETURNING *"
	deleteRole   = "DELETE FROM roles WHERE role_id=$1 RETURNING *"
)

type RoleRepository struct {
	d *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{d: db}
}

func (roleRepository *RoleRepository) ReadAll() ([]models.Role, error) {
	result := make([]models.Role, 0)
	// query
	rows, err := roleRepository.d.Query(readAllRoles)
	if err != nil {
		return nil, fmt.Errorf("roles : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var role models.Role
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

func (roleRepository *RoleRepository) ReadOne(id int64) (*models.Role, error) {
	// query
	row := roleRepository.d.QueryRow(readOneRole, id)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Delete: %v", err)
	}
	return result, nil
}

func (roleRepository *RoleRepository) Create(role models.Role) (*models.Role, error) {
	// query
	row := roleRepository.d.QueryRow(createRole, role.Role)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Delete: %v", err)
	}
	return result, nil
}

func (roleRepository *RoleRepository) Update(role models.Role) (*models.Role, error) {
	// query
	row := roleRepository.d.QueryRow(updateRole, role.Role, role.RoleID)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Update: %v", err)
	}
	return result, nil
}

func (roleRepository *RoleRepository) Delete(id int64) (*models.Role, error) {
	// query
	row := roleRepository.d.QueryRow(deleteRole, id)
	// scan result
	result, err := scanRole(row)
	if err != nil {
		return nil, fmt.Errorf("Delete: %v", err)
	}
	return result, nil
}

func scanRole(row *sql.Row) (*models.Role, error) {
	// read each row and load it into the entity
	var result models.Role
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
