package controller

import (
	"github.com/briantopping/freeipa-operator/pkg/controller/freeipaservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, freeipaservice.Add)
}
