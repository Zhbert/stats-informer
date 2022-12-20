package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GitHubPage(context *gin.Context) {
	context.HTML(http.StatusOK, "github.tmpl", gin.H{"title": "GitHub Projects"})
}
