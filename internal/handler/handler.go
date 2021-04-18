package handler

import (
	"awesomeProject1/internal/service"
)

func InitHandler(nameService service.NameService) {
	InitName(nameService)
}
