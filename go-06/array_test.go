package go_06

import "testing"

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