# SliceUtils

Slice utility library for Golang !

### Get
```
go get github.com/tspn/sliceutils
```

### Usage
```go
  v := []int{4, 3, 2, 1}
	su := sliceutils.Int(v)

	fmt.Println(su.Len(), su)                                             //4 [4 3 2 1]
	fmt.Println(su[2], su)                                                //2 [4 3 2 1]
	fmt.Println(su.Reverse(), su)                                         //[1 2 3 4] [4 3 2 1]
	fmt.Println(su.Push(1), su)                                           //[4 3 2 1 1] [4 3 2 1]
	fmt.Println(su.UnShift(1), su)                                        //[1 4 3 2 1] [4 3 2 1]
	fmt.Println(su.Pop(), su)                                             //1 [4 3 2]
	fmt.Println(reflect.DeepEqual([]int{4, 3, 2}, su.GetInt()))           //true

	s := []string{"Hello", "World"}
	su = sliceutils.String(s)
	fmt.Println(su.Reduce(func(value, sum interface{}) interface{} {      //Hello World
		sum = fmt.Sprintf("%s %s", sum.(string), value.(string))
		return sum
	}))

	f := []float64{1.1, 2.2, 3.3}
	su = sliceutils.Float64(f)
	su.Map(func(i interface{}) interface{} {
		return i.(float64) * 2
	}).ForEach(func(i interface{}) {                                      //2.2
		fmt.Println(i.(float64))                                            //4.4
	})                                                                    //6.6


```
