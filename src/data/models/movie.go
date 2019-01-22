package models

// Movie 是我们的一个简单的数据结构体
// 请注意公共标签（适用于我们web网络应用）
// 应保存在“web / viewmodels / movie.go”等其他文件中
//可以通过嵌入datamodels.Movie或
//声明新字段但我们将使用此数据模型
//作为我们应用程序中唯一的一个Movie模型，
//为了摇摇欲坠。
type Movie struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}