// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type RoleRepository struct {
}

//根据ID查找
func (repo *RoleRepository) FindRoleByID(DB *gorm.DB, id int) (role models.RoleDetail, err error) {
	err = DB.Where("id = ?", id).First(&role).Error
	return
}

// 创建角色
func (repo *RoleRepository) InsertRole(DB *gorm.DB, model models.RoleCreate) (result models.RoleDetail, err error) {
	model.ID = 0
	err = DB.Create(&model).Error
	if err == nil {
		return repo.FindRoleByID(DB, model.ID)
	}
	return
}

// 修改角色
func (repo *RoleRepository) UpdateRole(DB *gorm.DB, role models.RoleUpdate) (result models.RoleDetail, err error) {
	err = DB.Updates(role).Error
	return
}

// 删除角色
func (repo *RoleRepository) DeleteRole(DB *gorm.DB, ids []int) (err error) {
	return DB.Where("id IN (?)", ids).Delete(models.RoleDetail{}).Error
}

// 查询角色
func (repo *RoleRepository) FindAllRole(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (results []*models.RoleList, total int, err error) {
	if len(order) == 0 {
		order = "id desc"
	}
	db := DB.Model(&models.RoleList{}).Order(order)
	if len(query) > 0 {
		db = db.Where(query, queryArgs[:])
	}
	db.Count(&total)
	err = db.Offset((page - 1) * size).Limit(size).Find(&results).Error
	return
}
