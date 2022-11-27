package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"user/config"
	"user/database"
	e "user/entity"
	"user/models"
	"user/utils"

	"github.com/labstack/echo/v4"
)

var (
	article e.Article
	user    e.User
	query   string
	data    []interface{}
)

func InsertArticle(c echo.Context) (err error) {
	var (
		rows       sql.Result
		title      string
		content    string
		created_by int
	)
	currentTime := time.Now()
	c.Bind(&article)

	article.CreatedAt = fmt.Sprintf("%d-%d-%d", currentTime.Year(), currentTime.Month(), currentTime.Day())

	if !utils.IsValidAlphaNumericHyphen(c.FormValue("title")) {
		res := models.SetResponse(http.StatusBadRequest, "title can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		title = c.FormValue("title")
	}

	if c.FormValue("content") == "" {
		res := models.SetResponse(http.StatusBadRequest, "content can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		content = c.FormValue("content")
	}

	if c.FormValue("created_by") == "" {
		res := models.SetResponse(http.StatusBadRequest, "created_by can't be empty", []interface{}{})
		return c.JSON(http.StatusOK, res)
	} else {
		if i, _ := strconv.Atoi(c.FormValue("created_by")); i == 0 {
			res := models.SetResponse(http.StatusBadRequest, "created_by can't be zero", []interface{}{})
			return c.JSON(http.StatusOK, res)
		} else {
			created_by = i
		}
	}

	query = fmt.Sprintf(database.InsertArticle, article.Id, title, content, article.CreatedAt, created_by, 0)
	rows, err = config.DB.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return
	}

	id, _ := rows.LastInsertId()
	article.Id = int(id)
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")
	article.CreatedBy = created_by

	data = append(data, article)

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "failed to insert article", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func GetArticleByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		res := models.SetResponse(http.StatusBadRequest, "article id must be an integer", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}

	query = fmt.Sprintf(database.GetArticleById, id)
	err = config.DB.QueryRow(query).Scan(&article.Id, &article.Title, &article.Content, &article.CreatedAt, &article.CreatedBy, &article.Points)
	if err != nil {
		log.Println(err.Error())
	}

	data = append(data, article)

	if len(data) > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "no data was found", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
func GetPointsByUserId(c echo.Context) (err error) {
	var (
		article e.Article
		query   string
		data    []interface{}
		points  e.Points
	)

	id, _ := strconv.Atoi(c.FormValue("id"))

	query = fmt.Sprintf(database.GetPointByUserId, id)
	err = config.DB.QueryRow(query).Scan(&article.Points)
	if err != nil {
		log.Println(err.Error())
	}
	points.Points = article.Points

	data = append(data, points)

	if len(data) > 0 {
		res := models.SetResponse(http.StatusOK, "success", data)
		return c.JSON(http.StatusOK, res)
	} else {
		res := models.SetResponse(http.StatusBadRequest, "No Data Was Found", []interface{}{})
		return c.JSON(http.StatusOK, res)
	}
}
