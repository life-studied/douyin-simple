package dao

// GetCommentByIdListById 根据video_id返回视频评论列表
func GetCommentByIdListById(videoID int64) ([]Comment, error) {
	var comments []Comment
	err := SERVER_DB.Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetUserById 根据user_id返回用户结构体
func GetUserById(userID int64) (User, error) {
	var user User
	err := SERVER_DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// GetCommentById 通过commentID 返回comment结构体
func GetCommentById(commentID int64) (Comment, error) {
	var comment Comment
	err := SERVER_DB.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return Comment{}, err
	}
	return comment, nil
}

// CreateComment 创建评论
func CreateComment(comment *Comment) error {
	err := SERVER_DB.Create(comment).Error
	return err
}

// DeleteCommentById 根据id删除评论
func DeleteCommentById(commentID int64) error {
	err := SERVER_DB.Where("id = ?", commentID).Delete(Comment{}).Error
	return err
}
