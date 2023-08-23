package dao

import (
	"github.com/life-studied/douyin-simple/global"
	"github.com/life-studied/douyin-simple/model"
)

// GetCommentByIdListById 根据video_id返回视频评论列表
func GetCommentByIdListById(videoID int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := global.DB.Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetUserById 根据user_id返回用户结构体
func GetUserById(userID int64) (User, error) {
	var user User
	err := global.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// GetCommentById 通过commentID 返回comment结构体
func GetCommentById(commentID int64) (model.Comment, error) {
	var comment model.Comment
	err := global.DB.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

// CreateComment 创建评论
func CreateComment(comment *model.Comment) error {
	err := global.DB.Create(comment).Error
	return err
}

// DeleteCommentById 根据id删除评论
func DeleteCommentById(commentID int64) error {
	err := global.DB.Where("id = ?", commentID).Delete(model.Comment{}).Error
	return err
}
