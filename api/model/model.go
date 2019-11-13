package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
  gorm.model
  Name string `gorm:"unique" json:"name"`
  City string `json:"city"`
  Age int `json:"age"`
  Status bool `json:"status"`
}
