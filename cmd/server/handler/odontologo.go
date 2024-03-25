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
