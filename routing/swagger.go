package routing

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goshipdemo/docs" // docs is generated by Swag CLI, you have to import it.
)

func SwaggerRun(route *gin.Engine)  {
	// 引入swagger
	//url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	//route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}