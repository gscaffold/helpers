package main

import (
	"context"

	"github.com/gscaffold/helpers/databases/mysql"
	"github.com/gscaffold/helpers/logger"
	"github.com/gscaffold/utils"
	"gorm.io/gorm"
)

func main() {
	db := mysql.MustOpen("test")

	// insert
	{
		user := &User{
			UserID: "123",
			Name:   "张三",
		}
		err := db.Create(user).Error
		if err != nil {
			panic(err)
		}
	}

	// update
	{
		user := &User{
			UserID: "123",
			Name:   "李四",
		}
		err := db.Model(user).Where("user_id=?", user.UserID).Updates(user).Error
		if err != nil {
			panic(err)
		}
	}

	// query
	{
		user := &User{}
		err := db.Model(user).Where("user_id='123'").First(user).Error
		if err != nil {
			panic(err)
		}
		logger.Infof(context.TODO(), "query user:%s", utils.ToJSONStr(user))
	}
}

type User struct {
	*gorm.Model
	UserID string `gorm:"column:user_id" json:"user_id"`
	Name   string `gorm:"column:name" json:"name"`
}

func (u *User) TableName() string {
	return "users"
}

/*
create table sql

CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户唯一标识',
  `name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户姓名',
  `created_at` DATETIME(3) NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` DATETIME(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=117202 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';
*/
