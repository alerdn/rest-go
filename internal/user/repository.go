package user

import (
	"database/sql"
	"log"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection,
	}
}

func (ur *UserRepository) Create(user User) (int, error) {
	result, err := ur.connection.Exec(
		"INSERT INTO user(name, email, password) VALUES (?, ?, ?);",
		user.Name, user.Email, user.Password,
	)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}

func (ur *UserRepository) GetUsers() ([]User, error) {
	rows, err := ur.connection.Query("SELECT id, created_at, name, email, password FROM user")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var userList []User
	var user User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.CreatedAt, &user.Name, &user.Email, &user.Password); err != nil {
			log.Println(err)
			return nil, err
		}

		userList = append(userList, user)
	}

	return userList, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (User, error) {
	var user User

	err := ur.connection.QueryRow("SELECT id, created_at, name, email, password FROM user WHERE email = ?", email).Scan(&user.ID, &user.CreatedAt, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(id int) (User, error) {
	var user User

	err := ur.connection.QueryRow("SELECT id, created_at, name, email, password FROM user WHERE id = ?", id).Scan(&user.ID, &user.CreatedAt, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return User{}, err
	}

	return user, nil
}