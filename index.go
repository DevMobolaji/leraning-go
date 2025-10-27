package main

import (
	"fmt"
	"strconv"
)

var (
	//Grouping similar variables together
	//Global variables
	MyInt int = 10;
	MyString string = "Beejay"
)

// const  (
// 	x int32 = 10
// 	isRunning = false
// )

func main() {

	
	fmt.Println("Hello Beejay")

	MyVar := "Beejay"
	fmt.Println(MyVar)

	var isBool bool = true
	
	fmt.Println((isBool))

	var smalPositiveValue uint8

	smalPositiveValue = 10
	fmt.Println(smalPositiveValue)

	var smalPositiveInt int8

	smalPositiveInt = -2
	fmt.Println(smalPositiveInt)

	var myByte byte = 'M'
	//byte is an alias for uint8 
	//byte is used to represent raw data

	fmt.Println(myByte)

	var myRune rune = 'B'
		//rune is an alias for int32
		//rune is used to represent a Unicode code point

	fmt.Println("RUNE IS,", myRune)



	//////////////////////////////
	//FLOAT AND COMPLEX NUMBERS 


	///STRINGS LESSON 
	var FirstName, LastName string 

	FirstName = "Beejay"
	LastName = "Okechukwu"
	//strings are immutable in Go 
	//you cannot change a string after it is created

	Fullname := fmt.Sprintf("My name is  %s, %s\n", FirstName, LastName)
	fmt.Println(Fullname)

	/////VARIABLE DECLARATION
	fmt.Println(MyInt)
	result := fmt.Sprintf("%T", something())
	fmt.Println(result)


	///CONSTANTS
	//Constants are immutable values which are known at compile time and do not change for the life of the program
	//Constants can be character, string, boolean, or numeric values
	//Constants cannot be declared using the := syntax
	//Constants must be assigned a value at the time of declaration
	//The expression on the right side of a constant declaration must be a constant expression
	
	
	var x int = 42 // define x with an example value
	fmt.Println(x)


	res := sum(5, MyInt)
	fmt.Println(res)


	//IDENTIFIERS AND KEYWORDS
	//Identifiers are the names you give to variables, constants, functions, types, and other user-defined items
	//Keywords are the reserved words in Go that have a special meaning and cannot be used as identifiers
	//There are 25 keywords in Go 


	//OPERATORS AND OPERAND
	//Operators are special symbols that perform operations on one or more operands
	// +, - , *, /, %, ++, --, ==, !=, >, <, >=, <=, &&, ||, !, &, |, ^, <<, >>, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=

	//BITWISE OPERATORS
	//Bitwise operators operate on binary representations of integers
	//& (AND), | (OR), ^ (XOR), &^ (AND NOT), << (left shift), >> (right shift)

	var a uint8 = 202
	var b uint8 = 141
	
	c := a | b // OR
	fmt.Println("\n",c)

	// Binary representation
	fmt.Println(strconv.FormatUint(uint64(a), 2)) 
	fmt.Println(strconv.FormatUint(uint64(b), 2)) 
	fmt.Println(strconv.FormatUint(uint64(c), 2)) 


	d := a & b // AND
	fmt.Println("\n",d)

	// Binary representation
	fmt.Println(strconv.FormatUint(uint64(a), 2)) 
	fmt.Println(strconv.FormatUint(uint64(b), 2)) 
	fmt.Println(strconv.FormatUint(uint64(d), 2)) 


	e := a ^ b // XOR
	fmt.Println("\n",e)

	// Binary representation
	fmt.Println(strconv.FormatUint(uint64(a), 2)) 
	fmt.Println(strconv.FormatUint(uint64(b), 2)) 
	fmt.Println(strconv.FormatUint(uint64(e), 2)) 

	f := a &^ b // AND NOT 
	fmt.Println("\n",f)

	// Binary representation
	fmt.Println(strconv.FormatUint(uint64(a), 2)) 
	fmt.Println(strconv.FormatInt(int64(b), 2)) 
	fmt.Println(strconv.FormatUint(uint64(f), 2)) 

	g := ^a // INVERT
	fmt.Println("\n",f)

	// Binary representation
	fmt.Println(strconv.FormatUint(uint64(a), 2)) 
	fmt.Println(strconv.FormatUint(uint64(g), 2)) 

	//SHIFT OPERATORS
 	var shift uint8 = 20
	shift = shift >> 2

	fmt.Println("\n",shift)
	fmt.Println(strconv.FormatUint(uint64(shift), 2))

	shift = shift << 2
	fmt.Println("\n",shift)
	fmt.Println(strconv.FormatUint(uint64(shift), 2))

	//BIT MASKING 
	//I NEED TO MAKE MORE RESEARCH ON THIS



	//ARRAYS

	//ARR DECLARATION BASIC FORMAT
	var arr1 [5]int
	fmt.Println(arr1)

	//ARR DECLARATION WITH VALUES/ELEMENTS
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr[2])

	//SPARSE ARRAYS DECLARATION
	arr2 := [5]int{1:10, 3:30}
	fmt.Println(arr2)

	//IMPLICIT LENGTH ARRAYS
	arr3 := [...]int{1,2,3,4,5,6}
	fmt.Println(len(arr3))
	

	//MULTI-DIMENSIONAL ARRAYS
	arr4 := [2][3]int{{3,4},{3,6,8}}
	arr4[0][2] = 9
	fmt.Println(arr4)


	//SLICES

	slice := make([]int, 5, 20)
	fmt.Println((slice))
	slc := make([]int, 5) 
	//arry := [6]int{1,2,3,4,5,6}
	slc = append(slc, 4,6,6)
	slcs := copy(slice, slc)
	fmt.Println((slc))
	fmt.Println("Destination slice:", slcs)

	name := "Code & strings"

	//STRING TO BYTE/RUNE SLICE
	var byteName []rune = []rune(name)
	fmt.Println("Byte slice:", byteName[0])
	
	//MAPS
	//Maps are unordered collections of key-value pairs
	//Maps are also known as associative arrays or dictionaries in other programming languages
	//Maps are reference types and must be initialized before use
	
	//MAP DECLARATION
	var mymap = map[string]int{
		"lastname": 20,
	}

	mymap["name"] = 12
	mymap["age"] = 30

	foo, ok := mymap["aha"]
	if !ok {
		fmt.Println("Not found")
	}

	mymap["name"] = 40
	fmt.Println(foo)
	fmt.Println(len(mymap))

	//STRUCTS
	type person struct {
		name string
		age int
		address string
		language []string
	}

	//STRUCTS
	person1 := person {
		name: "beejay",
		age: 4,
		address: "5 daing ",
		language: []string{"english", "math"},
	}
	fmt.Println(person1)
	fmt.Println(person1.name)

	person1.age = 12
	fmt.Println(person1.age)

	//ANONYMOUS STRUCTS
	g1 := struct{
		id int
		country string
		} {
		id: 1,
		country: "Nigeria",
	}

	fmt.Println(g1)

	ageNum := 11;

	//IF-ELSE STATEMENT
	if ageNum < 12 {
		fmt.Println("Child")
	} else if ageNum >= 12 && ageNum < 20 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Adult")
	}

	aLoop := 10;

	//FOR LOOP
	for aLoop <= 0 {
		fmt.Println("this is the number: ", aLoop)
		aLoop++
	}

	for i := 0; i <= aLoop; i++ {
		fmt.Println(i)
	}


	//CLASSIC FOR LOOP
	Loops := []string{"Beejay", "Okechukwu", "Code"}
	for i := 0; i < len(Loops); i++ {
		fmt.Println("This is the names: ", Loops[i])
	}

	type User struct {
		Username string
		Email string
		Age int
	}

	users := User{
		Username: "Beejay",
		Email: "A@a.com",
		Age: 24,
	}

	fmt.Println(users)

	//mp := map[string] string {"hello": "bello"}
	rangea := []int{1,2,3,4,5,6}

	for i, v := range rangea {
		fmt.Printf("the index is %v and the value is %d\n" , i, v)
	}

	//SWITCH STATEMENT 
	days := "Sunday"

	switch days {
	case "Monday":
		fmt.Println("Hello this is ," + days)
	case "Tuesday": 
		fmt.Println("Hello this is ," + days)
	default: 
		fmt.Println("Hello this is ," + days)
	}


	outerloop: 
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == 1 && j == 1 {
					continue outerloop
				}
				fmt.Println(i, j)
			}
		}

	Looping := 5  

	if Looping == 10 {
		goto end
	}
	fmt.Println(Looping) 

end: {
	fmt.Println("This is the end of the program")
}

word := "Testings"

switch wordlen := len(word); wordlen {
case 4: 
	fmt.Println("wordlen is 4")
case 7: 
	fmt.Println("wordlen is 7")
default: {
	fmt.Println("wordlen is default")
	}
}

  v := Variadic(1,2,3,4,5,6,7,8)
  fmt.Println(v)

  fmt.Println(multiplyByNum(multiply, 5))

  var intPtr *int

  ag := 10
  
  intPtr = &ag;
  fmt.Println(intPtr)
}

	

//FUNCTION 
func something() string {
	//this function returns the string representation of MyInt
	return fmt.Sprintf("%d", MyInt) 
}

func sum(a, b int) int {
	return a + b
}

//VARIADIC PARAMETER
func Variadic (num ...int) int {
	total := 0

	for _, v := range num {
		total += v
	}
	return total
}

//FUNCTION THAT RETURNS ANOTHER FUNCTION
func multiply (x int) int {
	return x * 5
}

func multiplyByNum (ml func (int) int, x int) int {
	return ml(x)
}


