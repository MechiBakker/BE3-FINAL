package odontologo

import (
	"errors"

	"github.com/MechiBakker/BE3-FINAL/internal/domain"
	"github.com/MechiBakker/BE3-FINAL/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Odontologo, error)

	Create(p domain.Odontologo) (domain.Odontologo, error)

	Update(id int, p domain.Odontologo) (domain.Odontologo, error)

	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(p domain.Odontologo) (domain.Odontologo, error) {

	err := r.storage.Create(p)
	if err != nil {
		return domain.Odontologo{}, errors.New("Ha ocurrido un error al crear odontólogo")
	}
	return p, nil
}

func (r *repository) GetByID(id int) (domain.Odontologo, error) {
	odontologo, err := r.storage.Read(id)
	if err != nil {
		return domain.Odontologo{}, errors.New("El odontólogo no existe")
	}
	return odontologo, nil
}

func (r *repository) Update(id int, p domain.Odontologo) (domain.Odontologo, error) {

	err := r.storage.Update(p)
	if err != nil {
		return domain.Odontologo{}, errors.New("Ha ocurrido un error al actualizar odontólogo")
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
