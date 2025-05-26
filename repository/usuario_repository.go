package repository

import (
	"LearnGo/domain"
	"database/sql"
)

type UsuarioRepository struct {
	DB *sql.DB
}

func (r UsuarioRepository) Insert(usuario domain.Usuario) (int64, error) {
	var id int64
	err := r.DB.QueryRow(
		"INSERT INTO usuarios (name, email) VALUES ($1, $2) RETURNING id",
		usuario.Name, usuario.Email,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r UsuarioRepository) GetByID(id int) (domain.Usuario, error) {
	var u domain.Usuario
	err := r.DB.QueryRow("SELECT id, name, email FROM usuarios WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return domain.Usuario{}, err
	}
	return u, nil
}

func (r UsuarioRepository) Update(u domain.Usuario) error {
	_, err := r.DB.Exec("UPDATE usuarios SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, u.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r UsuarioRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM usuarios WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
