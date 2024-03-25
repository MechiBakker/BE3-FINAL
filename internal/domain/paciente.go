package domain

type Paciente struct {
	IdPaciente          int    `json:"idPaciente"`
	NombrePaciente      string `json:"nombrePaciente" binding:"required"`
	ApellidoPaciente    string `json:"apellidoPaciente" binding:"required"`
	DomicilioPaciente   string `json:"domicilioPaciente" binding:"required"`
	DniPaciente         string `json:"dniPaciente" binding:"required"`
	FechaDeAltaPaciente string `json:"fechaDeAltaPaciente" binding:"required"`
}
