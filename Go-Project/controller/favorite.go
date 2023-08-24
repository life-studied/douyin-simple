package controller

import (
	"fmt"
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
	userFromToken, exist := usersLoginInfo[token]
	if exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Token is invalid"})
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
	var video = service.FavoriteVideoID{
		VideoID: videoId,
	}
	var user = service.FavoriteUserID{
		UserID:   userFromToken.Id,
		UserName: userFromToken.Name,
	}
	if actionType == "1" {
		err := service.FavoriteVideo(video, user)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Favorite failed"})
			//打印报错
			fmt.Println(err.Error())
			return
		}

	} else if actionType == "2" {
		err := service.UnfavoriteVideo(video, user)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Unfavorite failed"})
			//打印报错
			fmt.Println(err.Error())
			return
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Invalid action_type"})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Success"})
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
	//检测token
	userFromToken, exist := usersLoginInfo[token]
	if exist {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Token is invalid",
			},
			VideoList: nil,
		})
		return
	}
	//转为service层使用的结构体类型
	user := service.FavoriteUserID{
		UserID:   userFromToken.Id,
		UserName: userFromToken.Name,
	}
	//用不到但是必须要有的接收参数
	_, userIdOk := c.GetQuery("user_id")
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
	//获取用户收藏的视频列表
	videoList, err := service.ReadFavoriteVideo(user)
	if err != nil {

	}
	//转为前端需要的结构体类型
	c.JSON(http.StatusOK, serviceToVideoList(videoList))
}

func serviceToVideoList(favoriteVideoList []service.FavoriteVideoList) (response VideoListResponse) {

	response.Response = Response{
		StatusCode: 0,
		StatusMsg:  "Success",
	}
	response.VideoList = make([]Video, 0, len(favoriteVideoList))
	for _, favoriteVideo := range favoriteVideoList {
		response.VideoList = append(response.VideoList, Video{
			Id:            favoriteVideo.Id,
			Author:        serviceToUser(favoriteVideo.Author),
			PlayUrl:       favoriteVideo.PlayUrl,
			CoverUrl:      favoriteVideo.CoverUrl,
			FavoriteCount: favoriteVideo.FavoriteCount,
			CommentCount:  favoriteVideo.CommentCount,
			IsFavorite:    true,
		})
	}
	return response
}

func serviceToUser(favoriteUser service.FavoriteUserID) (user User) {
	user = User{
		Id:            favoriteUser.UserID,
		Name:          favoriteUser.UserName,
		FollowCount:   favoriteUser.FollowCount,
		FollowerCount: favoriteUser.FollowerCount,
	}
	return user
}
