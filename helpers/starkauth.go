package helpers

import (
	"github.com/starkinfra/core-go/starkcore/user/project"
)

func Auth() project.Project {
	projectId := Env("PROJECT_ID")
	privateKey := Env("PRIVATE_KEY")

	return project.Project{
		Id:          projectId,
		PrivateKey:  privateKey,
		Environment: "sandbox",
	}
}
