package repository

import (
	"context"
	"log"
	"session-23-gin-jwt/internal/models"

	"github.com/jmoiron/sqlx"
)

type MysqlRepo struct {
	conn *sqlx.DB
}

func (m MysqlRepo) UpdateUser(ctx context.Context, id interface{}, user models.User) error {
	_, err := m.conn.ExecContext(ctx, `Update Users set firstName=?, secondName=?, username=?, password=? where id =?`, user.FirstName, user.SeconName, user.Username, user.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (m MysqlRepo) DeleteUser(ctx context.Context, id interface{}) error {
	_, err := m.conn.ExecContext(ctx, `DELETE FROM Users WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}

func (m MysqlRepo) CreateUser(ctx context.Context, user models.User) (interface{}, error) {

	result, err := m.conn.ExecContext(ctx, `INSERT INTO Users (ID, FirstName, SecondName, UserName, password) VALUES (?, ?, ?, ?, ?)`,
		user.ID, user.FirstName, user.SeconName, user.Username, user.Password)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return user.ID, nil

}

func (m MysqlRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User
	log.Println(userName)
	row := m.conn.QueryRowx("select Username,FirstName,SecondName from Users where userName=?", userName)
	err := row.Scan(&user.FirstName, &user.SeconName, &user.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (m MysqlRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	rows, err := m.conn.Query("SELECT id, firstName, secondName, userName FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.SeconName, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func NewMysqlReqo(db *sqlx.DB) DbRepository {
	return &MysqlRepo{
		db,
	}

}
