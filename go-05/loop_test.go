package go_05

import "testing"

func TestWhileLoop(t *testing.T)  {
	n:=0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestCondition(t *testing.T)  {
	if a :=1==1;a {
		t.Log(a,"1")
	}else {
		t.Log(a,"2")
	}
}

func TestSwitch(t *testing.T)  {
	for i:=0;i<5 ;i++  {
		switch i {
		case 0,2:
			t.Log("111111")
		case 1,3:
			t.Log("222222")
		default:
			t.Log("000000")
		}
	}
}

func TestSwitch2(t *testing.T)  {
	for i:=0;i<5 ;i++  {
		switch {
		case i%2==0:
			t.Log("111111%")
		case i%2==1:
			t.Log("222222%")
		default:
			t.Log("000000")
		}
	}
}