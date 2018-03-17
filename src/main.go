/*
 * 说明：Tutorial for Mongodb based on Golang and Mgo
 * 时间：2018-01-17 22:55
 * 更新：
 */
package main

import (
	"fmt"

	"new_preoject/src/dao"
)

func main() {
	session := dao.InitMongo()
	defer session.Close()

	d := dao.NewDao(session)
	userDao := dao.NewUserDao(d)

	if err := userDao.TestMgoError(); err != nil {
		fmt.Println(err)
	}
}
