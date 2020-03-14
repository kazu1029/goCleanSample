package iRepository

import "github.com/kazu1029/goCleanSample/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u *domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (nick_name, email) VALUES (?,?)",
		u.NickName, u.Email,
	)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

func (repo *UserRepository) FindByID(identifier int) (user domain.User, err error) {
	row, err := repo.Query(
		"SELECT id, nick_name, email FROM users WHERE id = ?",
		identifier,
	)
	defer row.Close()
	if err != nil {
		return
	}

	if err = row.Scan(&user.ID, &user.NickName, &user.Email); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users []domain.User, err error) {
	rows, err := repo.Query("SELECT id, nick_name, email FROM users")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.NickName, &user.Email); err != nil {
			continue
		}
		users = append(users, user)
	}

	return
}
