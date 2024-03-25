package turno

import (
	"github.com/MechiBakker/BE3-FINAL/internal/domain"
)

type Service interface {
	GetTurnoByID(id int) (domain.Turno, error)

	CreateTurno(p domain.Turno) (domain.Turno, error)

	DeleteTurno(id int) error

	UpdateTurno(id int, p domain.Turno) (domain.Turno, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateTurno(p domain.Turno) (domain.Turno, error) {
	p, err := s.r.CreateTurno(p)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}

func (s *service) GetTurnoByID(id int) (domain.Turno, error) {
	p, err := s.r.GetTurnoByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}

func (s *service) UpdateTurno(id int, u domain.Turno) (domain.Turno, error) {
	p, err := s.r.GetTurnoByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if u.DescripcionTurno != "" {
		p.DescripcionTurno = u.DescripcionTurno
	}
	if u.FechaTurno != "" {
		p.FechaTurno = u.FechaTurno
	}
	if u.IdOdontologo != "" {
		p.IdOdontologo = u.IdOdontologo
	}
	if u.IdPaciente != "" {
		p.IdPaciente = u.IdPaciente
	}
	p, err = s.r.UpdateTurno(id, p)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}

func (s *service) DeleteTurno(id int) error {
	err := s.r.DeleteTurno(id)
	if err != nil {
		return err
	}
	return nil
}
