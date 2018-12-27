package main

import (
	"github.com/sunweiconfidence/prepare/src/github.com/gorhill/cronexpr"
	"fmt"
	"time"
)

func main() {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)
    //每一秒执行一次
	//if expr,err = cronexpr.Parse("* * * * * * *");err!=nil {
		//fmt.Println(err)
		//return
	//}

	//linux crontab
	//秒粒度，年配置(2018-2099)
	//哪一秒，哪一分钟，哪小时，哪天，哪月，哪年，星期几
	//每隔5分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *");err!=nil {
		fmt.Println(err)
		return
	}

	//当前时间
	now = time.Now()
	//下次调度时间
	nextTime = expr.Next(now)

	time.AfterFunc(nextTime.Sub(now),func(){
		fmt.Println("被调度了:",nextTime)
	})
	time.Sleep(5*time.Second)
	//fmt.Println(now,nextTime)
	//expr = expr



}
