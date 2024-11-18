package daos

const (
	readOneUser  = "SELECT * FROM users u left join roles r on u.role_id = r.role_id WHERE u.user_id=$1"
	readAllUsers = "SELECT * FROM users"
	createUser   = "INSERT INTO users (role) VALUES ($1) RETURNING *"
	updateUser   = "UPDATE users SET role=$1, updated=current_timestamp WHERE role_id=$3 RETURNING *"
	deleteUser   = "DELETE FROM users WHERE role_id=$1 RETURNING *"
)
