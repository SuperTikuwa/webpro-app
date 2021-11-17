package dboperation

import "github.com/SuperTikuwa/webpro-app/model"

func CreateSchedule(schedule *model.Schedule) error {
	db := gormConnect()
	defer db.Close()

	return db.Create(*schedule).Error
}

func SelectAllScheduleByUserID(userID int64) ([]model.Schedule, error) {
	db := gormConnect()
	defer db.Close()

	var schedules []model.Schedule
	err := db.Where("user_id = ?", userID).Find(&schedules).Error
	return schedules, err
}
