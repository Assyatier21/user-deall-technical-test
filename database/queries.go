package database

const (
	InsertUser            = "INSERT INTO users VALUES('%d','%s','%s','%d','%s')"
	GetUserByUsernamePass = "SELECT * FROM users WHERE username='%s' AND password='%s' AND token='%s' AND role_id='%d'"
	GetUserByUsername     = "SELECT * FROM users WHERE username='%s'"

	GetArticleById   = "SELECT * FROM articles WHERE id='%d'"
	InsertArticle    = "INSERT INTO articles VALUES('%d','%s','%s','%s','%d','%d')"
	GetPointByUserId = "SELECT SUM(points) AS points FROM articles WHERE articles.created_by='%d'"
)
