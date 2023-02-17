package main

//
//import "fmt"
//
///**
//为什么要使用接口
//*/
//type MiPhone struct {
//	Name string
//}
//
//func (m *MiPhone) start() {
//	fmt.Println(m.Name, "开机了")
//}
//
//type ApplePhone struct {
//	Name string
//}
//
//func (p *ApplePhone) start() {
//	fmt.Println(p.Name, "开机了")
//}
//
///**
//定义一个Usber接口
//*/
////1.接口是一个规范
//type Usber interface {
//	start()
//	stop()
//}
//
////2.如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
//type Phone struct {
//	Name string
//}
//
////3.手机要实现usb接口的话必须得实现usb接口中的所有方法
//func (p Phone) start() {
//	fmt.Println(p.Name, "启动")
//}
//
//func (p Phone) stop() {
//	fmt.Println(p.Name, "关机")
//}
//
//func main() {
//
//	/**
//	为什么要使用接口
//	*/
//	mi8 := MiPhone{Name: "小米8"}
//	mi8.start() // 小米8 开机了
//
//	apple8 := ApplePhone{Name: "苹果8"}
//	apple8.start() // 苹果8 开机了
//
//	/**
//	定义一个接口
//	*/
//	p := Phone{Name: "华为8"}
//	var p1 Usber
//	p1 = p
//	p1.start()
//	p1.stop()
//
//	/**
//	值接收者 vs 指针接收者
//	*/
//	// 实例化值类型
//	// 会在代码运行时将接收者的值复制一份。
//	phone1 := Phone{Name: "oppo12"}
//	var u1 Usber
//	u1 = phone1
//	u1.start()
//	u1.stop()
//
//	// 实例化指针类型
//	// 修改的是本身的值， 在方法结束后，修改都是有效的
//	phone2 := &Phone{Name: "onePlus12"}
//	var u2 Usber
//	u2 = phone2
//	u2.start()
//	u2.stop()
//	/**
//	使用时机：
//		需要修改接收者中的值。例如 为 Person 添加一个 SetAge 方法，来修改实例变量的年龄
//		接收者是copy代价比较大的对象
//		保证一致性
//	*/
//
//}
