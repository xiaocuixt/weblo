package controllers
import (
  "net/http"
	"github.com/gin-gonic/gin"
)

func Abount(c *gin.Context)  {
	c.HTML(http.StatusOK, "home/about.tmpl", gin.H{
  })
}
