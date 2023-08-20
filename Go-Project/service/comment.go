package service

import (
	"github.com/life-studied/douyin-simple/model"
)

type CommentService struct{}

func (cs *CommentService) QueryComment(videoId int64, token string) ([]model.Comment, error) {
	//返回值定义
	var Comments []model.Comment
	var err error

	//dao层操作

	if err != nil {
		return nil, err
	}
	return Comments, nil
}
