package models

import (
	"gin/dao"

	"gorm.io/gorm"
)

// Todo 待办事项
type Todo struct {
	gorm.Model
	Title  string `gorm:"title" json:"title"`
	Status bool   `gorm:"status" json:"status"`
}

// GetByID 通过ID查找
func GetByID(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// Create 新增Todo项
func Create(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

// Update 更新数据
func Update(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

// GetAllList 获取所有数据
func GetAllList() (list []*Todo, err error) {

	if err = dao.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return
}

// GetListByPage 分页获取数据
func GetListByPage(limit int, pageSize int) (list []*Todo, err error) {
	err = dao.DB.Offset((pageSize - 1) * pageSize).Limit(limit).Find(&list).Error
	return
}

// Delete 删除
func Delete(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
