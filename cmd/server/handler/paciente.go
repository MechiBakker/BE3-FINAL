package handler

import (
	"errors"
	"strconv"

	"github.com/MechiBakker/BE3-FINAL/internal/domain"
	"github.com/MechiBakker/BE3-FINAL/internal/paciente"
	"github.com/MechiBakker/BE3-FINAL/pkg/web"
	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// POST
func (h *pacienteHandler) CreatePaciente() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente

		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		p, err := h.s.CreatePaciente(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "El paciente ha sido creado correctamente")
	}
}

// GET
func (h *pacienteHandler) GetPacienteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		paciente, err := h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el paciente"))
			return
		}
		web.Success(c, 200, paciente, "El paciente se ha encontrado por su ID")
	}
}

func validateFieldsPaciente(paciente *domain.Paciente) (bool, error) {
	switch {
	case paciente.NombrePaciente == "" || paciente.ApellidoPaciente == "" || paciente.DomicilioPaciente == "" || paciente.DniPaciente == "" || paciente.FechaDeAltaPaciente == "" /* || paciente.FechaDeAltaPaciente > time.Now() */ :
		return false, errors.New("No puede haber campos vacíos ni la fecha puede ser posterior a la actual")
	}
	return true, nil
}

// PUT
func (h *pacienteHandler) UpdatePaciente() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el ID indicado"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		valid, err := validateFieldsPaciente(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdatePaciente(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "El paciente ha sido correctamente actualizado")
	}
}

// PATCH
func (h *pacienteHandler) UpdatePacienteForField() gin.HandlerFunc {
	type Request struct {
		NombrePaciente      string `json:"nombrePaciente,omitempty"`
		ApellidoPaciente    string `json:"apellidoPaciente,omitempty"`
		DomicilioPaciente   string `json:"domicilioPaciente,omitempty"`
		DniPaciente         string `json:"dniPaciente,omitempty"`
		FechaDeAltaPaciente string `json:"fechaDeAltaPacientePaciente,omitempty"`
	}
	return func(c *gin.Context) {

		var r Request
		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el ID indicado"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		update := domain.Paciente{
			NombrePaciente:      r.NombrePaciente,
			ApellidoPaciente:    r.ApellidoPaciente,
			DomicilioPaciente:   r.DomicilioPaciente,
			DniPaciente:         r.DniPaciente,
			FechaDeAltaPaciente: r.FechaDeAltaPaciente,
		}
		p, err := h.s.UpdatePaciente(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "El paciente ha sido correctamente actualizado")
	}
}

// DELETE
func (h *pacienteHandler) DeletePaciente() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("idPaciente")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetPacienteByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el ID indicado"))
			return
		}
		err = h.s.DeletePaciente(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "El paciente ha sido correctamente eliminado")
	}
}
