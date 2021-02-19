package service

import "github.com/nmarsollier/go_router_design/dao"

// Nos va a permitir mockear respuestas para los tests
var daoHelloFunc func() string = dao.Hello

// SayHello es nuestro negocio
func SayHello(userName string) string {
	return daoHelloFunc() + " " + userName
}
