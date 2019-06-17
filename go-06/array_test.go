package go_06

import "testing"

//数组
func TestArray(t *testing.T)  {
	var arr [3]int
	t.Log(arr[1])

	arr1:=[...]int{1,2,3}
	t.Log(arr1[1])
}

func TestArrayTravel(t *testing.T)  {
	arr:=[...]int{1,2,3,4}

	for idx,e:=range arr{
		t.Log(idx,e)
	}

	t.Log(arr[1:len(arr)])
}

//切片
func TestSliceInit(t *testing.T)  {
	var s[] int
	t.Log(len(s),cap(s))
	s=append(s,1)
	t.Log(len(s),cap(s))

	s1:=[] int{1,2,3}
	t.Log(len(s1),cap(s1))

	s2:=make ([]int,2,3)
	t.Log(len(s2),cap(s2))
}