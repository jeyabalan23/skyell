
// backend/models/analysis.go
package models

import "time"

type Analysis struct {
    ID            uint           `gorm:"primaryKey"`
    URL           string
    HTMLVersion   string
    Title         string
    Headings      map[string]int `gorm:"type:json"`
    InternalLinks int
    ExternalLinks int
    BrokenLinks   []string       `gorm:"type:json"`
    HasLoginForm  bool
    Status        string
    CreatedAt     time.Time
}
