package user

import (
	"database/sql"
	"fmt"

	"github.com/sakhimpungose/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db:db}
}


func (s *Store) GetUserByEmailAddress(emailAddress string) (*types.User, error){
	rows, err := s.db.Query("SELECT ID, FirstName, LastName, EmailAddress, Password, CreatedAt WHERE EmailAddress=?", emailAddress)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.EmailAddress,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}