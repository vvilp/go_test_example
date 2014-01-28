package main
import "fmt"
import "sort"
import "strings"
func test_switch() {
	i := 2

	switch i {
	case 1:
		fmt.Println("1")	
	case 2:
		fmt.Println("2")
	}
}

func test_if() {
	a := 1
	b := 0
	if a == 1 && b == 0 {
		fmt.Println("a == 1 && b == 0")
	}

	if a == 1 || b == 0 {
		fmt.Println("a == 1 || b == 1")
	}
}

func test_for() {
I:	for i := 0 ; i < 10 ; i++ {
		for j := 0 ; j < 10 ; j++ {
			fmt.Println("i : ", i , "j : " , j)
			if i >= 5 {
				break I
			}
		}
	}
}

func test_array() {
	var a [5]int
	fmt.Println(a)

	//b := [5]int wrong code
	//fmt.Println(b)

	b := [5]int{0,0,0,0,0}
	fmt.Println(b)

	//Array cannot sort like this
	//c := [5]int{1,3,2,5,4} 
	//sort.Ints(c,4)
	//fmt.Println(c)
}

func test_slice() {
	a := make([]int, 3)
	fmt.Println(a)

	b := make([]int, 10)// the slice v now refers to a new array of 10 ints
	fmt.Println(b)
	b = append(b, 123);
	fmt.Println(b)

	c := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(c)
	fmt.Println(c)

	d := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println("d[:5]  : ",d[:5])
	fmt.Println("d[0:5] : ",d[0:5])
	fmt.Println("d[2:5] : ",d[2:5])
	fmt.Println("d[2:]  : ",d[2:])

	//e_array := [5]int{1,3,2,5,4} 
	//e := make([]int, e_array)
	//copy(e,&e_array)

}

func test_map() {
	a := make(map[string]int)
	a["Tom"] = 1;
	a["Tony"] = 2;
	fmt.Println(a)

	b := map[string]int {
		"abc" : 1,
		"qwe" : -1,
	}
	fmt.Println(b)
	fmt.Println(b["qwe"])

	value, present := b["asd"]
	fmt.Println(value, present)


	delete(b,"qwe");
	fmt.Println(b)
}

func test_range() {
	a := []int{1,2,3,4,5,6,7,8,9,0}

	//compare with python
	//for index, each in enumerate(sample):
	for index, num := range a {
		fmt.Println(index, num)
	}

	b := map[string]int {
		"abc" : 1,
		"qwe" : -1,
	}

	for k, v := range b {
		fmt.Println(k, v)
	}
}

func test_variadic_func(nums ...int) {
	for index, num := range nums {
		fmt.Println(index, num)
	}
}

func test_func_var(f func(string) string) {
	fmt.Println(f("qwertyuiop"))
}

func test_func_closure() {
	f := func(s string) string {
		return strings.ToUpper(s)
	}

	test_func_var(f)

	test_func_var(func(s string) string {
			return s + " Append String"
	})
}

func test_string_append(s string) {
	s = s + "  Append String"
}

func test_string_append_p(s *string) {
	*s = *s + "  Append String"
}

func test_pointor() {
	s := "asd"
	test_string_append(s)
	// will not change string
	fmt.Println(s)

	// will change string
	test_string_append_p(&s)
	fmt.Println(s)
}

func test_defer() {
	defer fmt.Println("1") 
	defer fmt.Println("2")
	// print before return
	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
}

func test_func_change_para(a []int) {
	a[0] += 1
	a = append(a, 999)
}
func test_func_change_para_p(a *[]int) {
	*a = append(*a, 999)
}
func test_func(){
	a := []int{0}
	test_func_change_para(a)
	// can not change slice a
	// can change the things in the slice
	fmt.Println(a)

	b := []int{0}
	test_func_change_para_p(&b)
	fmt.Println(b)
	//slice seems as variable to the function
	//so need pass the pointer
}

func test_max_integer() {
	a := 1 << 31 - 1
	fmt.Println(a)
}

func test_string_para(str *string) {
	*str += "qwe"
}

func main() {
	//test_switch()
	//test_if()
	//test_array()
	//test_slice()
	 // test_map()
	//test_range()
	//test_variadic_func(1,2,3,4)
	//test_func_closure()
	//test_pointor()

	//test_for()
	 //test_defer()
	// test_func()
	// test_max_integer()

	str := "123"
	test_string_para(&str)
	fmt.Println(str)
	
}