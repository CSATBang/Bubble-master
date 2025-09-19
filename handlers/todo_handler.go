package handlers

import (
	"buble-master/database"
	"buble-master/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	// 从请求中绑定JSON数据
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 存入数据库 - 使用原生SQL语句
	// 注意：使用GORM的Create方法会更简洁: DB.Create(&todo)
	result := database.DB.Exec(
		"insert into todos (title,status,created_at,updated_at) VALUES (?,?,?,?)", todo.Title, todo.Status, time.Now(), time.Now())
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//获取刚插入的ID
	var id uint
	database.DB.Raw("select last_insert_id()").Scan(&id)
	//返回响应
	c.JSON(http.StatusCreated, models.TodoResponse{
		ID:     id,
		Title:  todo.Title,
		Status: todo.Status,
	})
}

func GetAllTodos(c *gin.Context) {
	var todos []models.TodoResponse
	// 查询所有记录使用原生SQL语句
	// 可以使用GORM的Find方法: DB.Model(&Todo{}).Find(&todos)
	result := database.DB.Raw("select id,title,status from todos").Scan(&todos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	//返回空数组而不是null
	if todos == nil {
		todos = []models.TodoResponse{}
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	var todo models.TodoResponse
	id := c.Param("id") // 获取URL参数

	result := database.DB.Raw("select id,title,status from todos where id = ?", id).Scan(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	// 检查记录是否存在
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found!"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodoStatus(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	// 绑定请求数据
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 更新数据库
	// 可以使用GORM的Update方法: DB.Model(&Todo{}).Where("id = ?", id).Update("status", todo.Status)
	result := database.DB.Exec("update todos set status = ?,updated_at = ? where id = ? ", todo.Status, time.Now(), id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	// 检查是否找到并更新了记录
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "status updated successfully!"})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	// 删除记录
	// 可以使用GORM的Delete方法: DB.Delete(&Todo{}, id)
	result := database.DB.Exec("delete from todos where id=?", id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	// 检查是否找到并删除了记录

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found!"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "status delete successfully!"})
}
