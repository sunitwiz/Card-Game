/*
//create a repo in github and take the url
//git init
//git remote add origin https://github.com/sunitwiz/Card-Game.git
//git add .
//git commit -m "first commit"
//git push -u origin master






//commands in Go
go build "hello world.go"   //compiles the code
go run "hello world.go"    //runs the code
go fmt "hello world.go"   //formats the code syntax
go install "hello world.go"  //installs the code
go get "hello world.go"   //downloads the code
go test "hello world.go"   //tests the code
go test   //to test


incase of multiple go files we need to create go.mod file

✅ 1. Initialize a Go Module
Run the following command inside your project directory 

1.go mod init cardgame

This will create a go.mod file with the module name cardgame.

✅ 2. Verify Your go.mod File
After running go mod init cardgame, check the go.mod file:

cat go.mod
expected o/p
module cardgame
go 1.XX  # (Your Go version, e.g., go 1.21)


✅ 3. Run Your Go Program
Now, you should be able to run:

go run .



4. setting.json 
"code-runner.executorMap": {
        "go": "cd $dir && ([ -f go.mod ] && go run . || go run $fileName)"
    }

*/


Classes	❌ No classes, uses structs instead
Inheritance	❌ No classical inheritance, uses composition
 Polymorphism	✅ Achieved via interfaces
 Encapsulation	✅ Controlled using capitalization (exported vs unexported fields)
 Abstraction	✅ Supported via interfaces
Method Overloading	❌ Not supported (methods must have unique names)
 Operator Overloading	❌ Not supported


package main           //this defines it will create a executable fie

import "fmt"           //importing the fmt package
func main(){        //same line as main the "{"" should be on the same line as main}
	fmt.Println("Hello, World!")
}

1.
type of packages in Go
1. Executable package      //(main package  )genertes a file that we can run
                           // need to write package main  in the code
2. Reusable package        //other than main package code used as a helper code





2 .variable declaration

var card string = "Ace of Spades"  //var is a keyword to declare a variable
                                      //card is the name of the variable    
                                        //string is the type of the variable
                                        //Ace of Spades is the value of the variable
//or
card := "Ace of Spades"  //Go will automatically detect the type of the variable
//or                    //this type of declaration should be always declared inside a function
card = "Ace of Spades"  //if the variable is already declared then we can assign the value to it
//or
var card string  //if we declare a variable without assigning a value then it will be assigned a default value
fmt.Println(card)  //default value is an empty string




3. function statements  

func processArray(num int, arr []int) []int {     
    // Example: Append num to the array and return it
    // we can return multiple values in Golang
    arr = append(arr, num)
    return arr
}


4. Array and slice and map
var arr [5]int // An array of 5 integers (all initialized to 0)
arr[0] = 10    // Assigning a value
fmt.Println(arr) // Output: [10 0 0 0 0]

var slice []int      // A nil slice (no underlying array yet)
slice = append(slice,  20, 30) 
fmt.Println(slice)   // Output: [10 20 30]


slice := []int{1, 2, 3}   //I will define
var slice []int         //I will declare for empty


myMap := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range myMap {                     //every variable declaredneed to be used , if not used it will throw an error
    fmt.Println("Key:", key, "Value:", value)        //we can write _,value if we dont want to use the key
}




type Car struct {
    Brand string
    Year  int
}

cars := Car{} // ✅ Valid: Creates an empty struct
numbers := []int{}    // ✅ Valid: Creates an empty slice
scores := [3]int{}    // ✅ Valid: Creates an array with default values (0,0,0)
grades := map[string]int{} // ✅ Valid: Creates an empty map
type deck []string
cards := deck{} // ✅ Valid: Creates an empty slice of type deck


cars := int{} // ❌ Invalid: You cannot initialize a primitive type like this




5. loops
nums := []int{10, 20, 30, 40}
for index, value := range nums {
    fmt.Println( index, value)
}



for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}



6. if else
num := 10
if num > 5 {
    fmt.Println("Number is greater than 5")
} else {
    fmt.Println("Number is less than or equal to 5")
}



7. Type method

Feature--->	var deck []string	        type deck []string
What it is--->	A slice variable	    A custom type based on a slice
Can have methods?--->	❌ No	      ✅ Yes
Usage--->	Directly a []string	        Needs conversion (deck{"Ace", "King"})
Flexibility--->	Just a slice	        Can add behavior (methods)


Use type deck []string if you want to extend functionality (e.g., add methods like shuffle(), drawCard(), etc.)


type deck []string  // deck is now a custom type

func (d deck) print() { // Adding a method to deck
    for _, card := range d {
        fmt.Println(card)
    }
}
func main() {
    cards := deck{"Ace", "King", "Queen"} // Using the custom type
    cards.print() // Output: Ace, King, Queen
}



8. Receiver function
receiver functions (methods) allow you to define functions on custom types (like structs or custom slices). This makes Go’s approach feel object-oriented without using traditional OOP classes.
// ANy variable of receiverType has access to print method
func (ReceiverName receiverType ) methodName(params) returnType {
    // Function logic
}




// Define a custom type deck (slice of strings)
type deck []string

// Method with a receiver (prints all cards)

// ANy variable of type deck has access to print method
func (d deck) print() {
    for _, card := range d {
        fmt.Println(card)
    }
}

func main() {
    cards := deck{"Ace", "King", "Queen"}
    cards.print()  // Calls the method
}





9.os package 
to read and file 


10. Test cases
func TestNewDeck(t *testing.T) {   //testing.T is a struct provided by Go’s built-in testing package.
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(b))    //%v tells Go to print the value of the variable in its default format.
	}

}  


t.Log(msg)	Logs a message (useful for debugging).
t.Errorf(format, args...)	Logs an error but continues execution.
t.Fatalf(format, args...)	Logs an error and stops execution.
t.Fail()	Marks test as failed but continues execution.
t.FailNow()	Marks test as failed and stops execution.
t.Skip(msg)	Skips the test with a message.

How *testing.T is Used in Tests
In Go, test functions:

Must be in a _test.go file.
Must start with Test and take t *testing.T as a parameter.


 


