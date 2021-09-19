package migration

import (
	"go-assesment/helper"
	"go-assesment/model"
)

func Migrate() {
	helper.DB().AutoMigrate(&model.Question{})
	helper.DB().AutoMigrate(&model.Answer{})
}
