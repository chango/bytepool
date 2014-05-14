package bytepool

import (
	"reflect"
	"testing"
)

type MyThirdPartyData struct {
	A string
	B int
	C bool
}

func CreateNewData() interface{} {
	return &MyThirdPartyData{}
}

func ResetData(item interface{}) error {
	var resetItem *MyThirdPartyData
	var ok bool
	if resetItem, ok = item.(*MyThirdPartyData); ok && resetItem != nil {
		resetItem.A = ""
		resetItem.B = 0
		resetItem.C = false
	}
	return nil
}

// func TestJPoolEachItemIsOfASpecifiedSize(t *testing.T) {
// 	expected := 9
// 	p := New(1, expected)
// 	item := p.Checkout()
// 	defer item.Close()
// 	if cap(item.bytes) != expected {
// 		t.Errorf("expecting array to have a capacity of %d, got %d", expected, cap(item.bytes))
// 	}
// }

func TestArbitraryPoolDynamicallyCreatesAnItemWhenPoolIsEmpty(t *testing.T) {
	p := NewArbitrary(1, CreateNewData, ResetData)
	item1 := p.Checkout()
	item2 := p.Checkout()
	if item2.pool != nil {
		t.Error("The dynamically created item should have a nil pool")
	}

	item1.Close()
	item2.Close()
	if p.Len() != 1 {
		t.Errorf("Expecting a pool length of 1, got %d", p.Len())
	}
	if p.Misses() != 1 {
		t.Errorf("Expecting a miss count of 1, got %d", p.Misses())
	}

}

func TestArbitraryPoolReleasesAnItemBackIntoThePool(t *testing.T) {
	p := NewArbitrary(1, CreateNewData, ResetData)
	item1 := p.Checkout()
	pointer := reflect.ValueOf(item1).Pointer()
	item1.Close()

	item2 := p.Checkout()
	defer item2.Close()
	if reflect.ValueOf(item2).Pointer() != pointer {
		t.Error("Pool returned an unexected item")
	}
}
