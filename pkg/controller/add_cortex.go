package controller

import (
	"github.com/VineethReddy02/goModules/cortex-operator/pkg/controller/cortex"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cortex.Add)
}
