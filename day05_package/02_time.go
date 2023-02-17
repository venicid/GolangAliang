package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	time三种类型
		时间对象： golang中定义的一个对象
		格式化时间： 给人看的时间（2022-03-27 09:15:29）
		时间戳： 秒的整数形式
	*/
	var objTime time.Time
	var strTime string
	var stampTime int64

	objTime = time.Now()
	strTime = objTime.Format("2006-01-02 15:04:05") // 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	stampTime = objTime.Unix()                      // 时间戳(从1970年1月1日 开始到现在多少s的时间)

	fmt.Println(objTime)   // 2022-11-05 11:08:11.1100483 +0800 CST m=+0.003172901
	fmt.Println(strTime)   // 2022-11-05 11:08:11
	fmt.Println(stampTime) // 1667617691
	fmt.Println()

	/**
	互相转换
	*/
	// 时间对象 --> 时间戳
	stampTime = objTime.Unix()

	// 时间对象 --> 格式化时间
	strTime = objTime.Format("2006-01-02 15:04:05")

	year := objTime.Year()                                                              //年
	month := objTime.Month()                                                            //月
	day := objTime.Day()                                                                //日
	hour := objTime.Hour()                                                              //小时
	minute := objTime.Minute()                                                          //分钟
	second := objTime.Second()                                                          //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) // 2021-05-19 09:20:06
	// 注意：%02d 中的 2 表示宽度，如果整数不够 2 列就补上 0

	// 时间戳 --> 时间对象
	timeObj := time.Unix(stampTime, 0)
	fmt.Println(timeObj) // 2022-11-05 10:55:02 +0800 CST

	// 格式化时间  ---> 时间对象
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, loc) // 按照指定时区和指定格式解析字符串时间
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeObj)

	/**
	获取当前时间，几分钟前，几分钟后
	*/

	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	durationTime, _ := time.ParseDuration("-1m")
	beforeObjTime := objTime.Add(durationTime)
	fmt.Println(beforeObjTime)

	durationTime, _ = time.ParseDuration("1h20m")
	afterObjTime := objTime.Add(durationTime)
	fmt.Println(afterObjTime)

	/**
	获取两个时间的间隔
	time.Duration 是 time 包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
	const (
		Nanosecond Duration = 1
		Microsecond = 1000 * Nanosecond
		Millisecond = 1000 * Microsecond
		Second = 1000 * Millisecond
		Minute = 60 * Second
		Hour = 60 * Minute
	)
	*/
	duration := objTime.Sub(afterObjTime)
	fmt.Println(duration)
	fmt.Println(duration.Hours())
	fmt.Println(duration.Minutes())
	fmt.Println(duration.Seconds())

	// 2天2小时2分钟2秒后
	var twoDayTwoHour = 2*24*time.Hour + 2*time.Hour + 2*time.Minute
	fmt.Println(twoDayTwoHour) // 50h2m0s
	durationTime, _ = time.ParseDuration(twoDayTwoHour.String())
	twoDayTwoHourObjTime := objTime.Add(durationTime)
	fmt.Println(twoDayTwoHourObjTime) // 2022-11-07 13:42:15.0850399 +0800 CST m=+180120.004117901
}
