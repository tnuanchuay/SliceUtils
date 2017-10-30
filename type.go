package sliceutils

type SliceUtils []interface{}

func (c SliceUtils) GetInterface()[]interface{}{
	v := []interface{}{}
	for _, d := range c{
		v = append(v, d.(interface{}))
	}
	return v
}

func (c SliceUtils) GetFloat32()[]float32{
	v := []float32{}
	for _, d := range c{
		v = append(v, d.(float32))
	}
	return v
}

func (c SliceUtils) GetFloat64()[]float64{
	v := []float64{}
	for _, d := range c{
		v = append(v, d.(float64))
	}
	return v
}

func (c SliceUtils) GetString()[]string{
	v := []string{}
	for _, d := range c{
		v = append(v, d.(string))
	}
	return v
}

func (c SliceUtils) GetInt()[]int{
	v := []int{}
	for _, d := range c{
		v = append(v, d.(int))
	}
	return v
}

func Interface(s []interface{})SliceUtils{
	return SliceUtils(s)
}

func Uint(s []uint)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func Float32(s []float32)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func Float64(s []float64)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func Int(s []int)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func Int8(s []int8)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func Int64(s []int64)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func Int32(s []int32)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func Int16(s []int16)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}

func String(s []string)SliceUtils{
	var su SliceUtils = make([]interface{}, len(s))
	for i, d := range s{
		su[i] = d
	}
	return su
}