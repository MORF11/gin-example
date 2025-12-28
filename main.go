package main
// perepisat bd na postgre dlya leapcell 
import (
	"fmt"

	"gin/tools"
	"github.com/gin-gonic/gin"
)

var usrs []tls.User
var fs []string

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	defer tls.DBClose()
	
    router.GET("/", func (c *gin.Context) {c.HTML(200,"index.html",gin.H{"title":"aboba","dt":usrs,"fs":fs})})
	
    router.POST("/add",func (c *gin.Context) {
		tls.Addu(c.PostForm("aname"))
		c.Redirect(301,"/")
		usrs = tls.Retu()
	})
    router.POST("/del",func (c *gin.Context) {
		tls.Delu(c.PostForm("id"))
		c.Redirect(301,"/")
		usrs = tls.Retu()
	})
	
	router.POST("/fileup",func (c *gin.Context) {
		file,e := c.FormFile("file")
		if e != nil  || file == nil {
			fmt.Println("pizdaruliam") 
			return
		}
		fmt.Println(file,file.Filename)
		c.SaveUploadedFile(file,"files/"+file.Filename)
		c.Redirect(301,"/")
		fs = tls.Check_dir()
	})
	router.GET("/filedw/:name",func (c *gin.Context) {
		c.File("files/"+c.Param("name"))
		c.Redirect(301,"/")
	})
	
    router.GET("/get", func (c *gin.Context) {c.JSON(200,usrs)})
    router.POST("/post", func (c *gin.Context) {
		var n string
		c.BindJSON(&n)
		tls.Addu(n)
		c.JSON(201,usrs)
	})
	
    router.Run("localhost:8080")
}

func init() {
	fs = tls.Check_dir()
	usrs = tls.Retu()
	fmt.Println(fs,"\n",usrs)
}