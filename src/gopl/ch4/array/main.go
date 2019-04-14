package main

import "fmt"

/**
 * 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。
 * 数组的长度是固定的，内置的函数len可以返回数组中元素的个数
 */

func main() {
	// 默认情况下，数组的每个元素都被初始化为元素类型对应的零值
	var a [3]int  // array of 3 integers
	fmt.Println(a[0])  // print the first element
	fmt.Println(a[len(a)-1])  // print the last element a[2]

	// Print the indices and elements
	for i, v := range a{
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only
	for _, v := range a{
		fmt.Printf("%d\n", v)
	}

	// 使用数组字面值语法用一组值来初始化数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	// 在数组字面值中，如果在数组的长度位置出现的是“...”省略号，
	// 则表示数组的长度是根据初始化值的个数来计算
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)
	// 数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。
	// 数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。
	q := [3]int{1, 2, 3}
	// q = [4]int{1, 2, 3, 4}  //compile error: cannot assign [4]int to [3]int

	// 指定一个索引和对应值列表的方式初始化
	type Currency int
	const (
		USD Currency = iota //美元
		EUR //欧元
		GBP //英镑
		RMB //人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	// 如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，
	// 这时候我们可以直接通过==比较运算符来比较两个数组，只有当两个数组的所有元素都是相等的时候数组才是相等的。
	// 不相等比较运算符!=遵循同样的规则。
	e := [2]int{1, 2}
	f := [...]int{1, 2}
	g := [2]int{1, 3}
	fmt.Println(e == f, e == g, f == g) // "true false false"
	// h := [3]int{1, 2}
	// fmt.Println(e == h) // compile error: cannot compare [2]int == [3]int
}

// 当调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的参数变量，
// 所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。
// 因为函数参数传递的机制导致传递大的数组类型将是低效的，并且对数组参数的任何的修改都是发生在复制的数组上，
// 并不能直接修改调用时原始的数组变量。我们可以显式地传入一个数组指针，
// 那样的话函数通过指针对数组的任何修改都可以直接反馈到调用者
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}