package q

import "testing"

// func TestHelloEmptyArg(t *testing.T) {
// 	emptyResult := Hello("")
// 	if emptyResult != "Hello Dude!" {
// 		t.Errorf("hello(\"\")failed,expected %v,got %v", "Hello Dude!", emptyResult)
// 	} else {
// 		t.Logf("hello(\"\") Success,expected %v,got %v", "Hello Dude!", emptyResult)
// 	}
// }
func TestHello(t *testing.T) {
	result := Hello("Tatva")
	if result != "Hello Tatva!" {
		t.Errorf("hello(\"Tatva\" failed ,expected %v,got %v", "Hello Dude!", result)
	} else {
		t.Logf("hello(\"Tatva\") Success,expected %v,got %v", "Hello Dude!", result)

	}

}
