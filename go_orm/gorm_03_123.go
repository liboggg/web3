package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "test:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移表结构
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 查询某个用户发布的所有文章及其对应的评论信息
	//queryUserPostsAndComments(db, 1)

	// 查询评论数量最多的文章信息
	//queryMostCommentedPost(db)

	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	//createPost(db)

	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
	//如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	deleteComment(db)

}
func deleteComment(db *gorm.DB) {
	commentId := 4
	var comment Comment
	if err := db.Where("id = ?", commentId).First(&comment).Error; err != nil {
		fmt.Printf("Error finding comment: %v\n", err)
		return
	}
	db.Where("id = ?", commentId).Delete(&comment)
}

func (comment *Comment) AfterDelete(db *gorm.DB) (err error) {
	fmt.Println(*comment)
	fmt.Println(comment)
	if comment.PostId > 0 {
		var count int64
		db.Model(&Comment{}).Where(" postId = ?", comment.PostId).Count(&count)

		var statusStr string
		if count == 0 {
			statusStr = "无评论"
		} else {
			statusStr = fmt.Sprintf("有%d条评论", count)
		}
		db.Model(&Post{}).Where("id = ?", comment.PostId).Update("statusStr", statusStr)
	}
	return
}

func createPost(db *gorm.DB) {
	userId := 1
	db.Create(&Post{
		Title:     "测试自动创建",
		UserId:    userId,
		StatusStr: "",
	})
}

func (u *Post) AfterCreate(db *gorm.DB) (err error) {
	if u.Id > 0 && u.UserId > 0 {
		//查询用户文章数量
		var count int64
		db.Model(&Post{}).Where("userId = ?", u.UserId).Count(&count)
		db.Model(&User{}).Where("id = ?", u.UserId).Update("postCount", count)
	}
	return
}

// 查询某个用户发布的所有文章及其对应的评论信息
func queryUserPostsAndComments(db *gorm.DB, userId int) {
	var user User
	db.Preload("Posts.Comments").First(&user, userId)
	fmt.Printf("User: %+v\n", user)
	for _, post := range user.Posts {
		fmt.Printf("  Post: %+v\n", post)
		for _, comment := range post.Comments {
			fmt.Printf("    Comment: %+v\n", comment)
		}
	}
	fmt.Println("--------------------------------")
	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling user to JSON: %v\n", err)
		return
	}
	fmt.Println(string(userJSON))

}

// 查询评论数量最多的文章信息
func queryMostCommentedPost(db *gorm.DB) {
	//var post Post
	//db.Preload("Comments").Joins("JOIN `comment` ON post.id = comment.postId").
	//	Group("comment.postId").
	//	Order("COUNT(comment.id) DESC").
	//	First(&post)
	//fmt.Printf("Most Commented Post: %+v\n", post)
	//for _, comment := range post.Comments {
	//	fmt.Printf("  Comment: %+v\n", comment)
	//}

	var post Post
	db.Joins("JOIN `comment` ON post.id = comment.postId").
		Group("comment.postId").
		Order("COUNT(comment.id) DESC").
		First(&post)
	fmt.Printf("Most Commented Post: %+v\n", post)
}

// 用户
type User struct {
	Id        int    `db:"id" gorm:"primaryKey`
	Name      string `db:"name"`
	PostCount int    `db:"postCount"`
	Posts     []Post `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "user"
}

// 文章
type Post struct {
	Id        int       `db:"id" gorm:"primaryKey`
	Title     string    `db:"title"`
	UserId    int       `db:"userId" gorm:"column:userId"`
	StatusStr string    `db:"statusStr" gorm:"column:statusStr"`
	Comments  []Comment `gorm:"foreignKey:PostId"`
}

func (Post) TableName() string {
	return "post"
}

// 评论
type Comment struct {
	Id      int    `db:"id" gorm:"primaryKey`
	PostId  int    `db:"postId" gorm:"column:postId"`
	Content string `db:"content"`
}

func (Comment) TableName() string {
	return "comment"
}
