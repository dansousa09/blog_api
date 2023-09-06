package handler

import (
	"github.com/dansousa09/blog_api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

const (
	errFieldIsRequiredMsg = "field is required"
	errPostNotFoundMsg    = "post not found"
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}
