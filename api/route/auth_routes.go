package route

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(*db, entities.CollectionUser)
	// sc := controller.SignupController{
	// 	SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
	// 	Env:           env,
	// }
	// group.POST("/signup", sc.Signup)
	// group.GET("/verify-email/:token", sc.VerifyEmail)
}

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(*db, entities.CollectionUser)
	// lc := &controller.LoginController{
	// 	LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
	// 	Env:          env,
	// }
	// group.POST("/login", lc.Login)
}

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(*db, entities.CollectionUser)
	// rtc := &controller.RefreshTokenController{
	// 	RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
	// 	Env:                 env,
	// }
	// group.POST("/refresh", rtc.RefreshToken)
}
