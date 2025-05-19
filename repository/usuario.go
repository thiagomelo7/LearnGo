package repository

import (
	"LearnGo/domain"
	"context"
	"database/sql"
)

type Repository interface {
	InserirUsuario(ctx context.Context, tx *sql.Tx, user domain.Usuario) (int, error)
}

type repositoryImpl struct{}

func (r *repositoryImpl) InserirUsuario(ctx context.Context, tx *sql.Tx, user domain.Usuario) (int, error) {
	insert, values, err := Psq.Insert("usuarios").
		Columns("nome", "email", "senha").
		Values(user.Nome, user.Email, user.Senha).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}
	var ID int
	err = tx.QueryRowContext(ctx, insert, values...).Scan(&ID)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
