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
		fmt.Println(rows.Columns())
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
			&user.Role.Updated); err != nil {
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
	var result dtos.User
	// query
	row := userDao.d.QueryRow(readOneUser, id)
	if err := row.Scan(&result.UserID,
		&result.FirstName,
		&result.LastName,
		&result.Google,
		&result.Apple,
		&result.Role.RoleID,
		&result.Role.Role,
		&result.Trainer,
		&result.Created,
		&result.Updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("ReadOne: 0 rows found %v", err)
		}
		return nil, fmt.Errorf("ReadOne: %v", err)
	}
	return &result, nil
}
