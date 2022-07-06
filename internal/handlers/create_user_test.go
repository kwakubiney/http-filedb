package handlers_test

import (
	"testing"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/sqlite-go/internal/engine"
	"github.com/sqlite-go/internal/handlers"
	"github.com/sqlite-go/internal/server"
	"github.com/stretchr/testify/assert"
)

var routeHandlers *gin.Engine

func TestCreateUser(t *testing.T){
	DB, err := engine.DbOpen("", "", 0644, "test")
	if err != nil {
		log.Println(err)
	}
	handler := handlers.New(DB)
	server := server.New(handler)
	routeHandlers = server.SetupRoutes()

	req := handlers.MakeTestRequest(t, "/createUser", map[string]interface{}{
		"id": "23",
		"username": "keb",
		"email": "k@mail.com", 
	}, "POST")


	resp := handlers.BootstrapServer(req, routeHandlers)
	
	decodedResponse := handlers.DecodeResponse(t,resp)
	assert.Equal(t, map[string]interface {}{"ID":"23", "Username":"keb", "Email":"k@mail.com"}, decodedResponse["data"])
	assert.Equal(t, "record has successfully been created", decodedResponse["message"])

	DB.TestClose()
}