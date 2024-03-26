package handler

import (
	"errors"
	"strconv"

	"github.com/MechiBakker/BE3-FINAL/internal/domain"
	"github.com/MechiBakker/BE3-FINAL/internal/turno"
	"github.com/MechiBakker/BE3-FINAL/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.Service
}

func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// POST
// @Summary Obtener turno por ID
// @Description Retorna un turno dado su ID
// @Tags Turnos
// @Accept json
// @Produce json
// @Param idTurno path int true "ID del turno a obtener"
// @Success 200 {object} domain.Turno "Turno obtenido"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Router /turnos/{idTurno} [get]
func (h *turnoHandler) CreateTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		p, err := h.s.CreateTurno(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "El turno ha sido creado correctamente")
	}
}

// GET
// @Summary Obtener turno por ID
// @Description Retorna un turno dado su ID
// @Tags Turnos
// @Accept json
// @Produce json
// @Param idTurno path int true "ID del turno a obtener"
// @Success 200 {object} domain.Turno "Turno obtenido"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el turno"
// @Router /turnos/{idTurno} [get]
func (h *turnoHandler) GetTurnoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		turno, err := h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el turno indicado"))
			return
		}
		web.Success(c, 200, turno, "El turno se ha encontrado por su ID")
	}
}

func validateFieldsTurno(turno *domain.Turno) (bool, error) {
	switch {
	case turno.DescripcionTurno == "" || turno.FechaTurno == "" || turno.IdOdontologo == "" || turno.IdPaciente == "":
		return false, errors.New("No puede haber campos vacíos ni la fecha puede ser posterior a la actual")
	}
	return true, nil
}

// PUT
// @Summary Actualizar turno
// @Description Actualiza un turno existente por su ID
// @Tags Turnos
// @Accept json
// @Produce json
// @Param idTurno path int true "ID del turno a actualizar"
// @Param body body domain.Turno true "Información actualizada del turno"
// @Success 200 {object} domain.Turno "Turno actualizado"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Router /turnos/{idTurno} [put]
func (h *turnoHandler) UpdateTurno() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el ID indicado"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		valid, err := validateFieldsTurno(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.UpdateTurno(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "El turno ha sido correctamente actualizado")
	}
}

// PATCH
// @Summary Actualizar campos específicos de un turno
// @Description Actualiza campos específicos de un turno por su ID
// @Tags Turnos
// @Accept json
// @Produce json
// @Param idTurno path int true "ID del turno a actualizar"
// @Param body body Request true "Campos a actualizar del turno"
// @Success 200 {object} domain.Turno "Turno actualizado"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Router /turnos/{idTurno} [patch]
func (h *turnoHandler) UpdateTurnoForField() gin.HandlerFunc {
	type Request struct {
		DescripcionTurno string `json:"descripcionTurno,omitempty"`
		FechaTurno       string `json:"fechaTurno,omitempty"`
		IdOdontologo     string `json:"idOdontologo,omitempty"`
		IdPaciente       string `json:"idPaciente,omitempty"`
	}
	return func(c *gin.Context) {

		var r Request
		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el ID indicado"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		update := domain.Turno{
			DescripcionTurno: r.DescripcionTurno,
			FechaTurno:       r.FechaTurno,
			IdOdontologo:     r.IdOdontologo,
			IdPaciente:       r.IdPaciente,
		}
		p, err := h.s.UpdateTurno(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "El turno ha sido correctamente actualizado")
	}
}

// DELETE
// @Summary Eliminar un turno
// @Description Elimina un turno por su ID
// @Tags Turnos
// @Accept json
// @Produce json
// @Param idTurno path int true "ID del turno a eliminar"
// @Success 200 {string} string "El turno ha sido correctamente eliminado"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Router /turnos/{idTurno} [delete]
func (h *turnoHandler) DeleteTurno() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("idTurno")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetTurnoByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el ID indicado"))
			return
		}
		err = h.s.DeleteTurno(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "El turno ha sido correctamente eliminado")
	}
}
