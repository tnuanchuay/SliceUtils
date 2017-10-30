package sliceutils

import "testing"

//func TestGet(t *testing.T){
//	data := []int{1, 2}
//	su := Int(data)
//}

func TestMakeUint(t *testing.T){
	data := []uint{1, 2}
	su := Uint(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeInt64(t *testing.T){
	data := []int64{1, 2}
	su := Int64(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeInt32(t *testing.T){
	data := []int32{1, 2}
	su := Int32(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeInt16(t *testing.T){
	data := []int16{1, 2}
	su := Int16(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeInt8(t *testing.T){
	data := []int8{1, 2}
	su := Int8(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeInterface(t *testing.T){
	data := []interface{}{1, 2}
	su := Interface(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeFloat32(t *testing.T){
	data := []float32{1, 2}
	su := Float32(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeFloat64(t *testing.T){
	data := []float64{1, 2}
	su := Float64(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeInt(t *testing.T){
	data := []int{1, 2}
	su := Int(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}

func TestMakeString(t *testing.T){
	data := []string{"1", "2"}
	su := String(data)
	if len(su) != 2{
		t.Error("expect", 2, "actual", len(su))
	}
}