package main

import "fmt"
//struct
type Book struct {
	Name string
	isbn int
	Author_name string
}
//method
func add(x int,y int) int   {
	return x+y
}
func main() {
	//variables
	var1:=10
	fmt.Println(var1)
	var var2 int=4
	fmt.Println(var1+var2)
	//pointer
	
	p:= & var2
	fmt.Println(p,*p)

	//loop
	for i:=0;i<=10;i++{
		fmt.Println("%d\n",i)
	}

	//array,slice
	var my_slice = []int{1,2,3,4}
	fmt.Println(my_slice)
	my_slice=append(my_slice,90)
	fmt.Println(my_slice)
	//if
	age:=18
	if age>=18{
		fmt.Println("Eligible to vote")
	}else {
		fmt.Println("Not Eligible to vote")
	}
	//struct
	var book Book
	book.Name="Wings of Fire"
	book.isbn=1312
	book.Author_name="APJ Abdul Kalam"
	fmt.Println(book.Name,book.isbn,book.Author_name)
	var book2 = Book{Name:"Narcos",isbn:1211,Author_name:"Pablo Escobar"}
	fmt.Println(book2)
	//range
    a := [10]int{1,2,3,4,5,6,7,8,9,10}
	for index,values:=range a{
		fmt.Printf("a[%d]:%d\n",index,values)
	}
	//map
	m:=make(map[int]string)
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"
	m[4] = "Four"
	m[5] = "Five"
	fmt.Println(m[4])
	
	//add
	fmt.Println(add(1,3))
}