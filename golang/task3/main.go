package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// Student 学生表（students）
type Student struct {
	Id        uint
	Name      string
	Age       int
	Grade     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Employee 员工表（employees）
type Employee struct {
	Id         uint
	Name       string
	Department string
	Salary     float64
}

// Book 书籍表（books）
type Book struct {
	Id     uint
	Title  string
	Author string
	Price  float64
}

// Account 账户表（accounts）
type Account struct {
	Id      uint
	Balance float64
}

// Transaction 事物操作表（transactions）
type Transaction struct {
	Id            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/

// User （用户）
type User struct {
	Id         uint   `gorm:"primaryKey"`
	Username   string `gorm:"uniqueIndex;size:50;not null"`
	Email      string `gorm:"uniqueIndex;size:100;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Posts      []Post         `gorm:"foreignKey:AuthorID"`
	PostsCount int            `gorm:"default:0"`
}

// Post （文章）
type Post struct {
	Id            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:200;not null"`
	Content       string `gorm:"type:text;not null"`
	AuthorID      uint   `gorm:"index;not null"` //外键
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Comments      []Comment      `gorm:"foreignKey:PostID"`
	CommentStatus string         `gorm:"type:enum('无评论','有评论');default:'无评论'"`
}

// Comment （评论）
type Comment struct {
	Id        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   `gorm:"index;not null"` //外键
	PostID    uint   `gorm:"index;not null"` //外键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 数据库连接
func getConnect() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	return db
}

// 数据库连接
func getConnect1() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	return db
}

// 数据库连接 sqlx
func getConnectSqlx() *sqlx.DB {
	// sqlx 的连接
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	return db
}

// 实例化数据库表
func InitDB(db *gorm.DB) {
	//err := db.AutoMigrate(&Student{})
	//if err == nil {
	//	fmt.Println("Student创建成功！")
	//} else {
	//	panic(err)
	//}

	//err := db.AutoMigrate(&Account{}, &Transaction{})
	//if err == nil {
	//	fmt.Println("Account、Transaction创建成功！")
	//} else {
	//	panic(err)
	//}

	//err := db.AutoMigrate(&Employee{})
	//if err == nil {
	//	var employees []Employee
	//	db.Find(&employees)
	//	if len(employees) == 0 {
	//		employees = []Employee{
	//			{Name: "张三", Department: "技术部", Salary: 8000.00},
	//			{Name: "李四", Department: "销售部", Salary: 18000.00},
	//			{Name: "王五", Department: "技术部", Salary: 98000.00},
	//			{Name: "老六", Department: "售后部", Salary: 6000.00},
	//			{Name: "祝东钊", Department: "技术部", Salary: 16000.00},
	//		}
	//		db.Create(&employees)
	//	}
	//	fmt.Println("Employee创建成功！")
	//} else {
	//	panic(err)
	//}

	//err := db.AutoMigrate(&Book{})
	//if err == nil {
	//	var books []Book
	//	db.Find(&books)
	//	if len(books) == 0 {
	//		books = []Book{
	//			{Title: "C语言开发", Author: "张三", Price: 30.00},
	//			{Title: "java语言开发", Author: "李四", Price: 80.00},
	//			{Title: "go语言开发", Author: "王五", Price: 100.00},
	//			{Title: "php语言开发", Author: "赵六", Price: 40.00},
	//		}
	//		db.Create(&books)
	//	}
	//	fmt.Println("Book创建成功！")
	//} else {
	//	panic(err)
	//}

	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err == nil {
		db.Create(GetSampleUsers())
		db.Create(GetSamplePosts())
		db.Create(GetSampleComments())

		fmt.Println("User、Post、Comment创建成功！")
	} else {
		panic(err)
	}
}

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
func queryStudent(db *gorm.DB) {
	// 编写SQL语句向 students 表中插入一条新记录
	student := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	result := db.First(&student)
	if result.RowsAffected == 0 {
		db.Create(&student)
	} else {
		result0 := db.Model(&student).Update("Grade", "三年级")
		if result0.Error != nil {
			panic("修改失败！" + result0.Error.Error())
		} else {
			fmt.Println("修改成功！", result0)
		}
	}
	// 编写SQL语句向 students 表中插入多条新记录
	//students := []Student{{Name: "李四", Age: 20, Grade: "一年级"}, {Name: "王五", Age: 25, Grade: "五年级"}, {Name: "祝东钊", Age: 16, Grade: "一年级"}}
	//db.Create(&students)

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	result1 := db.Model(&Student{}).Where("name =?", "张三").Update("Grade", "四年级")
	if result1.Error != nil {
		panic("修改失败！" + result1.Error.Error())
	} else {
		fmt.Println("修改成功！", result1)
	}

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	students := []Student{}
	db.Where("age > ?", 18).Find(&students)
	fmt.Println("students 表中所有年龄大于 18 岁的学生信息数：", len(students), students)

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	result2 := db.Where("age < ?", 15).Delete(&Student{})
	if result2.Error != nil {
		panic("删除失败！" + result2.Error.Error())
	} else {
		fmt.Println("删除成功！", result2)
	}
}

/*
假设有两个表：
accounts 表（包含字段 id 主键， balance 账户余额）和transactions 表（包含字段 id 主键，from_account_id 转出账户ID，to_account_id 转入账户ID，amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func transferAccount(db *gorm.DB) {
	// 添加A、B两个账户
	//accounts := []*Account{
	//	{Balance: 300.0},
	//	{Balance: 80.0},
	//}
	//db.Create(accounts)

	err := transferMoney(1, 2, db)
	if err != nil {
		fmt.Println("转账失败：", err)
	} else {
		fmt.Println("转账成功！")
	}
}

func transferMoney(fromId uint, toId uint, db *gorm.DB) error {
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// 先查询转出账号中金额是否足够
	var account Account
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account, fromId)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败：%v", result.Error)
	}
	if account.Balance < 100 {
		tx.Rollback()
		return fmt.Errorf("账户：%d中余额不足！", fromId)
	}
	// 查询转入账号
	var toAccount Account
	result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&toAccount, toId)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("查询转入账户失败：%v", result.Error)
	}
	// 更新余额
	result = tx.Model(&Account{}).Where("id = ?", fromId).Update("balance", gorm.Expr("balance - ?", 100.0))
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新转出账户失败: %v", result.Error)
	}
	result = tx.Model(&Account{}).Where("id = ?", toId).Update("balance", gorm.Expr("balance + ?", 100.0))
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新转入账户失败: %v", result.Error)
	}
	// 写入表
	tr := Transaction{
		FromAccountId: fromId,
		ToAccountId:   toId,
		Amount:        100,
	}
	if result := tx.Create(&tr); result.Error != nil {
		return fmt.Errorf("创建交易记录失败：%v", result.Error)
	}
	return tx.Commit().Error
}

/*
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func getEmployee(db *sqlx.DB) {
	// 查询所有 "技术部" 的员工信息
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees where department = ? ORDER BY name", "技术部")
	if err != nil {
		panic(err)
	}
	fmt.Println("技术部员工:", employees)

	//查询工资最高的员工
	var employee Employee
	errr := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if errr != nil {
		panic(errr)
	}
	fmt.Println("工资最高的员工:", employee)
}

/*
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
func getBook(db *sqlx.DB) {
	var books []*Book
	err := db.Select(&books, "SELECT * FROM books where price > ?", 50.0)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Printf("书名：%s，作者：%s，价格：%f\n", book.Title, book.Author, book.Price)
	}
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func getUserQuery(db *gorm.DB) {
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	var user User
	var username string = "tech_guru"
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	var post []*Post
	result = db.Where("author_id = ?", user.Id).Find(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	for _, p := range post {
		fmt.Printf("%s发布的所有文章及其对应的评论信息:%s\n", username, p)
	}

	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	subQuery := db.Model(&Comment{}).
		Select("post_id, COUNT(*) AS comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Limit(1)
	var maxCommentPost Post
	result = db.Debug().Joins("JOIN (?) AS max_post ON posts.id = max_post.post_id", subQuery).Find(&maxCommentPost)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("查询评论数量最多的文章信息:", maxCommentPost)
}

/*
题目3：钩子函数,
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。,
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func initBookDB(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err == nil {
		fmt.Println("User、Post、Comment创建成功！")
	} else {
		panic(err)
	}

	// 创建用户
	user := User{Username: "张三122UUU", Email: "zhangsan@example.com111UUU"}
	db.Create(&user)

	// 创建文章 (触发AfterCreate钩子)
	post := Post{
		Title:    "我的第一篇博客",
		Content:  "欢迎来到我的博客...",
		AuthorID: user.Id,
	}
	db.Create(&post) // 此时用户的PostsCount自动+1

	// 创建评论
	comment := Comment{
		Content:  "很好的文章!",
		PostID:   post.Id,
		AuthorID: user.Id,
	}
	db.Create(&comment)

	// 删除评论 (触发AfterDelete钩子)
	db.Delete(&comment) // 删除后检查文章状态
}

// AfterCreate 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。,
func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	result := db.Model(&User{}).Where("id = ?", p.AuthorID).Update("posts_count", gorm.Expr("posts_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// AfterDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	var commentCount int64
	result := db.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount)
	if result.Error != nil {
		return result.Error
	}
	status := "有评论"
	if commentCount == 0 {
		status = "无评论"
		result = db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", status)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func main() {
	db := getConnect1()
	//InitDB(db)
	//queryStudent(db)
	transferAccount(db)
	//dbSqlx := getConnectSqlx()
	//getEmployee(dbSqlx)
	//getBook(dbSqlx)
	//getUserQuery(db)
	//initBookDB(db)
}

func GetSampleUsers() []User {
	return []User{
		{
			Id:       1,
			Username: "tech_guru",
			Email:    "guru@tech.com",
		},
		{
			Id:       2,
			Username: "code_master",
			Email:    "master@code.dev",
		},
		{
			Id:       3,
			Username: "dev_learner",
			Email:    "learner@dev.net",
		},
	}
}

func GetSamplePosts() []Post {
	return []Post{
		{
			Id:       1,
			Title:    "深入理解Golang并发模型",
			Content:  "Go语言的并发模型是其最强大的特性之一...",
			AuthorID: 1,
		},
		{
			Id:       2,
			Title:    "GORM高级技巧大全",
			Content:  "本文将介绍GORM的各种高级用法和最佳实践...",
			AuthorID: 1,
		},
		{
			Id:       3,
			Title:    "从零构建RESTful API",
			Content:  "使用Go和Gin框架构建高性能API服务...",
			AuthorID: 2,
		},
		{
			Id:       4,
			Title:    "数据库优化实战",
			Content:  "如何优化SQL查询提升应用性能...",
			AuthorID: 2,
		},
		{
			Id:       5,
			Title:    "微服务架构设计模式",
			Content:  "微服务架构的常见模式和反模式...",
			AuthorID: 3,
		},
	}
}

func GetSampleComments() []Comment {
	return []Comment{
		{
			Id:       1,
			Content:  "非常有深度的文章！",
			AuthorID: 2,
			PostID:   1,
		},
		{
			Id:       2,
			Content:  "期待更多关于channel的内容",
			AuthorID: 3,
			PostID:   1,
		},
		{
			Id:       3,
			Content:  "GORM的关联查询确实很方便",
			AuthorID: 1,
			PostID:   2,
		},
		{
			Id:       4,
			Content:  "解决了我的实际问题",
			AuthorID: 3,
			PostID:   3,
		},
		{
			Id:       5,
			Content:  "优化后性能提升明显",
			AuthorID: 1,
			PostID:   4,
		},
		{
			Id:       6,
			Content:  "实例代码能否分享一下？",
			AuthorID: 3,
			PostID:   4,
		},
		{
			Id:       7,
			Content:  "架构设计思路很清晰",
			AuthorID: 2,
			PostID:   5,
		},
	}
}
