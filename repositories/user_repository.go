package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ssjlee93/fitworks-data-user/models"
)

const (
	readOneUser  = "SELECT * FROM users u left join roles r on u.role_id = r.role_id WHERE u.user_id=$1"
	readAllUsers = "SELECT * FROM users u left join roles r on u.role_id = r.role_id ORDER BY u.user_id DESC"
	createUser   = "INSERT INTO users (first_name, last_name, google, apple, role_id, trainer_id) VALUES ($1, $2, $3, $4, $5, $6)"
	updateUser   = "UPDATE users SET first_name=$2, last_name=$3, google=$4, apple=$5, role_id=$6, trainer_id=$7, updated=current_timestamp WHERE user_id=$1"
	deleteUser   = "DELETE FROM users WHERE user_id=$1"
)

type UserRepository struct {
	d *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{d: db}
}

func (userRepository *UserRepository) ReadAll() ([]models.User, error) {
	log.Println("| - - UserRepository.ReadAll")
	result := make([]models.User, 0)
	// query
	rows, err := userRepository.d.Query(readAllUsers)
	if err != nil {
		return nil, fmt.Errorf("users : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user models.User
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

func (userRepository *UserRepository) ReadOne(id int64) (*models.User, error) {
	log.Println("| - - UserRepository.ReadOne")
	// query
	row := userRepository.d.QueryRow(readOneUser, id)
	result, err := scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("ReadOne: %v", err)
	}
	return result, nil
}

func (userRepository *UserRepository) Create(user models.User) error {
	log.Println("| - - UserRepository.Create")

	exec, err := userRepository.d.Exec(createUser,
		user.FirstName,
		user.LastName,
		user.Google,
		user.Apple,
		user.RoleID,
		user.TrainerID)

	if err != nil {
		log.Printf("Error UserRepository create : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func (userRepository *UserRepository) Update(user models.User) error {
	log.Println("| - - UserRepository.Update")
	exec, err := userRepository.d.Exec(updateUser,
		user.UserID,
		user.FirstName,
		user.LastName,
		user.Google,
		user.Apple,
		user.RoleID,
		user.TrainerID)

	if err != nil {
		log.Printf("Error UserRepository update : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func (userRepository *UserRepository) Delete(id int64) error {
	log.Println("| - - UserRepository Delete", id)
	exec, err := userRepository.d.Exec(deleteUser, id)
	if err != nil {
		log.Printf("UserRepository.Delete error : %v", err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func scanUser(row *sql.Row) (*models.User, error) {
	var user models.User
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
