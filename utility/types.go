package utility

import (
	"github.com/julienschmidt/httprouter"
)

type RoutesMapper map[string]httprouter.Handle
