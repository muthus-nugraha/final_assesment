package routers

import (
	"log"
	"os"

	"final_assignment/app/handler"
	"final_assignment/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var PORT = os.Getenv("PORT")

	UserHandler := handler.NewUserHandler()
	PhotoHandler := handler.NewPhotoHandler()
	CommentHandler := handler.NewCommentHandler()
	SocialMediaHandler := handler.NewSocialMediaHandler()
	r := gin.Default()

	r.POST("/User/Signup", UserHandler.Signup)
	r.POST("/User/Signin", UserHandler.Signin)
	routers := r.Group("")
	routers.Use(middleware.JwtAuthMiddleware())
	routers.DELETE("/User", UserHandler.RemoveUser)
	routers.PUT("/User", UserHandler.EditUser)

	routers.POST("/NewPhoto", PhotoHandler.NewPhoto)
	routers.GET("/Photos", PhotoHandler.GetPhotos)
	routers.PUT("/Photos/:photoId", PhotoHandler.EditPhoto)
	routers.DELETE("/Photos/:photoId", PhotoHandler.RemovePhoto)

	routers.POST("/NewComment", CommentHandler.NewComment)
	routers.GET("/Comments", CommentHandler.GetComments)
	routers.PUT("/Comments/:commentId", CommentHandler.EditComment)
	routers.DELETE("/Comments/:commentId", CommentHandler.RemoveComment)

	routers.POST("/NewSocialMedia", SocialMediaHandler.NewSocialMedia)
	routers.GET("/SocialMedia", SocialMediaHandler.GetSocialMedia)
	routers.PUT("/SocialMedia/:socialMediaId", SocialMediaHandler.EditSocialMedia)
	routers.DELETE("/SocialMedia/:socialMediaId", SocialMediaHandler.RemoveSocialMedia)

	log.Println("Start Server")
	r.Run(":" + PORT)
}
