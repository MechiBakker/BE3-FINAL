package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/MechiBakker/BE3-FINAL/cmd/server/handler"
	"github.com/MechiBakker/BE3-FINAL/internal/odontologo"
	"github.com/MechiBakker/BE3-FINAL/internal/paciente"
	"github.com/MechiBakker/BE3-FINAL/internal/turno"
	"github.com/MechiBakker/BE3-FINAL/pkg/store"
	"github.com/MechiBakker/BE3-FINAL/pkg/middleware"
	"github.com/MechiBakker/BE3-FINAL/cmd/server/docs"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Api de turnos Odontologicos
// @version 1.0
// @description Reserva de turnos odontológicos
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos_odontologia")
	if err != nil {
		log.Fatal(err)
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
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	docs.SwaggerInfo.Host = "localhost:8080"
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.SetTrustedProxies([]string{"127.0.0.1"})

	engine.GET("/api/v1/ping", func(c *gin.Context) { c.String(200, "pong") })

	odontologos := engine.Group("/api/v1/odontologos")
	{
		odontologos.POST("", odontologoHandler.CreateOdontologo())
		odontologos.GET(":idOdontologo", odontologoHandler.GetOdontologoByID())
		odontologos.PUT(":idOdontologo", middleware.Authentication(),odontologoHandler.UpdateOdontologo())
		odontologos.PATCH(":idOdontologo", middleware.Authentication(),odontologoHandler.UpdateOdontologoForField())
		odontologos.DELETE(":idOdontologo", middleware.Authentication(),odontologoHandler.DeleteOdontologo())
	
	}

	pacientes := engine.Group("/api/v1/pacientes")
	{
		pacientes.POST("", middleware.Authentication(),pacienteHandler.CreatePaciente())
		pacientes.GET(":idPaciente", pacienteHandler.GetPacienteByID())
		pacientes.PUT(":idPaciente", middleware.Authentication(),pacienteHandler.UpdatePaciente())
		pacientes.PATCH(":idPaciente", middleware.Authentication(),pacienteHandler.UpdatePacienteForField())
		pacientes.DELETE(":idPaciente", middleware.Authentication(),pacienteHandler.DeletePaciente())
	

	}

	turnos := engine.Group("/api/v1/turnos")
	{
		turnos.POST("", middleware.Authentication(),turnoHandler.CreateTurno())
		turnos.GET(":idTurno", turnoHandler.GetTurnoByID())
		turnos.PUT(":idTurno", middleware.Authentication(),turnoHandler.UpdateTurno())
		turnos.PATCH(":idTurno", middleware.Authentication(),turnoHandler.UpdateTurnoForField())
		turnos.DELETE(":idTurno", middleware.Authentication(),turnoHandler.DeleteTurno())
	}

	engine.Run(":8080")

}
