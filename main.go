package main

import (
	"gitee.com/fat_marmota/infra/log"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/factory"
	"github.com/dealmaker/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init begin
	log.InitZapSugared(true, false, 2)
	factory.BuildStreamlines()
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker", nil, "mysql")
	// Init end

	r := gin.Default()

	//store := memstore.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("user_info", store))

	r.POST("/auth/user/signup", handler.UserSignup)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
