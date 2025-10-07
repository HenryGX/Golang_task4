package HANDLER

import (
	"Golang_task4/SERVICE"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsCtrl struct{}

func (cc *CommentsCtrl) Create(c *gin.Context) {
	err := SERVICE.CreateComment(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (cc *CommentsCtrl) Query(c *gin.Context) {
	err := SERVICE.QueryComment(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
