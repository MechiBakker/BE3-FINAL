package store

import (
	"database/sql"
	"fmt"

	"github.com/MechiBakker/BE3-FINAL/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Create(odontologo domain.Odontologo) error {
	query := "INSERT INTO odontologos (idOdontologo, nombreOdontologo, apellidoOdontologo, matriculaOdontologo) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(odontologo)
	res, err := stmt.Exec(odontologo.IdOdontologo, odontologo.NombreOdontologo, odontologo.ApellidoOdontologo, odontologo.MatriculaOdontologo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Read(id int) (domain.Odontologo, error) {
	var odontologo domain.Odontologo
	query := "SELECT * FROM odontologos WHERE idOdontologo = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&odontologo.IdOdontologo, &odontologo.NombreOdontologo, &odontologo.ApellidoOdontologo, &odontologo.MatriculaOdontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s *sqlStore) Update(odontologo domain.Odontologo) error {
	query := "UPDATE odontologos SET nombreOdontologo = ?, apellidoOdontologo = ?, matriculaOdontologo = ? WHERE idOdontologo = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(odontologo.NombreOdontologo, odontologo.ApellidoOdontologo, odontologo.MatriculaOdontologo, odontologo.IdOdontologo)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {
	query := "DELETE FROM odontologos WHERE idOdontologo = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) CreatePaciente(paciente domain.Paciente) error {
	query := "INSERT INTO pacientes (idPaciente, nombrePaciente, apellidoPaciente, domicilioPaciente, dniPaciente, fechaDeAltaPaciente) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(paciente)
	res, err := stmt.Exec(paciente.IdPaciente, paciente.NombrePaciente, paciente.ApellidoPaciente, paciente.DomicilioPaciente, paciente.DniPaciente, paciente.FechaDeAltaPaciente)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) ReadPaciente(id int) (domain.Paciente, error) {
	var paciente domain.Paciente
	query := "SELECT * FROM pacientes WHERE idPaciente = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&paciente.IdPaciente, &paciente.NombrePaciente, &paciente.ApellidoPaciente, &paciente.DomicilioPaciente, &paciente.DniPaciente, &paciente.FechaDeAltaPaciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s *sqlStore) UpdatePaciente(paciente domain.Paciente) error {
	query := "UPDATE pacientes SET nombrePaciente = ?, apellidoPaciente = ?, domicilioPaciente = ?, dniPaciente = ?, fechaDeAltaPaciente = ? WHERE idPaciente = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(paciente.NombrePaciente, paciente.ApellidoPaciente, paciente.DomicilioPaciente, paciente.DniPaciente, paciente.FechaDeAltaPaciente, paciente.IdPaciente)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeletePaciente(id int) error {
	query := "DELETE FROM pacientes WHERE idPaciente = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) CreateTurno(turno domain.Turno) error {
	query := "INSERT INTO turnos (idTurno, descripcionTurno, fechaTurno, idOdontologo, idPaciente) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(turno)
	res, err := stmt.Exec(turno.IdTurno, turno.DescripcionTurno, turno.FechaTurno, turno.IdOdontologo, turno.IdPaciente)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) ReadTurno(id int) (domain.Turno, error) {
	var turno domain.Turno
	query := "SELECT * FROM turnos WHERE idTurno = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&turno.IdTurno, &turno.DescripcionTurno, &turno.FechaTurno, &turno.IdOdontologo, &turno.IdPaciente)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *sqlStore) UpdateTurno(turno domain.Turno) error {
	query := "UPDATE turnos SET descripcionTurno = ?, fechaTurno = ?, idOdontologo = ?, idPaciente = ? WHERE idTurno = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(turno.DescripcionTurno, turno.FechaTurno, turno.IdOdontologo, turno.IdPaciente, turno.IdTurno)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeleteTurno(id int) error {
	query := "DELETE FROM turnos WHERE idTurno = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
