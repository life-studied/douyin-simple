package service

import "github.com/RaymondCode/simple-demo/controller"

type CommentService struct{}

func (cs *CommentService) SelectCommentList(videoId int64, token string) ([]controller.Comment, error) {
	var Comments []controller.Comment //定义查询到的评论
	var err error
	return Comments, nil
}
