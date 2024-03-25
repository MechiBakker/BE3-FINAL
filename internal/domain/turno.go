package domain

type Turno struct {
	IdTurno          int    `json:"idTurno"`
	DescripcionTurno string `json:"descripcionTurno" binding:"required"`
	FechaTurno       string `json:"fechaTurno" binding:"required"`
	IdOdontologo     string `json:"idOdontologo" binding:"required"`
	IdPaciente       string `json:"idPaciente" binding:"required"`
}
