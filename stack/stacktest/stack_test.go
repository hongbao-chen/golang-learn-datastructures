package stacktest

import (
	"data-structure/stack/calculate"
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {

	str := "3+2*5+2*2"
	str2 := "3+2*(2+5*5)*2+13"
	str3 := "3+2*(2+5*5+2)*2+13"
	//3+2*2+13=20
	//3+2*(2+5*5)*2+13=124
	//3+2*(2+5*5+2)*2+13 = 132

	res := calculate.Calculate(str)
	res2 := calculate.Calculate(str2)
	res3 := calculate.Calculate(str3)

	fmt.Printf("%s=%d\n", str, res)
	fmt.Printf("%s=%d\n", str2, res2)
	fmt.Printf("%s=%d\n", str3, res3)

	//  结果：
	//	E:\goProject\data-structure-and-algorithms\stack\stacktest>go test
	//	3+2*5+2*2=17
	//	3+2*(2+5*5)*2+13=124
	//	3+2*(2+5*5+2)*2+13=132
	//	PASS
	//	ok      data-structure/stack/stacktest  0.159s

}
