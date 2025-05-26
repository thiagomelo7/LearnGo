package facade

import (
	"LearnGo/domain"
	"LearnGo/repository"
)

type UsuarioFacade struct {
	Repo repository.UsuarioRepository
}

var instance *UsuarioFacade

func Init(repo repository.UsuarioRepository) {
	instance = &UsuarioFacade{Repo: repo}
}

func Get() *UsuarioFacade {
	return instance
}

func (f UsuarioFacade) Create(u domain.Usuario) (int64, error) {
	return f.Repo.Insert(u)
}

func (f UsuarioFacade) Get(id int) (domain.Usuario, error) {
	return f.Repo.GetByID(id)
}

func (f UsuarioFacade) Update(u domain.Usuario) error {
	return f.Repo.Update(u)
}

func (f UsuarioFacade) Delete(id int) error {
	return f.Repo.Delete(id)
}
