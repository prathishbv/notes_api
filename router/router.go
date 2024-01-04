package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prathishbv/notes-api/controller"
	middleware "github.com/prathishbv/notes-api/middlerware"
	"github.com/prathishbv/notes-api/repository"
)

func NewRouter(userRepository repository.UsersRepository, authenticationController *controller.AuthenticationController, usersController *controller.UserController, notesController *controller.NotesController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/signup", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.GetUsers)

	notesRouter := router.Group("/notes", middleware.DeserializeUser(userRepository))
	notesRouter.GET("", notesController.GetNotes)
	notesRouter.GET("/:id", notesController.GetNote)
	notesRouter.POST("", notesController.CreateNote)
	notesRouter.PUT("/:id", notesController.UpdateNote)
	notesRouter.DELETE("/:id", notesController.DeleteNote)
	notesRouter.POST("/:id/share", notesController.ShareNote)

	searchRouter := router.Group("/search", middleware.DeserializeUser(userRepository))
	searchRouter.GET("", notesController.SearchNotes)

	return service
}
