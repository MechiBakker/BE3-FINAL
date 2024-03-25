package store

import "github.com/MechiBakker/BE3-FINALO/internal/domain"

type StoreInterface interface {
	Read(id int) (domain.Odontologo, error)

	Create(odontologo domain.Odontologo) error

	Update(odontologo domain.Odontologo) error

	Delete(id int) error

	ReadPaciente(id int) (domain.Paciente, error)

	CreatePaciente(paciente domain.Paciente) error

	UpdatePaciente(paciente domain.Paciente) error

	DeletePaciente(id int) error

	ReadTurno(id int) (domain.Turno, error)

	CreateTurno(paciente domain.Turno) error

	UpdateTurno(paciente domain.Turno) error

	DeleteTurno(id int) error
}
