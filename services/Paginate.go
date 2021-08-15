package services

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginator(c *gin.Context) func(db *gorm.DB) *gorm.DB{
	return func(db *gorm.DB) *gorm.DB{
		page,_ := strconv.Atoi(c.Query("page"))
		
		offset := (page - 1) * page

		return db.Offset(offset).Limit(10)
	}
}