package daos

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ssjlee93/fitworks-data-user/dtos"
)

const (
	readOneUser  = "SELECT * FROM users u left join roles r on u.role_id = r.role_id WHERE u.user_id=$1"
	readAllUsers = "SELECT * FROM users u left join roles r on u.role_id = r.role_id"
	createUser   = "INSERT INTO users (first_name, last_name, google, apple, role_id, trainer_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	updateUser   = "UPDATE users SET first_name=$1, last_name=$2, google=$3, apple=$4, role_id=$5, trainer_id=$6, updated=current_timestamp WHERE role_id=$3 RETURNING *"
	deleteUser   = "DELETE FROM users WHERE user_id=$1 RETURNING *"
)

type UserDAOImpl struct {
	d *sql.DB
}

func NewUserDAOImpl(db *sql.DB) *UserDAOImpl {
	return &UserDAOImpl{d: db}
}

func (userDao *UserDAOImpl) ReadAll() ([]dtos.User, error) {
	result := make([]dtos.User, 0)
	// query
	rows, err := userDao.d.Query(readAllUsers)
	if err != nil {
		return nil, fmt.Errorf("users : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user dtos.User
		if err := rows.Scan(
			&user.UserID,
			&user.FirstName,
			&user.LastName,
			&user.Google,
			&user.Apple,
			&user.RoleID,
			&user.TrainerID,
			&user.Created,
			&user.Updated,
			&user.Role.RoleID,
			&user.Role.Role,
			&user.Role.Created,
			&user.Role.Updated,
		); err != nil {
			return nil, fmt.Errorf("users : %v", err)
		}
		result = append(result, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("users : %v", err)
	}
	return result, nil
}

func (userDao *UserDAOImpl) ReadOne(id int64) (*dtos.User, error) {
	// query
	row := userDao.d.QueryRow(readOneUser, id)
	result, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("ReadOne: %v", err)
	}
	return result, nil
}

func (userDao *UserDAOImpl) Create(user dtos.User) (*dtos.User, error) {

	row := userDao.d.QueryRow(createUser,
		user.FirstName,
		user.LastName,
		user.Google,
		user.Apple,
		user.RoleID,
		user.TrainerID)

	result, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("Create: %v", err)
	}
	return result, nil
}

func (userDao *UserDAOImpl) Update(user dtos.User) (*dtos.User, error) {

	row := userDao.d.QueryRow(updateUser,
		user.FirstName,
		user.LastName,
		user.Google,
		user.Apple,
		user.RoleID,
		user.TrainerID)

	result, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("Update: %v", err)
	}
	return result, nil
}

func (userDao *UserDAOImpl) Delete(id int64) (*dtos.User, error) {

	row := userDao.d.QueryRow(deleteUser, id)

	result, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("Delete: %v", err)
	}
	return result, nil
}

func scanUser(row *sql.Row) (*dtos.User, error) {
	var user dtos.User
	if err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Google,
		&user.Apple,
		&user.RoleID,
		&user.TrainerID,
		&user.Created,
		&user.Updated,
		&user.Role.RoleID,
		&user.Role.Role,
		&user.Role.Created,
		&user.Role.Updated,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("scanning row failed: 0 rows found %v", err)
		}
		return nil, fmt.Errorf("scanning row failed: %v", err)
	}
	return &user, nil
}
