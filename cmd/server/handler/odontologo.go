package handler

import (
	"errors"
	"strconv"

	"github.com/MechiBakker/BE3-FINAL/internal/domain"
	"github.com/MechiBakker/BE3-FINAL/internal/odontologo"
	"github.com/MechiBakker/BE3-FINAL/pkg/web"
	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	s odontologo.Service
}

func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

// POST
// @Summary Crear un nuevo odontólogo
// @Description Crea un nuevo odontólogo con los datos proporcionados
// @Tags Odontólogos
// @Accept json
// @Produce json
// @Param odontologo body domain.Odontologo true "Datos del odontólogo a crear"
// @Success 201 {object} domain.Odontologo "Odontólogo creado correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Router /odontologos [post]
func (h *odontologoHandler) CreateOdontologo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domain.Odontologo

		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		p, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "El odontólogo ha sido creado correctamente")
	}
}

// GET
// @Summary Obtener un odontólogo por ID
// @Description Obtiene un odontólogo por su ID
// @Tags Odontólogos
// @Accept json
// @Produce json
// @Param idOdontologo path int true "ID del odontólogo a obtener"
// @Success 200 {object} domain.Odontologo "Odontólogo obtenido correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el ID indicado"
// @Router /odontologos/{idOdontologo} [get]
func (h *odontologoHandler) GetOdontologoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el odontólogo"))
			return
		}
		web.Success(c, 200, odontologo, "El odontólogo se ha encontrado por su ID")
	}
}

func validateEmptys(odontologo *domain.Odontologo) (bool, error) {
	switch {
	case odontologo.NombreOdontologo == "" || odontologo.ApellidoOdontologo == "" || odontologo.MatriculaOdontologo == "":
		return false, errors.New("No puede haber campos vacíos")
	}
	return true, nil
}

// PUT
// @Summary Actualizar un odontólogo por ID
// @Description Actualiza un odontólogo por su ID con todos los campos
// @Tags Odontólogos
// @Accept json
// @Produce json
// @Param idOdontologo path int true "ID del odontólogo a actualizar"
// @Param odontologo body domain.Odontologo true "Datos del odontólogo a actualizar"
// @Success 200 {object} domain.Odontologo "Odontólogo actualizado correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el ID indicado"
// @Failure 409 {object} ErrorResponse "Error al actualizar el odontólogo"
// @Router /odontologos/{idOdontologo} [put]
func (h *odontologoHandler) UpdateOdontologo() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado el ID indicado"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var odontologo domain.Odontologo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		p, err := h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "El odontólogo ha sido correctamente actualizado")
	}
}

// PATCH
// @Summary Actualizar parcialmente un odontólogo por ID
// @Description Actualiza parcialmente un odontólogo por su ID con campos específicos
// @Tags Odontólogos
// @Accept json
// @Produce json
// @Param idOdontologo path int true "ID del odontólogo a actualizar"
// @Param odontologo body UpdateRequest true "Campos a actualizar"
// @Success 200 {object} domain.Odontologo "Odontólogo actualizado correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el ID indicado"
// @Failure 409 {object} ErrorResponse "Error al actualizar el odontólogo"
// @Router /odontologos/{idOdontologo} [patch]
func (h *odontologoHandler) UpdateOdontologoForField() gin.HandlerFunc {
	type Request struct {
		NombreOdontologo    string `json:"nombreOdontologo,omitempty"`
		ApellidoOdontologo  string `json:"apellidoOdontologo,omitempty"`
		MatriculaOdontologo string `json:"matriculaOdontologo,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado ese ID"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("Invalid Json"))
			return
		}
		update := domain.Odontologo{
			NombreOdontologo:    r.NombreOdontologo,
			ApellidoOdontologo:  r.ApellidoOdontologo,
			MatriculaOdontologo: r.MatriculaOdontologo,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "El odontólogo ha sido correctamente actualizado")
	}
}

// DELETE
// @Summary Eliminar un odontólogo por ID
// @Description Elimina un odontólogo por su ID
// @Tags Odontólogos
// @Accept json
// @Produce json
// @Param idOdontologo path int true "ID del odontólogo a eliminar"
// @Success 200 {object} string "Odontólogo eliminado correctamente"
// @Failure 400 {object} ErrorResponse "Error de solicitud"
// @Failure 404 {object} ErrorResponse "No se encontró el ID indicado"
// @Router /odontologos/{idOdontologo} [delete]
func (h *odontologoHandler) DeleteOdontologo() gin.HandlerFunc {

	return func(c *gin.Context) {

		idParam := c.Param("idOdontologo")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se ha encontrado ese ID"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, nil, "El odontólogo ha sido correctamente eliminado")
	}
}
