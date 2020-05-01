package main

import (
	"reflect"
	"testing"
)

func TestLazyList(t *testing.T) {
	var ll List = NewLazyList(EvalerFunc(testCounter), Int(0))
	res := make([]Expr, 0, 10)
	for !ll.Empty() {
		val, err := ll.Head()
		if err != nil {
			t.Fatalf("Head() failed: %v", err)
		}
		res = append(res, val)
		ll, err = ll.Tail()
		if err != nil {
			t.Fatalf("Tail() failed: %v", err)
		}
	}
	exp := []Expr{
		Int(1),
		Int(2),
		Int(3),
		Int(4),
		Int(5),
		Int(6),
		Int(7),
		Int(8),
		Int(9),
		Int(10),
	}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Incorrect sequence generated by LazyList: expected %v, actual %v", exp, res)
	}
}

// 1..10 generator
func testCounter(args []Expr) (Expr, error) {
	prev := int64(args[0].(Int))
	if prev >= 10 {
		return QEmpty, nil
	}
	return &Sexpr{
		List:   []Expr{Int(prev + 1)},
		Quoted: true,
	}, nil
}

func TestLazyListState(t *testing.T) {
	var ll List = NewLazyList(EvalerFunc(fibGen), QList(Int(1), Int(1)))
	res := make([]Expr, 0, 10)
	for i := 0; i < 6; i++ {
		val, err := ll.Head()
		if err != nil {
			t.Fatalf("Head() failed: %v", err)
		}
		res = append(res, val)
		ll, err = ll.Tail()
		if err != nil {
			t.Fatalf("Tail() failed: %v", err)
		}
	}
	exp := []Expr{
		Int(1),
		Int(2),
		Int(3),
		Int(5),
		Int(8),
		Int(13),
	}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Incorrect sequence generated by LazyList: expected %v, actual %v", exp, res)
	}
}

func fibGen(args []Expr) (Expr, error) {
	state := args[0].(*Sexpr)
	a := int64(state.List[0].(Int))
	b := int64(state.List[1].(Int))
	a, b = b, a+b
	return &Sexpr{
		List: []Expr{
			Int(a),
			QList(Int(a), Int(b)),
		},
		Quoted: true,
	}, nil
}

func TestLazyListFiniteString(t *testing.T) {
	ll := NewLazyList(EvalerFunc(testCounter), Int(0))
	act := ll.String()
	exp := "'(1 2 3 4 5 6 7 8 9 10)"
	if act != exp {
		t.Errorf("Incorrect String() result of LazyList: expected %q, actual %q", exp, act)
	}
}
