package admin_service

import (
	"fmt"
	"time"

	admin_model "github.com/jacob/modules/modules/admin/model"
)

func AdminQueryValidator(query *admin_model.AdminQuery) (time.Time, time.Time, error) {

	// 檢查是否提供了時間參數
	if query.StartTime == "" || query.EndTime == "" {
		return time.Time{}, time.Time{}, fmt.Errorf("請提供起始時間和結束時間")
	}

	// 將字符串轉換為時間
	startTime, err := time.Parse("2006-01-02T15:04:05", query.StartTime)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("無效的起始時間")
	}
	endTime, err := time.Parse("2006-01-02T15:04:05", query.EndTime)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("無效的結束時間")
	}

	// 檢查結束時間是否早於起始時間
	if endTime.Before(startTime) {
		return time.Time{}, time.Time{}, fmt.Errorf("結束時間不能早於起始時間")
	}

	return startTime, endTime, nil
}
