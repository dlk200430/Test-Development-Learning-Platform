package controllers

import (
	"github.com/testdev-learning/DataHub/backend/models"
	"github.com/testdev-learning/DataHub/backend/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ExecuteTestCase(c *gin.Context) {
	id := c.Param("id")
	var testCase models.TestCase
	if err := utils.DB.First(&testCase, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用例不存在"})
		return
	}

	startTime := time.Now()
	report := models.TestReport{
		TestCaseID: testCase.ID,
		Status:     "running",
		StartTime:  startTime,
	}
	utils.DB.Create(&report)

	var client http.Client
	var req *http.Request
	var err error

	finalURL := buildURL(testCase.URL, testCase.Params)

	if testCase.Method == "POST" || testCase.Method == "PUT" || testCase.Method == "PATCH" {
		req, err = http.NewRequest(testCase.Method, finalURL, bytes.NewBuffer([]byte(testCase.Body)))
	} else {
		req, err = http.NewRequest(testCase.Method, finalURL, nil)
	}

	if err != nil {
		report.Status = "failed"
		report.Error = err.Error()
		report.EndTime = time.Now()
		report.Duration = int64(time.Since(startTime).Milliseconds())
		utils.DB.Save(&report)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建请求失败"})
		return
	}

	setHeaders(req, testCase.Headers)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		report.Status = "failed"
		report.Error = err.Error()
		report.EndTime = time.Now()
		report.Duration = int64(time.Since(startTime).Milliseconds())
		utils.DB.Save(&report)
		c.JSON(http.StatusOK, gin.H{"message": "执行完成", "report": report})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	report.Response = string(body)
	report.EndTime = time.Now()
	report.Duration = int64(time.Since(startTime).Milliseconds())

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		report.Status = "passed"
	} else {
		report.Status = "failed"
	}

	utils.DB.Save(&report)
	c.JSON(http.StatusOK, gin.H{"message": "执行完成", "report": report})
}

func buildURL(baseURL, params string) string {
	if params == "" {
		return baseURL
	}

	var paramMap map[string]string
	if err := json.Unmarshal([]byte(params), &paramMap); err != nil {
		return baseURL
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return baseURL
	}

	query := parsedURL.Query()
	for k, v := range paramMap {
		query.Add(k, v)
	}
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String()
}

func setHeaders(req *http.Request, headers string) {
	if headers == "" {
		return
	}

	var headerMap map[string]string
	if err := json.Unmarshal([]byte(headers), &headerMap); err != nil {
		return
	}

	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
}

func GetTestReports(c *gin.Context) {
	testCaseID := c.Query("test_case_id")
	var reports []models.TestReport
	query := utils.DB

	if testCaseID != "" {
		id, _ := strconv.Atoi(testCaseID)
		query = query.Where("test_case_id = ?", id)
	}

	if err := query.Order("created_at DESC").Find(&reports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reports})
}

func GetTestReport(c *gin.Context) {
	id := c.Param("id")
	var report models.TestReport
	if err := utils.DB.Preload("TestCase").First(&report, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "报告不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": report})
}