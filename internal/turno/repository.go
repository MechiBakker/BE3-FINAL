package turno

import (
	"errors"

	"github.com/MechiBakker/BE3-FINAL/internal/domain"
	"github.com/MechiBakker/BE3-FINAL/pkg/store"
)

type Repository interface {
	GetTurnoByID(id int) (domain.Turno, error)

	CreateTurno(p domain.Turno) (domain.Turno, error)

	UpdateTurno(id int, p domain.Turno) (domain.Turno, error)

	DeleteTurno(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) CreateTurno(p domain.Turno) (domain.Turno, error) {

	err := r.storage.CreateTurno(p)
	if err != nil {
		return domain.Turno{}, errors.New("Ha ocurrido un error al crear turno")
	}
	return p, nil
}

func (r *repository) GetTurnoByID(id int) (domain.Turno, error) {
	turno, err := r.storage.ReadTurno(id)
	if err != nil {
		return domain.Turno{}, errors.New("El turno no existe")
	}
	return turno, nil
}

func (r *repository) UpdateTurno(id int, p domain.Turno) (domain.Turno, error) {
	err := r.storage.UpdateTurno(p)
	if err != nil {
		return domain.Turno{}, errors.New("Ha ocurrido un error al actualizar turno")
	}
	return p, nil
}

func (r *repository) DeleteTurno(id int) error {
	err := r.storage.DeleteTurno(id)
	if err != nil {
		return err
	}
	return nil
}
