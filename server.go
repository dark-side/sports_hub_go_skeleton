package main

import (
	"time"

	"go_be_plgrnd/config"
	"go_be_plgrnd/handler"
	"go_be_plgrnd/middleware"
	"go_be_plgrnd/repository"
	"go_be_plgrnd/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)

	authHandler handler.AuthHandler = handler.NewAuthHandler(authService, jwtService)
	userHandler handler.UserHandler = handler.NewUserHandler(userService, jwtService)

	articleRepository                    repository.ArticleRepository                    = repository.NewArticleRepository(db)
	commentRepository                    repository.CommentRepository                    = repository.NewCommentRepository(db)
	likeRepository                       repository.LikeRepository                       = repository.NewLikeRepository(db)
	activeStorageAttachmentRepository    repository.ActiveStorageAttachmentRepository    = repository.NewActiveStorageAttachmentRepository(db)
	activeStorageBlobRepository          repository.ActiveStorageBlobRepository          = repository.NewActiveStorageBlobRepository(db)
	activeStorageVariantRecordRepository repository.ActiveStorageVariantRecordRepository = repository.NewActiveStorageVariantRecordRepository(db)

	articleService                    service.ArticleService                    = service.NewArticleService(articleRepository)
	commentService                    service.CommentService                    = service.NewCommentService(commentRepository)
	likeService                       service.LikeService                       = service.NewLikeService(likeRepository)
	activeStorageAttachmentService    service.ActiveStorageAttachmentService    = service.NewActiveStorageAttachmentService(activeStorageAttachmentRepository)
	activeStorageBlobService          service.ActiveStorageBlobService          = service.NewActiveStorageBlobService(activeStorageBlobRepository)
	activeStorageVariantRecordService service.ActiveStorageVariantRecordService = service.NewActiveStorageVariantRecordService(activeStorageVariantRecordRepository)

	articleHandler       handler.ArticleHandler                    = handler.NewArticleHandler(articleService, jwtService)
	commentHandler       handler.CommentHandler                    = handler.NewCommentHandler(commentService, jwtService)
	likeHandler          handler.LikeHandler                       = handler.NewLikeHandler(likeService, jwtService)
	attachmentHandler    handler.ActiveStorageAttachmentHandler    = handler.NewActiveStorageAttachmentHandler(activeStorageAttachmentService, jwtService)
	blobHandler          handler.ActiveStorageBlobHandler          = handler.NewActiveStorageBlobHandler(activeStorageBlobService, jwtService)
	variantRecordHandler handler.ActiveStorageVariantRecordHandler = handler.NewActiveStorageVariantRecordHandler(activeStorageVariantRecordService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           365 * 24 * time.Hour,
	}))

	//TOD: double check the routes from documentation
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	articleRoutes := r.Group("api/articles", middleware.AuthorizeJWT(jwtService))
	{
		articleRoutes.POST("/", articleHandler.CreateArticle)
		articleRoutes.PUT("/:id", articleHandler.UpdateArticle)
		articleRoutes.DELETE("/:id", articleHandler.DeleteArticle)
		articleRoutes.GET("/:id", articleHandler.GetArticleByID)
		articleRoutes.GET("/", articleHandler.GetAllArticles)
	}

	commentRoutes := r.Group("api/comments", middleware.AuthorizeJWT(jwtService))
	{
		commentRoutes.POST("/", commentHandler.CreateComment)
		commentRoutes.PUT("/:id", commentHandler.UpdateComment)
		commentRoutes.DELETE("/:id", commentHandler.DeleteComment)
		commentRoutes.GET("/:id", commentHandler.GetCommentByID)
		commentRoutes.GET("/article/:articleId", commentHandler.GetCommentsByArticleID)
	}

	likeRoutes := r.Group("api/likes", middleware.AuthorizeJWT(jwtService))
	{
		likeRoutes.POST("/", likeHandler.CreateLike)
		likeRoutes.PUT("/:id", likeHandler.UpdateLike)
		likeRoutes.DELETE("/:id", likeHandler.DeleteLike)
		likeRoutes.GET("/:id", likeHandler.GetLikeByID)
		likeRoutes.GET("/likeable/:type/:id", likeHandler.GetLikesByLikeable)
	}

	attachmentRoutes := r.Group("api/attachments", middleware.AuthorizeJWT(jwtService))
	{
		attachmentRoutes.POST("/", attachmentHandler.CreateAttachment)
		attachmentRoutes.PUT("/:id", attachmentHandler.UpdateAttachment)
		attachmentRoutes.DELETE("/:id", attachmentHandler.DeleteAttachment)
		attachmentRoutes.GET("/:id", attachmentHandler.GetAttachmentByID)
		attachmentRoutes.GET("/record/:type/:id", attachmentHandler.GetAttachmentsByRecord)
	}

	blobRoutes := r.Group("api/blobs", middleware.AuthorizeJWT(jwtService))
	{
		blobRoutes.POST("/", blobHandler.CreateBlob)
		blobRoutes.PUT("/:id", blobHandler.UpdateBlob)
		blobRoutes.DELETE("/:id", blobHandler.DeleteBlob)
		blobRoutes.GET("/:id", blobHandler.GetBlobByID)
		blobRoutes.GET("/key/:key", blobHandler.GetBlobByKey)
	}

	variantRoutes := r.Group("api/variant-records", middleware.AuthorizeJWT(jwtService))
	{
		variantRoutes.POST("/", variantRecordHandler.CreateVariantRecord)
		variantRoutes.PUT("/:id", variantRecordHandler.UpdateVariantRecord)
		variantRoutes.DELETE("/:id", variantRecordHandler.DeleteVariantRecord)
		variantRoutes.GET("/:id", variantRecordHandler.GetVariantRecordByID)
		variantRoutes.GET("/blob/:blobId", variantRecordHandler.GetVariantRecordsByBlobID)
	}

	r.Run(":8080")
}
