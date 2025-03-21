package daos

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ssjlee93/fitworks-data-user/dtos"
)

const (
	readOneUser  = "SELECT * FROM users u left join roles r on u.role_id = r.role_id WHERE u.user_id=$1"
	readAllUsers = "SELECT * FROM users u left join roles r on u.role_id = r.role_id ORDER BY u.user_id DESC"
	createUser   = "INSERT INTO users (first_name, last_name, google, apple, role_id, trainer_id) VALUES ($1, $2, $3, $4, $5, $6)"
	updateUser   = "UPDATE users SET first_name=$2, last_name=$3, google=$4, apple=$5, role_id=$6, trainer_id=$7, updated=current_timestamp WHERE user_id=$1"
	deleteUser   = "DELETE FROM users WHERE user_id=$1"
)

type UserDAOImpl struct {
	d *sql.DB
}

func NewUserDAOImpl(db *sql.DB) *UserDAOImpl {
	return &UserDAOImpl{d: db}
}

func (userDao *UserDAOImpl) ReadAll() ([]dtos.User, error) {
	log.Println("| - - - UserDAO.ReadAll")
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

	// handle any other errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("users : %v", err)
	}
	return result, nil
}

func (userDao *UserDAOImpl) ReadOne(id int64) (*dtos.User, error) {
	log.Println("| - - - UserDAO.ReadOne")
	// query
	row := userDao.d.QueryRow(readOneUser, id)
	result, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("ReadOne: %v", err)
	}
	return result, nil
}

func (userDao *UserDAOImpl) Create(user dtos.User) error {
	log.Println("| - - - UserDAO.Create")

	exec, err := userDao.d.Exec(createUser,
		user.FirstName,
		user.LastName,
		user.Google,
		user.Apple,
		user.RoleID,
		user.TrainerID)

	if err != nil {
		log.Printf("Error UserDao create : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func (userDao *UserDAOImpl) Update(user dtos.User) error {
	log.Println("| - - - UserDAO.Update")
	exec, err := userDao.d.Exec(updateUser,
		user.UserID,
		user.FirstName,
		user.LastName,
		user.Google,
		user.Apple,
		user.RoleID,
		user.TrainerID)

	if err != nil {
		log.Printf("Error UserDao update : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func (userDao *UserDAOImpl) Delete(id int64) error {
	log.Println("| - - - UserDAO Delete", id)
	exec, err := userDao.d.Exec(deleteUser, id)
	if err != nil {
		log.Printf("UserDAO.Delete error : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
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
