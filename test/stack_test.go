package test

import (
	"algo/collection"
	"testing"
)

func TestStack(t *testing.T){
	s := new(collection.Stack)// new(T), return the pointer of instance,the memory of instance is zeroed
	//fmt.Println(reflect.TypeOf(s)) //out:*collection.Stack
	//var s collection.Stack//basic instantiation method
	//fmt.Println(reflect.TypeOf(s))//out collection.Stack

	s.Push(1)

	if s.Empty() {
		t.Error("Empty() Error")
	}else{
		t.Log("Empty() Ok", s.Empty())
	}
	if s.Size() != 1{
		t.Error("Size() Error")
	}else{
		t.Log("Size() Ok", s.Size())
	}

	i := s.Pop()
	if i != 1 {
		t.Error("Pop() Error")
	}else{
		t.Log("Pop() Ok", i)
	}


}