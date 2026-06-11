package controllers

import (
	"github.com/testdev-learning/DataHub/backend/models"
	"github.com/testdev-learning/DataHub/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestCaseCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Protocol    string `json:"protocol"`
	Method      string `json:"method"`
	URL         string `json:"url" binding:"required"`
	Headers     string `json:"headers"`
	Body        string `json:"body"`
	Params      string `json:"params"`
	Expected    string `json:"expected"`
}

func CreateTestCase(c *gin.Context) {
	var req TestCaseCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := c.Get("user_id")

	testCase := models.TestCase{
		Name:        req.Name,
		Description: req.Description,
		Protocol:    req.Protocol,
		Method:      req.Method,
		URL:         req.URL,
		Headers:     req.Headers,
		Body:        req.Body,
		Params:      req.Params,
		Expected:    req.Expected,
		Status:      "draft",
		CreatorID:   userId.(uint),
	}

	if err := utils.DB.Create(&testCase).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "创建成功", "data": testCase})
}

func GetTestCases(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var testCases []models.TestCase
	if err := utils.DB.Where("creator_id = ?", userId).Find(&testCases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": testCases})
}

func GetTestCase(c *gin.Context) {
	id := c.Param("id")
	var testCase models.TestCase
	if err := utils.DB.First(&testCase, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用例不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": testCase})
}

func UpdateTestCase(c *gin.Context) {
	id := c.Param("id")
	var testCase models.TestCase
	if err := utils.DB.First(&testCase, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用例不存在"})
		return
	}

	var req TestCaseCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	testCase.Name = req.Name
	testCase.Description = req.Description
	testCase.Protocol = req.Protocol
	testCase.Method = req.Method
	testCase.URL = req.URL
	testCase.Headers = req.Headers
	testCase.Body = req.Body
	testCase.Params = req.Params
	testCase.Expected = req.Expected

	if err := utils.DB.Save(&testCase).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "data": testCase})
}

func DeleteTestCase(c *gin.Context) {
	id := c.Param("id")
	if err := utils.DB.Delete(&models.TestCase{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}