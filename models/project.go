package models

type ProjectData struct {
  ID string
  Sponsor string
  Name string
}

type ProjectDesc struct {
  ID uint `gorm:"primaryKey"`
  Timestamp int64 `json:"timestamp"`
  ProjectID string `json:"project_id"`
  ProjectName string `json:"project_name"`
  ProjectSponsor string `json:"project_sponsor"`
}

type QueryParams struct {
  Sponsor string
  StartEpoch *int64
  EndEpoch *int64
}
