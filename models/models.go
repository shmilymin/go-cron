package models

import (
	"time"
)

type Model struct {
	Id         uint64    `form:"id"`
	CreateTime time.Time `form:"createTime"`
	UpdateTime time.Time `form:"updateTime"`
}
