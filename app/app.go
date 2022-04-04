package app

import (
	"github.com/nicolito128/tasks-api/app/routes"
)

func Init() {
	_ = routes.Run()
}
