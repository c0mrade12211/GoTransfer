package servers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HTTPserver(port int) {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save file",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
		})
	})

	r.Run(":" + strconv.Itoa(port))
}
