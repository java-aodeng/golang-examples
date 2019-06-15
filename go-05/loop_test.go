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

}