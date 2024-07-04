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

func GetAllArticle(c *gin.Context) {
	var result gin.H

	articles, err := repository.GetAllArticle(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": articles,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetArticleById(c *gin.Context) {
	var article structs.Article
	id, _ := strconv.Atoi(c.Param("id"))

	article.Id = int64(id)

	var result gin.H

	articles, err := repository.GetArticleById(database.DbConnection, article)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": articles,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertArticle(c *gin.Context) {
	var article structs.Article

	err := c.ShouldBindJSON(&article)
	if err != nil {
		panic(err)
	}

	err = repository.InsertArticle(database.DbConnection, article)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Article",
	})
}

func UpdateArticle(c *gin.Context) {
	var article structs.Article
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&article)
	if err != nil {
		panic(err)
	}

	article.Id = int64(id)
	article.UpdatedAt = time.Now()

	err = repository.UpdateArticle(database.DbConnection, article)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Article",
	})

}

func DeleteArticle(c *gin.Context) {
	var article structs.Article
	id, err := strconv.Atoi(c.Param("id"))

	article.Id = int64(id)

	err = repository.DeleteArticle(database.DbConnection, article)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Article",
	})
}
