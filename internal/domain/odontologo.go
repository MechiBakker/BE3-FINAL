package domain

type Odontologo struct {
	IdOdontologo        int    `json:"idOdontologo"`
	NombreOdontologo    string `json:"nombreOdontologo" binding:"required"`
	ApellidoOdontologo  string `json:"apellidoOdontologo" binding:"required"`
	MatriculaOdontologo string `json:"matriculaOdontologo" binding:"required"`
}
