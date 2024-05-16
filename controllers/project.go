package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/AlissonBarbosa/shylockgo-api/models"
	"github.com/gin-gonic/gin"
)

func getLatestEpoch() (int64, error) {
  var maxEpoch int64
  if err := models.DB.Table("project_descs").Select("MAX(epoch)").Row().Scan(&maxEpoch); err != nil {
    return 0, err
  }
  return maxEpoch, nil
}

func parseQueryParams(c *gin.Context) models.QueryParams {
  params := models.QueryParams{
    Sponsor: c.Query("sponsor"),
  }

  if startStr := c.Query("start"); startStr != "" {
    if start, err := strconv.ParseInt(startStr, 10, 64); err == nil {
      params.StartEpoch = &start
    }
  }

  if endStr := c.Query("end"); endStr != "" {
    if end, err := strconv.ParseInt(endStr, 10, 64); err == nil {
      params.EndEpoch = &end
    }
  }
  return params
}

func getProjectsDesc(params models.QueryParams) ([]models.ProjectDesc, error) {
  var projectsDesc []models.ProjectDesc
  query := models.DB

  if params.Sponsor != "" {
    query = query.Where("project_sponsor LIKE ?", "%"+params.Sponsor+"%")
  }

  if params.StartEpoch != nil {
    query = query.Where("timestamp >= ?", *params.StartEpoch)
  }

  if params.EndEpoch != nil {
    query = query.Where("timestamp <= ?", *params.EndEpoch)
  }

  if err := query.Find(&projectsDesc).Error; err != nil {
    slog.Error("Error quering database")
    return nil, err
  }
  
  return projectsDesc, nil
}

func GetProjects(c *gin.Context) {
  params := parseQueryParams(c)
  projectsDesc, err := getProjectsDesc(params)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "status": "error",
      "message": "Something wrong on server side",
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "data": projectsDesc,
  })
}
