package HANDLER

import (
	"Golang_task4/SERVICE"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostCtrl struct{}

func (p *PostCtrl) Add(c *gin.Context) {
	err := SERVICE.PostAdd(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (p *PostCtrl) Query(c *gin.Context) {
	err := SERVICE.PostQuery(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (p *PostCtrl) QueryList(c *gin.Context) {
	err := SERVICE.PostQueryList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (p *PostCtrl) Edit(c *gin.Context) {
	err := SERVICE.PostEdit(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (p *PostCtrl) Delete(c *gin.Context) {
	err := SERVICE.PostDelete(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
