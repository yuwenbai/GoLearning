package testins

import "fmt"

var ts *testinsc

type testinsc struct {
	number int
}

func (o *testinsc) SetNumber(num int) {
	o.number = num
}
func (o *testinsc) OutNumber() {
	fmt.Printf("OutNumber %v ", o.number)
}

func TestF(num int) {
	if ts == nil {
		ts = new(testinsc)
	}
	ts.SetNumber(num)
}
func TestX() {
	ts.OutNumber()
}
