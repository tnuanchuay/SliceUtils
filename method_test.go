package sliceutils

import (
	"testing"
	"reflect"
)

func TestSortDesc(t *testing.T){
	d := []int{4, 2, 9, 6, 4, 5, 7}
	su := Int(d)
	su = su.Sort(func (x, y interface{}) bool{
		return x.(int) > y.(int)
	})

	if !reflect.DeepEqual(su.GetInt(), []int{9, 7, 6, 5, 4, 4, 2}){
		t.Error("expect", []int{9, 7, 6, 5, 4, 4, 2}, "actual", su.GetInt())
	}
}

func TestSortAsc(t *testing.T){
	d := []int{4, 2, 9, 6, 4, 5, 7}
	su := Int(d)
	su = su.Sort(func (x, y interface{}) bool{
		return x.(int) < y.(int)
	})

	if !reflect.DeepEqual(su.GetInt(), []int{2, 4, 4, 5, 6, 7, 9}){
		t.Error("expect", []int{2, 4, 4, 5, 6, 7, 9}, "actual", su.GetInt())
	}
}

func TestSwap(t *testing.T){
	d := []int{1, 2, 3, 4}
	su := Int(d)
	su = su.Swap(1, 3)
	if !reflect.DeepEqual(su.GetInt(), []int{1, 4, 3, 2}){
		t.Error("expect", []int{1, 4, 3, 2}, "actual", su.GetInt())
	}
}

func TestReverse(t *testing.T){
	d := []int{1, 2, 3, 4}
	su := Int(d)
	su = su.Reverse()
	if !reflect.DeepEqual(su.GetInt(), []int{4, 3, 2, 1}) {
		t.Error("expect", []int{4, 3, 2, 1}, "actual", su.GetInt())
	}
}

func TestReduceIf0Member(t *testing.T){
	d := []int{}
	su := Int(d)
	v := su.Reduce(func(v, s interface{}) interface{}{
		return s.(int) + v.(int)
	})

	if v != nil {
		t.Error("expect", nil, "actual", v)
	}
}

func TestReduceIf1Member(t *testing.T){
	d := []int{1}
	su := Int(d)
	v := su.Reduce(func(v, s interface{}) interface{}{
		return s.(int) + v.(int)
	})

	if v != 1 {
		t.Error("expect", 1, "actual", v)
	}
}

func TestReduce(t *testing.T){
	d := []int{1, 2, 3}
	su := Int(d)
	v := su.Reduce(func(v, s interface{}) interface{}{
		return s.(int) + v.(int)
	})

	if v != 6 {
		t.Error("expect", 6, "actual", v)
	}
}

func TestMap(t *testing.T){
	d := []int{1, 2, 3}
	su := Int(d)
	su = su.Map(func (i interface{}) interface{}{
		return i.(int) * 2
	})

	if !reflect.DeepEqual(su.GetInt(), []int{2, 4, 6}){
		t.Error("expect equal")
	}
}

func TestFilterForEach(t *testing.T){
	d := []int{1, 2, 3, 4, 5, 6}
	su := Int(d)
	sum := 0
	su.Filter(func (i interface{}) bool {
		return i.(int) % 2 == 0
	}).ForEach(func (i interface{}){
		sum = sum + i.(int)
	})

	if sum != 12{
		t.Error("expect", 12, "actual", sum)
	}
}

func TestFilter(t *testing.T){
	d := []int{1, 2, 3, 4}
	su := Int(d)
	su = su.Filter(func (i interface{}) bool {
		return i.(int) % 2 == 0
	})

	if su.Len() != 2{
		t.Error("expect", 2, "actual", su.Len())
	}

	su.ForEach(func (i interface{}){
		if i.(int) % 2 != 0{
			t.Error("expect", i, "mod 2 equal 0")
		}
	})
}

func TestForEach(t *testing.T){
	d := []int{1, 2, 3}
	sum := 0
	su := Int(d)
	su.ForEach(func (i interface{}){
		sum += i.(int)
	})

	if sum != 6 {
		t.Error("expect", 6, "actual", sum)
	}
}

func TestShift0Member(t *testing.T){
	d := []int{}
	su := Int(d)
	v := su.Shift()
	if su.Len() != 0{
		t.Error("expect", 0, "actual", su.Len())
	}

	if v != nil{
		t.Error("expect", nil, "actual", v)
	}
}

func TestShift1Member(t *testing.T){
	d := []int{9}
	su := Int(d)
	v := su.Shift()
	if su.Len() != 0{
		t.Error("expect", 0, "actual", su.Len())
	}

	if v != 9{
		t.Error("expect", 9, "actual", v)
	}
}

func TestUnShift(t *testing.T){
	d := []int{2, 3, 4}
	su := Int(d)
	su.UnShift(1)
	if su.Len() != 4{
		t.Error("expect", 4, "actual", su.Len())
	}

	if !reflect.DeepEqual(su.GetInt(), []int{1, 2, 3, 4}){
		t.Error("expect", []int{1, 2, 3, 4}, "actual", su.GetInt())
	}
}

func TestShift(t *testing.T){
	d := []int{9, 2}
	su := Int(d)
	v := su.Shift()
	if su.Len() != 1{
		t.Error("expect", 1, "actual", su.Len())
	}

	if v != 9{
		t.Error("expect", 9, "actual", v)
	}
}

func TestPopIf0Member(t *testing.T){
	d := []int{}
	su := Int(d)
	v := su.Pop()
	if su.Len() != 0{
		t.Error("expect", 0, "actual", su.Len())
	}

	if v != nil{
		t.Error("expect", nil, "actual", v)
	}
}

func TestPopIf1Member(t *testing.T){
	d := []int{1}
	su := Int(d)
	v := su.Pop()
	if su.Len() != 0{
		t.Error("expect", 0, "actual", su.Len())
	}

	if v != 1{
		t.Error("expect", 1, "actual", v)
	}
}

func TestPop(t *testing.T){
	d := []int{1, 2, 3}
	su := Int(d)
	v := su.Pop()
	if su.Len() != 2{
		t.Error("expect", 2, "actual", su.Len())
	}

	if v != 3{
		t.Error("expect", 3, "actual", v)
	}
}

func TestQueue(t *testing.T){
	d := []int{1, 2, 3}
	su := Int(d)
	su.Queue(9)
	if su.Len() != 4 {
		t.Error("expect", 4, "actual", su.Len())
	}

	if su.Tail() != 9{
		t.Error("expect", 9, "actual", su.Tail())
	}
}

func TestPush(t *testing.T){
	d := []int{1, 2, 3}
	su := Int(d)
	su.Push(9)
	if su.Len() != 4 {
		t.Error("expect", 4, "actual", su.Len())
	}

	if su.Tail() != 9{
		t.Error("expect", 9, "actual", su.Tail())
	}
}

func TestTailIfNil(t *testing.T){
	d := []int{}
	su := Int(d)
	if su.Tail() != nil{
		t.Error("expect", nil, "actual", su.Tail())
	}
}

func TestTail(t *testing.T){
	d := []int{1, 2}
	su := Int(d)
	if su.Tail() != 2{
		t.Error("expect", 2, "actual", su.Tail())
	}
}

func TestHeadIfNil(t *testing.T){
	d := []int{}
	su := Int(d)
	if su.Head() != nil{
		t.Error("expect", nil, "actual", su.Head())
	}
}

func TestHead(t *testing.T){
	d := []int{1, 2}
	su := Int(d)
	if su.Head() != 1{
		t.Error("expect", 1, "actual", su.Head())
	}
}

func TestCount(t *testing.T){
	d := []interface{}{1, 2}
	su := Interface(d)
	if su.Count() != 2{
		t.Error("expect", 2, "actual", su.Len())
	}
}

func TestLen(t *testing.T){
	d := []interface{}{1, 2}
	su := Interface(d)
	if su.Len() != 2{
		t.Error("expect", 2, "actual", su.Len())
	}
}