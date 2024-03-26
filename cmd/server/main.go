package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/MechiBakker/BE3-FINAL/cmd/server/handler"
	"github.com/MechiBakker/BE3-FINAL/internal/odontologo"
	"github.com/MechiBakker/BE3-FINAL/internal/paciente"
	"github.com/MechiBakker/BE3-FINAL/internal/turno"
	"github.com/MechiBakker/BE3-FINAL/pkg/store"
	"github.com/gin-gonic/gin"
)

// @title Api de turnos Odontologicos
// @version 1.0
// @description Reserva de turnos odontológicos
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	db, err := sql.Open("mysql", "root:root@/my_db")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewSqlStore(db)

	repo := odontologo.NewRepository(storage)
	service := odontologo.NewService(repo)
	odontologoHandler := handler.NewOdontologoHandler(service)

	repoPaciente := paciente.NewRepository(storage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	repoTurno := turno.NewRepository(storage)
	serviceTurno := turno.NewService(repoTurno)
	turnoHandler := handler.NewTurnoHandler(serviceTurno)

	engine := gin.Default()

	engine.SetTrustedProxies([]string{"127.0.0.1"})

	engine.GET("/api/v1/ping", func(c *gin.Context) { c.String(200, "pong") })

	odontologos := engine.Group("/api/v1/odontologos")
	{
		odontologos.POST("", odontologoHandler.CreateOdontologo())
		odontologos.GET(":idOdontologo", odontologoHandler.GetOdontologoByID())
		odontologos.PUT(":idOdontologo", odontologoHandler.UpdateOdontologo())
		odontologos.PATCH(":idOdontologo", odontologoHandler.UpdateOdontologoForField())
		odontologos.DELETE(":idOdontologo", odontologoHandler.DeleteOdontologo())
	}

	pacientes := engine.Group("/api/v1/pacientes")
	{
		pacientes.POST("", pacienteHandler.CreatePaciente())
		pacientes.GET(":idPaciente", pacienteHandler.GetPacienteByID())
		pacientes.PUT(":idPaciente", pacienteHandler.UpdatePaciente())
		pacientes.PATCH(":idPaciente", pacienteHandler.UpdatePacienteForField())
		pacientes.DELETE(":idPaciente", pacienteHandler.DeletePaciente())

	}

	turnos := engine.Group("/api/v1/turnos")
	{
		turnos.POST("", turnoHandler.CreateTurno())
		turnos.GET(":idTurno", turnoHandler.GetTurnoByID())
		turnos.PUT(":idTurno", turnoHandler.UpdateTurno())
		turnos.PATCH(":idTurno", turnoHandler.UpdateTurnoForField())
		turnos.DELETE(":idTurno", turnoHandler.DeleteTurno())
	}

	engine.Run(":8080")

}
