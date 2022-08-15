package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig is the configuration for the application
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	ErrorLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
}