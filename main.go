package main

import (
	"fmt"
	"sort"
	"sync"
)

type Vector[T any] []T

// 自定义泛型参数类型-单一
type Integer interface {
	int
}

// 近似类型元素，只要底层类型是String都是包含在内。
type AnyString interface {
	~string
}

// 联合类型元素
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// 取交集
type MyString1 interface {
	~string
	string
}

func main() {
	//k := MapKeys(map[int]int{1: 2, 2: 4})
	//fmt.Println(k)

	OrderedSliceDemo()

}

// int64orfloat 泛型函数
func int64orfloat[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func ExampleGenericMethod() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("ints result is %v,floats result is %v",
		//可以显示指定泛型类型参数
		//int64orfloat[string, int64](ints),
		//int64orfloat[string, float64](floats),
		//隐式，编译器根据入参类型进行推导，如果没有入参，则必须显示指定
		int64orfloat(ints),
		int64orfloat(floats),
	)
}

// Push 自定义泛型类型也可以添加方法，接收者必须带上约束类型
func (v *Vector[T]) Push(x ...T) {
	*v = append(*v, x...)
}

func ExampleVectorPush() {
	//泛型类型必须通过类型参数实例化才能使用，一旦实例化就不能修改
	var v Vector[string]
	v.Push("hello world", "jjcai")
	fmt.Println(v)
}

func SumInteger[T Integer](a T, b T) T {
	return a + b
}

func SayHi[T AnyString](name T) {
	fmt.Println(name)
}

func DemoSayHi() {
	type MyString string
	var s MyString = "john"
	SayHi(s)
}

// MapKeys 获取map里所有的key切片
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type Set[T comparable] map[T]struct{}

func MakeSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (receiver Set[T]) Add(v T) {
	receiver[v] = struct{}{}
}

func (receiver Set[T]) Del(v T) {
	delete(receiver, v)
}

func (receiver Set[T]) Contains(v T) bool {
	_, ok := receiver[v]
	return ok
}

func (receiver Set[T]) Len() int {
	return len(receiver)
}

func (receiver Set[T]) Iterate(f func(T)) {
	for k := range receiver {
		f(k)
	}
}

func SetDemo() {
	s := MakeSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(2)
	if s.Contains(2) {
		fmt.Println("s contians 2")
	} else {
		fmt.Println("s not contains 2")
	}
	fmt.Println(s)
}

type ThreadSafeSet[T comparable] struct {
	l sync.RWMutex
	m map[T]struct{}
}

// 联合类型元素，根据操作定义 --排序
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// 定义新类型(数组)，使用自定义泛型类型
type orderedSlice[T Ordered] []T

// 实现sort.Interface接口的三个方法
func (s orderedSlice[T]) Len() int {
	return len(s)
}

func (s orderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s orderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// OrderedSlice 封装排序方法
func OrderedSlice[T Ordered](s []T) {
	sort.Sort(orderedSlice[T](s))
}

// OrderedSliceDemo 实现泛型实现自定义切片类型排序
func OrderedSliceDemo() {
	s := []int{1, 3, 2, 4, 10}
	//OrderedSlice(s)
	//fmt.Println(s)

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	fmt.Println(s)
}
