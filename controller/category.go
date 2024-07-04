package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tani-hub-v2/database"
	"tani-hub-v2/repository"
	"tani-hub-v2/structs"
	"time"
)

func GetAllCategory(c *gin.Context) {
	var result gin.H

	categories, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetCategoryById(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.Id = int64(id)

	var result gin.H

	categories, err := repository.GetCategoryById(database.DbConnection, category)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Category",
	})
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.Id = int64(id)
	category.UpdatedAt = time.Now()

	err = repository.UpdateCategory(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Category",
	})

}

func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, err := strconv.Atoi(c.Param("id"))

	category.Id = int64(id)

	err = repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Category",
	})
}
