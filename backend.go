package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id            int    `json:"id"`
	Title         string `json:"title" form:"postTitle"`
	Content       string `json:"content" form:"postContent"`
	Author        string `json:"author"`
	TimePublished string `json:"time_published"`
	IsPublished   int    `json:"is_published"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

type Message struct {
	FullName string `form:"fullName" json:"full_name"`
	Email    string `form:"email" json:"email"`
	Message  string `form:"message" json:"message"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}

type RespStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Session struct {
	Token string `json:"token"`
}

type RespLogin struct {
	Session string     `json:"session"`
	Resp    RespStatus `json:"resp"`
}

type Login struct {
	Email    string `json:"email" form:"inputEmail"`
	Password string `json:"password" form:"inputPassword"`
}

type PostId struct {
	Id string `uri:"id" binding:"required"`
}

type SubmitPost struct {
	Session Session `json:"session"`
	Post    Post    `json:"post"`
}

func home(c *gin.Context) {
	var posts Posts
	posts.Posts = make([]Post, 0)
	rows, err := stmtHome.Query()
	if err != nil {
		c.JSON(http.StatusNotFound, posts)
		return
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.TimePublished, &post.IsPublished)
		if err != nil {
			c.JSON(http.StatusNotFound, posts)
			return
		}
		posts.Posts = append(posts.Posts, post)
	}
	c.JSON(http.StatusOK, posts)
}

func messagePost(c *gin.Context) {
	var msg Message
	var respStatus RespStatus
	respStatus.Status = http.StatusNotFound
	respStatus.Message = ""
	err := c.ShouldBind(&msg)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	_, err = stmtMessage.Exec(msg.FullName, msg.Email, msg.Message)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	respStatus.Status = http.StatusOK
	c.JSON(http.StatusOK, respStatus)
}

func contactPost(c *gin.Context) {
	var session Session
	var msgs Messages
	msgs.Messages = make([]Message, 0)
	err := c.ShouldBind(&session)
	if err != nil {
		c.JSON(http.StatusNotFound, msgs)
		return
	}
	// session validation
	rows, err := stmtContact.Query()
	if err != nil {
		c.JSON(http.StatusNotFound, msgs)
		return
	}
	for rows.Next() {
		var msg Message
		err = rows.Scan(&msg.FullName, &msg.Email, &msg.Message)
		if err != nil {
			c.JSON(http.StatusNotFound, msgs)
			return
		}
		msgs.Messages = append(msgs.Messages, msg)
	}
	c.JSON(http.StatusOK, msgs)
}

func login(c *gin.Context) {

}

func logout(c *gin.Context) {
	// remove session
}

func articlesPost(c *gin.Context) {
	var posts Posts
	posts.Posts = make([]Post, 0)
	var session Session
	err := c.ShouldBind(&session)
	if err != nil {
		c.JSON(http.StatusNotFound, posts)
		return
	}
	// session validation
	rows, err := stmtArticles.Query()
	if err != nil {
		c.JSON(http.StatusNotFound, posts)
		return
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.TimePublished, &post.IsPublished)
		if err != nil {
			c.JSON(http.StatusNotFound, posts)
			return
		}
		posts.Posts = append(posts.Posts, post)
	}
	c.JSON(http.StatusOK, posts)
}

func articlesIdPost(c *gin.Context) {
	var post Post
	var postId PostId
	err := c.ShouldBindUri(&postId)
	if err != nil {
		c.JSON(http.StatusNotFound, post)
		return
	}
	var session Session
	err = c.ShouldBind(&session)
	if err != nil {
		c.JSON(http.StatusNotFound, post)
		return
	}
	// session validation
	rows, err := stmtArticlesId.Query(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, post)
		return
	}
	for rows.Next() {
		err = rows.Scan(&post.Title, &post.Content, &post.Author, &post.TimePublished, &post.IsPublished)
		log.Println(post)
		if err != nil {
			c.JSON(http.StatusNotFound, post)
			return
		}
	}
	log.Println(post)
	c.JSON(http.StatusOK, post)
}

func articlesAddPost(c *gin.Context) {
	var respStatus RespStatus
	respStatus.Status = http.StatusNotFound
	respStatus.Message = ""
	var submitPost SubmitPost
	err := c.ShouldBind(&submitPost)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	// session validation
	_, err = stmtArticlesAdd.Exec(submitPost.Post.Title, submitPost.Post.Content, submitPost.Post.Author,
		submitPost.Post.TimePublished, submitPost.Post.IsPublished)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	respStatus.Status = http.StatusOK
	c.JSON(http.StatusOK, respStatus)
}

func articlesAddIdPost(c *gin.Context) {
	var respStatus RespStatus
	respStatus.Status = http.StatusNotFound
	respStatus.Message = ""
	var postId PostId
	err := c.ShouldBindUri(&postId)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	var submitPost SubmitPost
	err = c.ShouldBind(&submitPost)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	// session validation
	_, err = stmtArticlesAddId.Exec(submitPost.Post.Title, submitPost.Post.Content, submitPost.Post.Author,
		submitPost.Post.TimePublished, submitPost.Post.IsPublished, submitPost.Post.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	respStatus.Status = http.StatusOK
	c.JSON(http.StatusOK, respStatus)
}

var db *sql.DB
var stmtHome *sql.Stmt
var stmtArticles *sql.Stmt
var stmtMessage *sql.Stmt
var stmtContact *sql.Stmt
var stmtArticlesId *sql.Stmt
var stmtArticlesAdd *sql.Stmt
var stmtArticlesAddId *sql.Stmt

const UserDb = "root"
const PwdDb = "yoga123"
const DbName = "yogablog"

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8", UserDb, PwdDb, DbName)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	stmtHome, err = db.Prepare("select post_id, post_title, post_content, post_author, time_published, is_published from post where is_published = 1;")
	if err != nil {
		log.Fatalln(err)
	}
	stmtArticles, err = db.Prepare("select post_id, post_title, post_content, post_author, time_published, is_published from post")
	if err != nil {
		log.Fatalln(err)
	}
	stmtMessage, err = db.Prepare("insert into `message` values (null, ?, ?, ?);")
	if err != nil {
		log.Fatalln(err)
	}
	stmtContact, err = db.Prepare("select full_name, email, message from `message`")
	if err != nil {
		log.Fatalln(err)
	}
	stmtArticlesId, err = db.Prepare("select post_title, post_content, post_author, time_published, is_published from post where post_id = ?;")
	if err != nil {
		log.Fatalln(err)
	}
	stmtArticlesAdd, err = db.Prepare("insert into `post` values (null, ?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatalln(err)
	}
	stmtArticlesAddId, err = db.Prepare("update `post` set post_title = ?, post_content = ?, post_author = ?, time_published = ?, is_published = ? where post_id = ?;")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/home", home)
	router.POST("/message", messagePost)
	router.POST("/contact", contactPost)
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/articles", articlesPost)
	router.POST("/articles/id/:id", articlesIdPost)
	router.POST("/articles/add", articlesAddPost)
	router.POST("/articles/add/:id", articlesAddIdPost)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
