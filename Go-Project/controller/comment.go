// Controller 层负责处理来自路由的请求并调用 Service 层的方法进行处理

package controller

import (
	"github.com/life-studied/douyin-simple/response"
	"github.com/life-studied/douyin-simple/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentAction 处理评论操作的请求
func CommentAction(c *gin.Context) {
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentID, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	value, _ := c.Get("userid")
	userID, _ := value.(int64)

	var responseObj response.Comment_Action_Response
	var err error

	// 根据 actionType 调用不同的 Service 方法
	if actionType == "1" {
		responseObj, err = service.CreateComment(userID, videoID, commentText)
	} else if actionType == "2" {
		responseObj, err = service.DeleteComment(userID, videoID, commentID)
	}

	// 根据处理结果返回相应的 JSON 响应
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseObj)
	} else {
		c.JSON(http.StatusOK, responseObj)
	}
}

// CommentList 处理获取评论列表的请求
func CommentList(c *gin.Context) {
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userID, _ := c.Get("userid")
	comments, err := service.GetCommentList(videoID, userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Comment_List_Response{
			Response:    response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "获取评论失败"},
			CommentList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.Comment_List_Response{
		Response:    response.Response{StatusCode: 0, StatusMsg: "OK"},
		CommentList: comments,
	})
}
