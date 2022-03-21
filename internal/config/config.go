package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// This configuration file might be access from any part of the application.
// We need to be very careful this configuration file does not import anything
// other than what it absolutely has to.

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
