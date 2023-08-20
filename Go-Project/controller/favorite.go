package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/service"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
    //接收参数，并判断是否合法
    token, tokenOk := c.GetQuery("token")
    if tokenOk {
        c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Lack of token"})
        return
    }
    videoId, videoIdOk := c.GetQuery("video_id")
    if videoIdOk {
        c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Lack of video_id"})
        return
    }
    actionType, actionTypeOk := c.GetQuery("action_type")
    if actionTypeOk {
        c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Lack of action_type"})
        return
    }
    //判断操作类型
    if actionType == "1" {
        service.FavoriteVideo(videoId, token)
    } else if actionType == "2" {
        service.UnfavoriteVideo(videoId, token)
    } else {
        c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Invalid action_type"})
    }

    if _, exist := usersLoginInfo[token]; exist {
        c.JSON(http.StatusOK, Response{StatusCode: 0})
    } else {
        c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
    }
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
    //接收参数，并判断是否合法
    token, tokenOk := c.GetQuery("token")
    if tokenOk {
        c.JSON(http.StatusOK, VideoListResponse{
            Response: Response{
                StatusCode: 1,
                StatusMsg:  "Lack of token",
            },
            VideoList: nil,
        })
        return
    }
    userId, userIdOk := c.GetQuery("user_id")
    if userIdOk {
        c.JSON(http.StatusOK, VideoListResponse{
            Response: Response{
                StatusCode: 1,
                StatusMsg:  "Lack of user_id",
            },
            VideoList: nil,
        })
        return
    }
    videoList := service.ReadFavoriteVideo(token, userId)
    c.JSON(http.StatusOK, VideoListResponse{
        Response: Response{
            StatusCode: 0,
        },
        VideoList: videoList,
    })
}
