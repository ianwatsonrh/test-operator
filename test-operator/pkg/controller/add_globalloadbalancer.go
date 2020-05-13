package controller

import (
	"github.com/ianwatsonrh/test-operator/pkg/controller/globalloadbalancer"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, globalloadbalancer.Add)
}
