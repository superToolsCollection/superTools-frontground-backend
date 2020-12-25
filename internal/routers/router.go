package routers

import (
	"net/http"
	"time"

	_ "superTools-frontground-backend/docs"
	"superTools-frontground-backend/global"
	"superTools-frontground-backend/internal/dao"
	"superTools-frontground-backend/internal/middleware"
	"superTools-frontground-backend/internal/routers/api"
	"superTools-frontground-backend/internal/routers/api/sd"
	"superTools-frontground-backend/internal/routers/api/v1/bedtimeStory"
	"superTools-frontground-backend/internal/routers/api/v1/mall"
	"superTools-frontground-backend/internal/routers/api/v1/tools"
	"superTools-frontground-backend/internal/routers/api/v1/user"
	"superTools-frontground-backend/internal/service"
	"superTools-frontground-backend/pkg/limiter"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/**
* @Author: super
* @Date: 2020-08-21 21:14
* @Description:
**/

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	//放到需要token的请求中
	//r.Use(middleware.JWT())
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//获取token
	r.GET("/auth", api.GetAuth)
	RegisterController(r, USER, global.DBEngine)

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	tool := r.Group("api/v1")
	{
		tool.GET("/morse", tools.GetMorse)
		tool.GET("/qrcode", tools.GetQRcode)
		tool.GET("/rgb2hex", tools.RgbToHex)
	}

	return r
}

func RegisterController(r *gin.Engine, name string, db *gorm.DB) {
	switch name {
	case PRODUCT:
		registerProduct(r, db)
	case ORDER:
		registerOrder(r, db)
	case USER:
		registerUser(r, db)
	case BEDTIME:
		registerBedtime(r, db)
	case HEALTH:
		registerHealth(r, db)
	}
}

func registerHealth(r *gin.Engine, db *gorm.DB) {
	// The health check handlers
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
}

func registerBedtime(r *gin.Engine, db *gorm.DB) {
	story := bedtimeStory.NewStory()
	tag := bedtimeStory.NewTag()
	bedtime := r.Group("/api/v1/bedtime")
	bedtime.GET("/stories_only/:id", story.GetOnly)
	bedtime.Use(middleware.JWT())
	{
		bedtime.POST("/tags", tag.Create)
		bedtime.DELETE("/tags/:id", tag.Delete)
		bedtime.PUT("/tags/:id", tag.Update)
		bedtime.PATCH("/tags/:id/state", tag.Update)
		bedtime.GET("/tags", tag.List)

		bedtime.POST("/stories", story.Create)
		bedtime.DELETE("/stories/:id", story.Delete)
		bedtime.PUT("/stories/:id", story.Update)
		bedtime.PATCH("/stories/:id/state", story.Update)
		bedtime.GET("stories/:id", story.Get)
		bedtime.GET("/stories", story.List)
	}
}

func registerUser(r *gin.Engine, db *gorm.DB) {
	userManager := dao.NewUserManager("users", db)
	userService := service.NewUserService(userManager)
	userController := user.NewUserController(userService)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/signin", userController.SignIn)
		userGroup.POST("/register", userController.Register)
		userGroup.PUT("/update", userController.Update)
	}
}

func registerOrder(r *gin.Engine, db *gorm.DB) {
	orderManager := dao.NewOrderManager("orders", db)
	orderService := service.NewOrderService(orderManager)
	orderController := mall.NewOrderController(orderService)

	g := r.Group("/api/v1/mall")
	{
		g.GET("/orders/:id", orderController.GetOrder)
		g.GET("/all_orders", orderController.GetAllOrder)
		g.GET("/orders", orderController.GetOrderList)
		g.GET("/all_orders_user", orderController.GetOrderByUserID)
		g.GET("/orders_user", orderController.GetOrderListByUserID)
		g.POST("/orders", orderController.Insert)
		g.DELETE("/orders", orderController.Delete)
		g.PUT("/orders", orderController.Update)
	}
}

func registerProduct(r *gin.Engine, db *gorm.DB) {
	productManager := dao.NewProductManager("products", db)
	productService := service.NewProductService(productManager)
	productController := mall.NewProductController(productService)

	g := r.Group("/api/v1/mall")
	{
		g.GET("/products/:id", productController.GetProduct)
		g.GET("/all_products", productController.GetAllProduct)
		g.GET("/products", productController.GetProductList)
		g.POST("/products", productController.Insert)
		g.DELETE("/products", productController.Delete)
		g.PUT("/products", productController.Update)
	}
}
