package controllers

import (
	"go-lang/initializers"
	"go-lang/models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "All posts",
		"data":    posts,
		"total":   len(posts),
	})
}

func PaginatePostsByOffset(c *gin.Context) {
	var posts []models.Post
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	initializers.DB.Offset(offset).Limit(limit).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "Offset paginated posts",
		"data":    posts,
		"total":   len(posts),
	})
}

func PaginatePostsByPage(c *gin.Context) {
	var posts []models.Post
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset := (page - 1) * limit
	initializers.DB.Offset(offset).Limit(limit).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "Page paginated posts",
		"data":    posts,
		"total":   len(posts),
	})
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post with id " + id + " not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post with id " + id,
		"data":    post,
	})
}

func CreatePost(c *gin.Context) {
	var post models.Post
	err := c.BindJSON(&post)
	if err != nil {
		panic(err)
	}
	initializers.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post with id " + id + " not found",
		})
		return
	}
	err := c.BindJSON(&post)
	if err != nil {
		panic(err)
	}
	initializers.DB.Model(&post).Updates(post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post with id " + id + " is updated",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post with id " + id + " not found",
		})
		return
	}
	initializers.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post with id " + id + " is deleted",
	})
}

func GetPostsConcurrently(c *gin.Context) {
	var posts []models.Post
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		initializers.DB.Find(&posts)
	}()

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"message": "All posts",
		"data":    posts,
		"total":   len(posts),
	})
}

func PaginatePostsByOffsetConcurrently(c *gin.Context) {
	var posts []models.Post
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		initializers.DB.Offset(offset).Limit(limit).Find(&posts)
	}()

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"message": "Offset paginated posts",
		"data":    posts,
		"total":   len(posts),
	})
}

func PaginatePostsByPageConcurrently(c *gin.Context) {
	var posts []models.Post
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset := (page - 1) * limit
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		initializers.DB.Offset(offset).Limit(limit).Find(&posts)
	}()

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"message": "Page paginated posts",
		"data":    posts,
		"total":   len(posts),
	})
}
