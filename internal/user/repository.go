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
