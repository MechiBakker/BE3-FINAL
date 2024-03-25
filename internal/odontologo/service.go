package odontologo

import (
	"github.com/MechiBakker/BE3-FINAL/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Odontologo, error)

	Create(p domain.Odontologo) (domain.Odontologo, error)

	Delete(id int) error

	Update(id int, p domain.Odontologo) (domain.Odontologo, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(p domain.Odontologo) (domain.Odontologo, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return p, nil
}

func (s *service) GetByID(id int) (domain.Odontologo, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return p, nil
}

func (s *service) Update(id int, u domain.Odontologo) (domain.Odontologo, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	if u.NombreOdontologo != "" {
		p.NombreOdontologo = u.NombreOdontologo
	}
	if u.ApellidoOdontologo != "" {
		p.ApellidoOdontologo = u.ApellidoOdontologo
	}
	if u.MatriculaOdontologo != "" {
		p.MatriculaOdontologo = u.MatriculaOdontologo
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
