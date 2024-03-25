package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/MechiBakker/BE3-FINAL/internal/domain"
)

type jsonStore struct {
	pathToFile string
}

func (s *jsonStore) loadOdontologos() ([]domain.Odontologo, error) {
	var odontologos []domain.Odontologo
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &odontologos)
	if err != nil {
		return nil, err
	}
	return odontologos, nil
}

func (s *jsonStore) saveOdontologos(odontologos []domain.Odontologo) error {
	bytes, err := json.Marshal(odontologos)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewJsonStore(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore) Read(id int) (domain.Odontologo, error) {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return domain.Odontologo{}, err
	}
	for _, odontologo := range odontologos {
		if odontologo.IdOdontologo == id {
			return odontologo, nil
		}
	}
	return domain.Odontologo{}, errors.New("El odontólogo no existe")
}

func (s *jsonStore) Create(odontologo domain.Odontologo) error {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return err
	}
	odontologo.IdOdontologo = len(odontologos) + 1
	odontologos = append(odontologos, odontologo)
	return s.saveOdontologos(odontologos)
}

func (s *jsonStore) Update(odontologo domain.Odontologo) error {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return err
	}
	for i, p := range odontologos {
		if p.IdOdontologo == odontologo.IdOdontologo {
			odontologos[i] = odontologo
			return s.saveOdontologos(odontologos)
		}
	}
	return errors.New("Ha ocurrido un error al actualizar odontólogo")
}

func (s *jsonStore) Delete(id int) error {
	odontologos, err := s.loadOdontologos()
	if err != nil {
		return err
	}
	for i, p := range odontologos {
		if p.IdOdontologo == id {
			odontologos = append(odontologos[:i], odontologos[i+1:]...)
			return s.saveOdontologos(odontologos)
		}
	}
	return errors.New("Ha ocurrido un error al eliminar odontólogo")
}

func (s *jsonStore) loadPacientes() ([]domain.Paciente, error) {
	var pacientes []domain.Paciente
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &pacientes)
	if err != nil {
		return nil, err
	}
	return pacientes, nil
}

func (s *jsonStore) savePacientes(pacientes []domain.Paciente) error {
	bytes, err := json.Marshal(pacientes)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewJsonStorePaciente(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore) ReadPaciente(id int) (domain.Paciente, error) {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return domain.Paciente{}, err
	}
	for _, paciente := range pacientes {
		if paciente.IdPaciente == id {
			return paciente, nil
		}
	}
	return domain.Paciente{}, errors.New("El paciente no existe")
}

func (s *jsonStore) CreatePaciente(paciente domain.Paciente) error {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return err
	}
	paciente.IdPaciente = len(pacientes) + 1
	pacientes = append(pacientes, paciente)
	return s.savePacientes(pacientes)
}

func (s *jsonStore) UpdatePaciente(paciente domain.Paciente) error {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return err
	}
	for i, p := range pacientes {
		if p.IdPaciente == paciente.IdPaciente {
			pacientes[i] = paciente
			return s.savePacientes(pacientes)
		}
	}
	return errors.New("Ha ocurrido un error al actualizar paciente")
}

func (s *jsonStore) DeletePaciente(id int) error {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return err
	}
	for i, p := range pacientes {
		if p.IdPaciente == id {
			pacientes = append(pacientes[:i], pacientes[i+1:]...)
			return s.savePacientes(pacientes)
		}
	}
	return errors.New("Ha ocurrido un error al eliminar paciente")
}

func (s *jsonStore) loadTurnos() ([]domain.Turno, error) {
	var turnos []domain.Turno
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &turnos)
	if err != nil {
		return nil, err
	}
	return turnos, nil
}

func (s *jsonStore) saveTurnos(turnos []domain.Turno) error {
	bytes, err := json.Marshal(turnos)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewJsonStoreTurno(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore) ReadTurno(id int) (domain.Turno, error) {
	turnos, err := s.loadTurnos()
	if err != nil {
		return domain.Turno{}, err
	}
	for _, turno := range turnos {
		if turno.IdTurno == id {
			return turno, nil
		}
	}
	return domain.Turno{}, errors.New("El turno no existe")
}

func (s *jsonStore) CreateTurno(turno domain.Turno) error {
	turnos, err := s.loadTurnos()
	if err != nil {
		return err
	}
	turno.IdTurno = len(turnos) + 1
	turnos = append(turnos, turno)
	return s.saveTurnos(turnos)
}

func (s *jsonStore) UpdateTurno(turno domain.Turno) error {
	turnos, err := s.loadTurnos()
	if err != nil {
		return err
	}
	for i, p := range turnos {
		if p.IdTurno == turno.IdTurno {
			turnos[i] = turno
			return s.saveTurnos(turnos)
		}
	}
	return errors.New("Ha ocurrido un error al actualizar turno")
}

func (s *jsonStore) DeleteTurno(id int) error {
	turnos, err := s.loadTurnos()
	if err != nil {
		return err
	}
	for i, p := range turnos {
		if p.IdTurno == id {
			turnos = append(turnos[:i], turnos[i+1:]...)
			return s.saveTurnos(turnos)
		}
	}
	return errors.New("Ha ocurrido un error al eliminar turno")
}
