package handlers_test

import (
	"testing"

	"log"
	"github.com/sqlite-go/internal/engine"
	"github.com/sqlite-go/internal/handlers"
	"github.com/sqlite-go/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestReadUser(t *testing.T){
	db, err := engine.DbOpen("", "", 0644, "test")
	if err != nil {
		log.Println(err)
	}
	handler := handlers.New(db)
	server := server.New(handler)
	routeHandlers = server.SetupRoutes()

	req := handlers.MakeTestRequest(t, "/createUser", map[string]interface{}{
		"id": "23",
		"username": "keb",
		"email": "k@mail.com", 
	}, "POST")


	resp := handlers.BootstrapServer(req, routeHandlers)
	
	decodedResponse := handlers.DecodeResponse(t,resp)
	assert.Equal(t, map[string]interface {}{"Email":"k@mail.com", "ID":"23", "Username":"keb"}, decodedResponse["data"])
	assert.Equal(t, "record has successfully been created", decodedResponse["message"])


	req = handlers.MakeTestRequest(t, "/readUser/23", map[string]interface{}{}, "GET")

	resp = handlers.BootstrapServer(req, routeHandlers)
	
	decodedResponse = handlers.DecodeResponse(t,resp)
	assert.Equal(t, "ID: 23, Name: keb, Email: k@mail.com", decodedResponse["data"])
	assert.Equal(t, "record has successfully been read", decodedResponse["message"])

	db.TestClose()
}


func TestReadUsers(t *testing.T){
	db, err := engine.DbOpen("", "", 0644, "test")
	if err != nil {
		log.Println(err)
	}
	handler := handlers.New(db)
	server := server.New(handler)
	routeHandlers = server.SetupRoutes()

	req := handlers.MakeTestRequest(t, "/createUser", map[string]interface{}{
		"id": "23",
		"username": "keb",
		"email": "k@mail.com", 
	}, "POST")


	resp := handlers.BootstrapServer(req, routeHandlers)
	
	decodedResponse := handlers.DecodeResponse(t,resp)
	assert.Equal(t, map[string]interface {}{"Email":"k@mail.com", "ID":"23", "Username":"keb"}, decodedResponse["data"])
	assert.Equal(t, "record has successfully been created", decodedResponse["message"])


	req = handlers.MakeTestRequest(t, "/readUser/23", map[string]interface{}{}, "GET")

	resp = handlers.BootstrapServer(req, routeHandlers)
	
	decodedResponse = handlers.DecodeResponse(t,resp)
	assert.Equal(t, "ID: 23, Name: keb, Email: k@mail.com", decodedResponse["data"])
	assert.Equal(t, "record has successfully been read", decodedResponse["message"])

	db.TestClose()
}
