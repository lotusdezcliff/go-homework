package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name      string    `gorm:"size:64;not null"`
    Email     string    `gorm:"size:128;uniqueIndex;not null"`
    PostCount int       `gorm:"default:0"`
    Posts     []Post
}

type Post struct {
    gorm.Model
    Title         string    `gorm:"size:128;not null"`
    Content       string    `gorm:"type:text"`
    UserID        uint      `gorm:"index"`
    CommentStatus string    `gorm:"size:16;default:'待評論'"`
    Comments      []Comment
}

type Comment struct {
    gorm.Model
    Content string `gorm:"type:text;not null"`
    Author  string `gorm:"size:64"`
    PostID  uint   `gorm:"index"`
}

func GetUserPostsWithComments(db *gorm.DB, userID uint) (User, error) {
	var user User
	err := db.Preload("Posts.Comments").First(&user, userID).Error
	return user, err
}

func GetTopCommentedPost(db *gorm.DB) (Post, error) {
	var topPost Post
	err := db.Model(&Post{}).
		Select("posts.*, count(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count desc").
		Preload("Comments").
		First(&topPost).Error
	return topPost, err
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "無評論").Error
	}

	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("資料庫連線失敗:", err)
		return
	}

	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		fmt.Println("自動建表失敗:", err)
		return
	}

	user := User{Name: "Aaron"}
	db.Create(&user)

	post := Post{Title: "Hello World", Content: "Hello World！", UserID: user.ID}
	db.Create(&post)

	comment := Comment{Content: "Nice！", PostID: post.ID}
	db.Create(&comment)

	fetchedUser, err := GetUserPostsWithComments(db, user.ID)
	if err != nil {
		fmt.Println("查詢失敗或找不到用戶:", err)
		return
	}

	fmt.Printf("成功查詢到用戶: %+v\n", fetchedUser)
}
