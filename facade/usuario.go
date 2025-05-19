package facade

import (
	"LearnGo/domain"
	"LearnGo/repository"
	"context"
	"database/sql"
)

type UsuarioFacade interface {
	// InserirUsuario insere um novo usuário no sistema.
	InserirUsuario(ctx context.Context, user domain.Usuario) (int, error)
}

type usuarioFacade struct {
	repo repository.Repository
	db   *sql.DB
}

func NewUsuarioFacade(repo repository.Repository, db *sql.DB) UsuarioFacade {
	return &usuarioFacade{repo: repo, db: db}
}

func (f *usuarioFacade) InserirUsuario(ctx context.Context, user domain.Usuario) (int, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	ID, err := f.repo.InserirUsuario(ctx, tx, user)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return ID, nil
}
