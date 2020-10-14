package model

import (
	"j2pay-server/model/response"

	"time"
)

type SystemMessage struct {
	Id         int
	Title      string      `gorm:"default:'';comment:'标题';"`
	BeginTime  time.Time   `gorm:"comment:'开始时间';type:timestamp;";json:"beginTime"`
	EndTime    time.Time   `gorm:"comment:'结束时间';type:timestamp;";json:"endTime"`
	Status     int         `gorm:"default:1;comment:'是否作废 0：否，1：是';"`
	AdminUsers []AdminUser `gorm:"many2many:system_message_user;"`
}

// 获取所有系统消息
func (s *SystemMessage) GetAll(page, pageSize int, where ...interface{}) (response.SystemMessagePage, error) {
	all := response.SystemMessagePage{
		Total:       s.GetCount(where...),
		PerPage:     pageSize,
		CurrentPage: page,
		Data:        []response.SystemMessageList{},
	}
	offset := GetOffset(page, pageSize)
	err := Getdb().Table("system_message").
		Limit(pageSize).
		Offset(offset).
		Find(&all.Data, where...).Error
	if err != nil {
		return response.SystemMessagePage{}, err
	}
	return all, err
}

// 根据ID获取公告详情
func (s *SystemMessage) Detail(id ...int) (res response.SystemMessageList, err error) {
	searchId := s.Id
	if len(id) > 0 {
		searchId = id[0]
	}
	err = Getdb().Table("system_message").
		Where("id = ?", searchId).
		First(&res).
		Error
	return
}

//新增公告
func (s *SystemMessage) Create(user []int) error {
	Getdb().Save(&s)
	for _,v  := range user {
			Getdb().Save(&SystemMessageUser{
				SystemMessageId: s.Id,
				AdminUserId: v,
			})
	}
	return nil
}

// 编辑系统公告
func (s *SystemMessage) Edit(users []int) error {
	tx := Getdb().Begin()
	updateInfo := map[string]interface{}{
		"title":      s.Title,
		"begin_time": s.BeginTime,
		"end_time":   s.EndTime,
	}
	if s.Title != "" {
		updateInfo["title"] = s.Title
	}

	if err := tx.Model(&SystemMessage{Id: s.Id}).
		Updates(updateInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(SystemMessageUser{}, "system_message_id = ?", s.Id).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, v := range users {
		err := tx.Create(&SystemMessageUser{
			SystemMessageId: s.Id,
			AdminUserId: v,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// 删除系统公告
func (s *SystemMessage) Del() error {
	tx := Getdb().Begin()
	if err := tx.Delete(s, "id = ?", s.Id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(SystemMessageUser{}, "system_message_id = ?", s.Id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 获取所有系统公告数量
func (s *SystemMessage) GetCount(where ...interface{}) (count int) {
	if len(where) == 0 {
		Getdb().Model(&s).Count(&count)
		return
	}
	Getdb().Model(&s).Where(where[0], where[1:]...).Count(&count)
	return
}

func GetAllMessage() (mapping map[int]response.AdminSystemMessage) {
	var systemMessages []response.AdminSystemMessage

	mapping = make(map[int]response.AdminSystemMessage)
	Getdb().Table("system_message").Select("id,title,begin_time,end_time").Order("id desc").Find(&systemMessages)
	for _, systemMessage := range systemMessages {
		mapping[systemMessage.Id] = systemMessage
	}
	return
}
