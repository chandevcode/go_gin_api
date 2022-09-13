package books

import (
	"net/http"

	"github.com/chanlevelup/go_gin_api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)
	c.Status(http.StatusOK)
}
