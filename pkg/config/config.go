package config

import (
	"html/template"
	"log"
)

// AppConfig is the configuration for the application
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
