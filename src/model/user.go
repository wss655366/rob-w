package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"` // omitempty值为空时忽略该字段解析
	Account  string        `json:"account"`
	Password string        `json:"password"`
	Name     string        `json:"name"`
	Age      int           `json:"age"`
	Email    string        `json:"email"`
	Friends  []string      `json:"friends"`  // 数组
	Comments []Comments    `json:"comments"` // 内嵌数组文档
	Address  Address       `json:"address"`  // 内嵌文档
	// 数据库私有字段
	CreateAt string `json:"create_at" bson:"create_at"`
	ModifyAt string `json:"modify_at" bson:"modify_at"`
	IsDelete bool   `json:"is_delete" bson:"is_delete"`
	DeleteAt string `json:"delete_at" bson:"delete_at"`
}

type Address struct {
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Remark   string `json:"remark"`
}

type Comments struct {
	Id      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Content string        `json:"content"`
	UserRef mgo.DBRef     `json:"user_ref" bson:"user_ref,omitempty"`
	// 数据库私有字段
	CreateAt string `json:"create_at" bson:"create_at"`
	ModifyAt string `json:"modify_at" bson:"modify_at"`
	IsDelete bool   `json:"is_delete" bson:"is_delete"`
	DeleteAt string `json:"delete_at" bson:"delete_at"`
}
