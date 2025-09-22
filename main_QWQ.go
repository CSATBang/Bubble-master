package main

/*

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// 全局数据库连接变量  tip:在生产环境中，考虑依赖注入
var (
	DB *gorm.DB
)

// Todo模型定义
// 使用gorm标签来指定数据库字段属性
type Todo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`   //主键，自增
	Title     string    `gorm:"type:varchar(100);not null"` //非空
	Status    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TodoResponse响应结构体
// 用于api响应，不返回时间字段
type TodoResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 转换函数，将Todo转换成TodoResponse
func (t *Todo) ToResponse() TodoResponse {
	return TodoResponse{
		ID:     t.ID,
		Title:  t.Title,
		Status: t.Status,
	}
}

// ConnectDB 初始化数据库连接
// 返回错误以便在主函数中处理
func ConnectDB() error {
	var err error
	// DSN (Data Source Name) 格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root:xxxxxxx@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect databases:%w", err)
	}
	fmt.Println("DB链接成功!")
	return nil
}
func main() {
	//创建数据库 create database if not exists bubble default charset=utf8;

	//初始化数据库连接
	err := ConnectDB()
	if err != nil {
		log.Fatalf("Databases connection failed :%v", err)
		return
	}
	// 设置路由

	r := gin.Default()
	// 添加中间件（可选）
	// r.Use(gin.Recovery()) // 默认已包含
	// r.Use(middleware.Logger())

	//api:v1
	v1Group := r.Group("v1")
	{

		//待办事项API
		//添加待办事项
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			// 从请求中绑定JSON数据
			if err := c.BindJSON(&todo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// 存入数据库 - 使用原生SQL语句
			// 注意：使用GORM的Create方法会更简洁: DB.Create(&todo)
			result := DB.Exec(
				"insert into todos (title,status,created_at,updated_at) VALUES (?,?,?,?)", todo.Title, todo.Status, time.Now(), time.Now())
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			//获取刚插入的ID
			var id uint
			DB.Raw("select last_insert_id()").Scan(&id)
			//返回响应
			c.JSON(http.StatusOK, TodoResponse{
				ID:     id,
				Title:  todo.Title,
				Status: todo.Status,
			})
		})
		//查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todos []TodoResponse
			// 查询所有记录使用原生SQL语句
			// 可以使用GORM的Find方法: DB.Model(&Todo{}).Find(&todos)
			result := DB.Raw("select id,title,status from todos").Scan(&todos)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			//返回空数组而不是null
			if todos == nil {
				todos = []TodoResponse{}
			}
			c.JSON(http.StatusOK, todos)

		})
		//查看某一个
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			var todo TodoResponse
			id := c.Param("id") // 获取URL参数

			result := DB.Raw("select id,title,status from todos where id = ?", id).Scan(&todo)
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

		})
		//修改待办事项状态
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			var todo Todo
			id := c.Param("id")
			// 绑定请求数据
			if err := c.BindJSON(&todo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// 更新数据库
			// 可以使用GORM的Update方法: DB.Model(&Todo{}).Where("id = ?", id).Update("status", todo.Status)
			result := DB.Exec("update todos set status = ?,updated_at = ? where id = ? ", todo.Status, time.Now(), id)
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
		})

		//删除待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			// 删除记录
			// 可以使用GORM的Delete方法: DB.Delete(&Todo{}, id)
			result := DB.Exec("delete from todos where id=?", id)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			// 检查是否找到并删除了记录

			if result.RowsAffected == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "todo not found!"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "status delete successfully!"})
		})
	}
	//启动服务器
	port := ":10090"
	fmt.Printf("Server starting on port %s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed to start:%v", err)
	}
}
*/
