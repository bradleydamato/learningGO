package main 

import "fmt"
import "strings"

//basename removes direcotry components and a .suffix

//e.g. a => a , a.go > a, a/b/c.go > c 

func basename(s string) string{
	//discard last '/' and everthing before 
	for i := len(s) - 1; i >= 0 ; i--{
		if s[i] == '/' {
			s = s[i+1:]
			break
				}
					}
	for i := len(s)-1; i>=0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break 
				}
					}
			return s

				}

func basenameshort(s string) string {
	slash := strings.LastIndex(s,"/") // returns -1 if there is not '/'
	s = s[slash+1:]
	if dot := strings.LastIndex(s,"."); dot >= 0{
			s = s[:dot]
						}
			return s


}

func comma (s string) string {
	n := len(s)
	if n <= 3{
		return s
}
	return comma(s[:n-3]) + "," + s[n-3:]
}




func main() {
	test_string := "home/brad/hi.go"
	test_number_string := "123456789"
	fmt.Println(basename(test_string))
	fmt.Println(basenameshort(test_string))
	fmt.Println(comma(test_number_string))
		
}
