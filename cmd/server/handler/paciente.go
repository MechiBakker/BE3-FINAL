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
// @Summary Crear un nuevo paciente
// @Description Crea un nuevo paciente
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param paciente body domain.Paciente true "Datos del paciente a crear"
// @Success 201 {object} domain.Paciente "Paciente creado correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Router /pacientes [post]
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
// @Summary Obtener un paciente por ID
// @Description Obtiene un paciente por su ID
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param idPaciente path int true "ID del paciente a obtener"
// @Success 200 {object} domain.Paciente "Paciente obtenido correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el paciente"
// @Router /pacientes/{idPaciente} [get]
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
// @Summary Actualizar un paciente por ID
// @Description Actualiza un paciente por su ID con todos los campos
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param idPaciente path int true "ID del paciente a actualizar"
// @Param paciente body domain.Paciente true "Datos del paciente a actualizar"
// @Success 200 {object} domain.Paciente "Paciente actualizado correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el ID indicado"
// @Failure 409 {object} ErrorResponse "Error al actualizar el paciente"
// @Router /pacientes/{idPaciente} [put]
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
// @Summary Actualizar parcialmente un paciente por ID
// @Description Actualiza parcialmente un paciente por su ID con campos específicos
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param idPaciente path int true "ID del paciente a actualizar"
// @Param paciente body UpdateRequest true "Campos a actualizar"
// @Success 200 {object} domain.Paciente "Paciente actualizado correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el ID indicado"
// @Failure 409 {object} ErrorResponse "Error al actualizar el paciente"
// @Router /pacientes/{idPaciente} [patch]
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
// @Summary Eliminar un paciente por ID
// @Description Elimina un paciente por su ID
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param idPaciente path int true "ID del paciente a eliminar"
// @Success 200 {string} string "El paciente ha sido correctamente eliminado"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el ID indicado"
// @Router /pacientes/{idPaciente} [delete]
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
