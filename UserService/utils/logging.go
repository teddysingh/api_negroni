package utils

import (
	"github.com/urfave/negroni"
)

// Logger - common logger to use across application
var Logger = negroni.NewLogger()
