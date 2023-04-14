package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func GolangBase() {
	// 局部匿名函数
	func(s string) {
		fmt.Println(s)
	}("hello golang...")
	func1 := func(s string) {
		fmt.Println(s)
	}
	func1("hello golang...")
	funcTest1(func(s string) {
		fmt.Println(s)
	})
	funcRes := funcTest2()
	funcRes(1, 2)
	// 闭包
	num := 25
	func2 := func() {
		fmt.Printf("Age is %d...\n", num)
	}
	func2()

	res1 := funcTest3() // 得到一个闭包
	fmt.Println(res1()) // 2
	fmt.Println(res1()) // 3
	fmt.Println(res1()) // 4
	fmt.Println(res1()) // 5 只要闭包还在使用外界变量 外界变量一直存在 所以可以一直x++

	// defer语句
	fmt.Println("mainFunc: ", funcTest4()) // 2

	// 数组
	arr1 := [5]int{0: 9, 3: 2}
	fmt.Println(arr1)
	arr2 := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Println(arr2)
	arr3 := [...][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Println(arr3)

	// 切片
	slc1 := make([]int, 3, 5)
	fmt.Printf("slc1: %v, %p\n", slc1, slc1)

	slc1 = append(slc1, []int{1, 1, 1, 1}...)
	fmt.Printf("slc1: %v, %p\n", slc1, slc1)

	slc2 := make([]int, len(slc1)+1)
	copy(slc2, slc1)
	fmt.Printf("slc2: %v, %p\n", slc2, slc2)

	// 数值、字符串转换
	// 数值->字符串
	var num1 int32 = 10
	str1 := strconv.FormatInt(int64(num1), 10) // FormatInt(i int64, base int) i为int64类型需要强转 base为转换的进制（base \in [2, 36]）
	fmt.Println("Int convert to string: ", str1)
	str2 := strconv.FormatInt(int64(num1), 2)
	fmt.Println("Int convert to string: ", str2)

	var num2 float64 = 3.75686756756756756756
	// FormatFloat(f float64, fmt byte, prec int, bitSize int)
	// f 必须为64位；fmt 格式匹配(f 小数格式，e 指数个数)；bitSize 输出字符串的精度选择（32 or 64）超出的进行四舍五入
	// prec: precision精度控制位 保留小数点后几位 -1表示保留全部（但不会超出对应float精度）
	str3 := strconv.FormatFloat(num2, 'f', 4, 64)
	fmt.Println("Float convert to string: ", str3) // 3.7569
	str4 := strconv.FormatFloat(num2, 'f', -1, 64)
	fmt.Println("Float convert to string: ", str4) // 3.756867567567568
	str5 := strconv.FormatFloat(num2, 'f', -1, 32)
	fmt.Println("Float convert to string: ", str5) // 3.7568676

	var num3 bool = true
	fmt.Println("Bool convert to string: ", strconv.FormatBool(num3)) // true

	str6 := strconv.Itoa(int(num1))
	fmt.Println(str6) // 10 (string)

	if num4, err := strconv.Atoi(str6); err == nil {
		fmt.Println(num4) // 10 (int)
	}

	// 数值转字符串--Sprintf方法
	str7 := fmt.Sprintf("%d", num1)
	fmt.Println(str7)
	str8 := fmt.Sprintf("%f", num2)
	fmt.Println(str8)
	str9 := fmt.Sprintf("%t", num3)
	fmt.Println(str9)

	// 指针
	ptr0 := new(int) // p0指向int类型 => 存储的是匿名变量的地址
	// 存储匿名变量的地址 匿名变量的值 指针地址
	fmt.Println(ptr0, *ptr0, &ptr0) // 0xc0000b20b0 0 0xc00009c020

	x := *ptr0           // 复制ptr0所引用的值
	ptr1, ptr2 := &x, &x // 这里ptr1 ptr2赋值的是x的地址
	fmt.Println(ptr1, ptr2, ptr0)
	fmt.Println(ptr1 == ptr2) // true
	fmt.Println(ptr1 == ptr0) //false

	ptr3 := &*ptr0            // <=> ptr3 := &(*ptr0) <=> ptr3 := ptr0 这俩存储的地址是一个
	fmt.Println(ptr0 == ptr3) // true

	*ptr0, *ptr1 = 123, 456
	fmt.Printf("%v, %v\n", *ptr0, x)   // 123 456
	fmt.Printf("%T, %T\n", *ptr0, x)   // int int
	fmt.Printf("%T, %T\n", ptr0, ptr1) // *int *int 俩地址
	*ptr2 = 456
	fmt.Printf("%v %v %v\n", x, *ptr1, *ptr2) // 456 456 456
	*ptr3 = 789
	fmt.Printf("%v %v\n", *ptr0, *ptr3) // 789 789

	num4 := 3
	double(&num4)
	fmt.Println(num4) // 6
	ptr4 := &num4
	double(ptr4)
	fmt.Println(num4, *ptr4, ptr4 == nil) // 12 12 false

	ptr5 := nextInt()
	fmt.Println(ptr5, *ptr5) // 0xc000018158 3

	num5 := int64(5)
	ptr6 := &num5
	*ptr6++                                 // *优先级高于++ <=>(*ptr6)++
	fmt.Println(*ptr6, num5)                // 6 6
	fmt.Println(ptr6, &num5, ptr6 == &num5) // 0xc0000b2100 0xc0000b2100 true

	*&num5++                 // 7
	*&*&num5++               // 8
	**&ptr6++                // 9
	*&*ptr6++                // 10
	fmt.Println(*ptr6, num5) // 10 10

	type MyInt int64
	type Ta *int64
	type Tb *MyInt
	// 四个不同类型的指针
	var pa0 Ta
	var pa1 *int64
	var pb0 Tb
	var pb1 *MyInt

	fmt.Println(pa0 == pa1) // true
	fmt.Println(pb0 == pb1) // true
	fmt.Println(pa0 == nil) // true
	fmt.Println(pa1 == nil) // true
	fmt.Println(pb0 == nil) // true
	fmt.Println(pb1 == nil) // true

	// 编译不通过
	// _ = pa0 == pb0
	// _ = pa1 == pb1 // _ = pa1 == (*int64)(pb1) 强转后可以比较
	// _ = pa0 == Tb(nil)

	// reflect 反射
	// 反射基本类型
	funcReflect(num1)

	type person struct {
		name, sex string
		age       int
	}
	p := person{"小王子", "男", 18}
	funcReflect(p) // type: person kind: struct

}
func funcTest1(f func(s string)) {
	f("hello golang...")
}
func funcTest2() func(int, int) {
	return func(a, b int) {
		fmt.Println(a + b)
	}
}
func funcTest3() func() int {
	x := 1
	return func() int { x++; return x }
}
func funcTest4() int {
	// defer语句
	d := 1
	defer fmt.Println("First Nums: ", d) // 1
	d++                                  // 此时d=2 return中保留的是该变量值
	defer func() {
		d++
		fmt.Println("Second Nums: ", d) // 3
	}()
	return d
}
func double(x *int) {
	*x += *x
	x = nil // 函数体内对传入的指针实参修改 不能反映到函数外【该修改发生在此指针的一个副本上】
}

func nextInt() *int {
	a := 3
	fmt.Println(&a) // 0xc000018158
	return &a
}
func funcReflect(input interface{}) {
	t := reflect.TypeOf(input) // 获得任意值的类型对象
	fmt.Println("Input's type: ", t)
	fmt.Printf("type: %v kind: %v\n", t.Name(), t.Kind()) // Name()就是对象类型 Kind是种类（底层的类型）
	v := reflect.ValueOf(input)                           // 获取值
	fmt.Println("Input's value: ", v)
}
