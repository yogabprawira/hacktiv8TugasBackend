package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strconv"

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
	Session  Session    `json:"session"`
	Username string     `json:"username"`
	Name     string     `json:"name"`
	Resp     RespStatus `json:"resp"`
}

type Login struct {
	Username string `json:"username" form:"inputUsername"`
	Password string `json:"password" form:"inputPassword"`
}

type ReqLogin struct {
	Session Session `json:"session"`
	Login   Login   `json:"login"`
}

type PostId struct {
	Id string `uri:"id" binding:"required"`
}

type SubmitPost struct {
	Session Session `json:"session"`
	Post    Post    `json:"post"`
}

func checkSession(token string) bool {
	rows, err := stmtCheckSession.Query(token)
	if err != nil {
		return false
	}
	for rows.Next() {
		return true
	}
	return false
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
	if !checkSession(session.Token) {
		c.JSON(http.StatusNotFound, msgs)
		return
	}
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
	var reqLogin ReqLogin
	var respLogin RespLogin
	respLogin.Resp.Status = http.StatusNotFound
	respLogin.Resp.Message = ""
	err := c.ShouldBind(&reqLogin)
	if err != nil {
		c.JSON(http.StatusNotFound, respLogin)
		return
	}
	rows, err := stmtCheckLogin.Query(reqLogin.Login.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, respLogin)
		return
	}
	isMatch := false
	var name string
	for rows.Next() {
		var pwd string
		err = rows.Scan(&pwd, &name)
		if err != nil {
			continue
		}
		if reflect.DeepEqual(pwd, reqLogin.Login.Password) {
			isMatch = true
			break
		}
	}
	if !isMatch {
		c.JSON(http.StatusNotFound, respLogin)
		return
	}
	_, err = stmtLogin.Exec(reqLogin.Session.Token, reqLogin.Login.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, respLogin)
		return
	}
	respLogin.Resp.Status = http.StatusOK
	respLogin.Username = reqLogin.Login.Username
	respLogin.Name = name
	respLogin.Session.Token = reqLogin.Session.Token
	c.JSON(http.StatusNotFound, respLogin)
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
	if !checkSession(session.Token) {
		c.JSON(http.StatusNotFound, posts)
		return
	}
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
	if !checkSession(session.Token) {
		c.JSON(http.StatusNotFound, post)
		return
	}
	id, err := strconv.Atoi(postId.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, post)
		return
	}
	rows, err := stmtArticlesId.Query(id)
	if err != nil {
		c.JSON(http.StatusNotFound, post)
		return
	}
	for rows.Next() {
		err = rows.Scan(&post.Title, &post.Content, &post.Author, &post.TimePublished, &post.IsPublished)
		if err != nil {
			c.JSON(http.StatusNotFound, post)
			return
		}
	}
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
	if !checkSession(submitPost.Session.Token) {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
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
	if !checkSession(submitPost.Session.Token) {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	_, err = stmtArticlesAddId.Exec(submitPost.Post.Title, submitPost.Post.Content, submitPost.Post.Author,
		submitPost.Post.TimePublished, submitPost.Post.IsPublished, submitPost.Post.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	respStatus.Status = http.StatusOK
	c.JSON(http.StatusOK, respStatus)
}

func articlesIdPublishPost(c *gin.Context) {
	var respStatus RespStatus
	respStatus.Status = http.StatusNotFound
	respStatus.Message = ""
	var postId PostId
	err := c.ShouldBindUri(&postId)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	var session Session
	err = c.ShouldBind(&session)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	if !checkSession(session.Token) {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	id, err := strconv.Atoi(postId.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	_, err = stmtArticlesIdPublish.Exec(id)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	respStatus.Status = http.StatusOK
	c.JSON(http.StatusOK, respStatus)
}

func articlesIdUnpublishPost(c *gin.Context) {
	var respStatus RespStatus
	respStatus.Status = http.StatusNotFound
	respStatus.Message = ""
	var postId PostId
	err := c.ShouldBindUri(&postId)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	var session Session
	err = c.ShouldBind(&session)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	if !checkSession(session.Token) {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	id, err := strconv.Atoi(postId.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	_, err = stmtArticlesIdUnpublish.Exec(id)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	respStatus.Status = http.StatusOK
	c.JSON(http.StatusOK, respStatus)
}

func articlesIdDeletePost(c *gin.Context) {
	var respStatus RespStatus
	respStatus.Status = http.StatusNotFound
	respStatus.Message = ""
	var postId PostId
	err := c.ShouldBindUri(&postId)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	var session Session
	err = c.ShouldBind(&session)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	if !checkSession(session.Token) {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	id, err := strconv.Atoi(postId.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, respStatus)
		return
	}
	_, err = stmtArticlesIdDelete.Exec(id)
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
var stmtArticlesIdPublish *sql.Stmt
var stmtArticlesIdUnpublish *sql.Stmt
var stmtArticlesIdDelete *sql.Stmt
var stmtCheckSession *sql.Stmt
var stmtLogin *sql.Stmt
var stmtCheckLogin *sql.Stmt

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
	stmtArticlesIdPublish, err = db.Prepare("update `post` set is_published = 1 where post_id = ?;")
	if err != nil {
		log.Fatalln(err)
	}
	stmtArticlesIdUnpublish, err = db.Prepare("update `post` set is_published = 0 where post_id = ?;")
	if err != nil {
		log.Fatalln(err)
	}
	stmtArticlesIdDelete, err = db.Prepare("delete from `post` where post_id = ?;")
	if err != nil {
		log.Fatalln(err)
	}
	stmtCheckSession, err = db.Prepare("select session from `user` where session = ?;")
	if err != nil {
		log.Fatalln(err)
	}
	stmtLogin, err = db.Prepare("update `user` set session = ? where username = ?;")
	if err != nil {
		log.Fatalln(err)
	}
	stmtCheckLogin, err = db.Prepare("select password_hash, name from `user` where username = ?;")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/home", home)
	router.POST("/message", messagePost)
	router.POST("/contact", contactPost)
	router.POST("/login", login)
	router.POST("/articles", articlesPost)
	router.POST("/articles/id/:id", articlesIdPost)
	router.POST("/articles/id/:id/publish", articlesIdPublishPost)
	router.POST("/articles/id/:id/unpublish", articlesIdUnpublishPost)
	router.POST("/articles/id/:id/delete", articlesIdDeletePost)
	router.POST("/articles/add", articlesAddPost)
	router.POST("/articles/add/:id", articlesAddIdPost)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
