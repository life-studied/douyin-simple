package service

import (
	"github.com/life-studied/douyin-simple/controller"
)

func FavoriteVideo(videoId string, token string) {

}

func UnfavoriteVideo(videoId string, token string) {

}

func ReadFavoriteVideo(token string, userId string) (videoList []controller.Video) {
	videoList = controller.DemoVideos
	return videoList
}

func WriteFavoriteVideo(token string, userId string) {

}

func DeleteFavoriteVideo(token string, userId string) {

}
