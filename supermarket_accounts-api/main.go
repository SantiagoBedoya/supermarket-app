package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/SantiagoBedoya/supermarket_accounts-api/accounts"
	"github.com/SantiagoBedoya/supermarket_accounts-api/api"
	"github.com/SantiagoBedoya/supermarket_accounts-api/repositories/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := mysql.NewMySQLRepository(os.Getenv("MYSQL_URI"))
	service := accounts.NewService(repository)
	handler := api.NewHandler(service)

	router := gin.Default()

	apiRoutes := router.Group("/api/v1")
	{
		apiRoutes.POST("/sign-in", handler.SignIn)
		apiRoutes.POST("/sign-up", handler.SignUp)
		apiRoutes.GET("/verify-token", handler.VerifyToken)
	}

	errs := make(chan error, 2)
	go func() {
		log.Println("Server is running on", handlePort())
		errs <- http.ListenAndServe(handlePort(), router)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated: %s", <-errs)
}

func handlePort() string {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
