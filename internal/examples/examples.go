package examples

// Basic type examples

// Integer examples
const IntExamples = `
// Integer type declarations
var a int = 42                  // Platform dependent (32 or 64 bit)
var b int8 = 127                // -128 to 127
var c int16 = 32767             // -32768 to 32767 
var d int32 = 2147483647        // -2147483648 to 2147483647
var e int64 = 9223372036854775807
var f uint = 42                 // Unsigned, platform dependent
var g uint8 = 255               // 0 to 255
var h uint16 = 65535            // 0 to 65535
var i uint32 = 4294967295       // 0 to 4294967295
var j uint64 = 18446744073709551615

// Short declaration (type inferred, defaults to int)
k := 42
`

// Integer operations
const IntOperations = `
// Basic arithmetic
a := 10
b := 3

sum := a + b         // 13
difference := a - b  // 7
product := a * b     // 30
quotient := a / b    // 3 (integer division, truncated)
remainder := a % b   // 1 (modulus)

// Increment and decrement
a++                  // a is now 11
b--                  // b is now 2

// Bitwise operations
c := 5               // 101 in binary
d := 3               // 011 in binary

andResult := c & d   // 001 = 1 (bitwise AND)
orResult := c | d    // 111 = 7 (bitwise OR)
xorResult := c ^ d   // 110 = 6 (bitwise XOR)
notResult := ^c      // -6 (bitwise NOT, flips all bits)

// Bit shifting
leftShift := c << 1  // 1010 = 10 (shift left by 1)
rightShift := c >> 1 // 010 = 2 (shift right by 1)
`

// Float examples
const FloatExamples = `
// Float declarations
var a float32 = 3.14
var b float64 = 3.141592653589793

// Short declaration (defaults to float64)
c := 3.14       // float64
d := float32(3.14)  // explicitly float32

// Scientific notation
e := 1.234e5    // 123400
f := 1.234e-2   // 0.01234

// Special values
var posInf = 1.0 / 0.0       // +Inf (positive infinity)
var negInf = -1.0 / 0.0      // -Inf (negative infinity)
var nan = 0.0 / 0.0          // NaN (not a number)
`

// Float operations
const FloatOperations = `
// Basic operations
a := 10.5
b := 3.2

sum := a + b         // 13.7
difference := a - b  // 7.3
product := a * b     // 33.6
quotient := a / b    // 3.28125

// Math functions (from math package)
import "math"

squareRoot := math.Sqrt(a)   // 3.24
power := math.Pow(a, 2)      // 110.25
ceiling := math.Ceil(a)      // 11.0
floor := math.Floor(a)       // 10.0
rounded := math.Round(a)     // 11.0

// Check for special values
isInf := math.IsInf(1.0/0.0, 1)  // true
isNaN := math.IsNaN(0.0/0.0)     // true
`

// Float precision example
const FloatPrecision = `
// Precision issues
a := 0.1
b := 0.2
c := a + b

fmt.Println(c)                   // 0.30000000000000004
fmt.Println(c == 0.3)            // false (due to precision)

// Better comparison for floating point
const epsilon = 1e-9
fmt.Println(math.Abs(c-0.3) < epsilon)  // true

// For displaying, use formatting
fmt.Printf("%.1f\n", c)          // 0.3
`

// Bool examples
const BoolExamples = `
// Boolean declarations
var a bool = true
var b bool = false
var c bool       // Zero value is false

// Short declaration
d := true

// From comparisons
e := 5 > 3       // true
f := 10 == 9     // false
g := 10 != 9     // true
h := 7 <= 7      // true
`

// Bool operations
const BoolOperations = `
// Logical operators
a := true
b := false

andResult := a && b    // false (logical AND)
orResult := a || b     // true (logical OR)
notA := !a             // false (logical NOT)
notB := !b             // true

// Short-circuit evaluation
x := 10
result := x > 5 && someExpensiveFunction()  // If x <= 5, someExpensiveFunction() won't be called

// In conditionals
if a && !b {
    fmt.Println("a is true and b is false")
}

// As flags
isEnabled := true
if isEnabled {
    // Feature is enabled
}
`

// String examples
const StringExamples = `
// String declarations
var a string = "Hello, Go!"
var b string        // Zero value is "" (empty string)

// Short declaration
c := "Learning Go strings"

// Raw string literals (preserves formatting)
multiline := ` + "`" + `This is a 
multi-line string
with "quotes" preserved` + "`" + `

// String concatenation
greeting := "Hello" + ", " + "world!"  // "Hello, world!"
`

// String operations
const StringOperations = `
// String operations
s := "Hello, Go!"

// Length (in bytes, not characters)
length := len(s)            // 10

// Accessing individual bytes (not characters)
firstByte := s[0]           // 'H' (72 in ASCII)

// Substring (slicing)
substring := s[0:5]         // "Hello"
prefix := s[:5]             // "Hello"
suffix := s[7:]             // "Go!"

// String comparison
equals := s == "Hello, Go!"  // true
notEquals := s != "hello"    // true
lessThan := "a" < "b"        // true (lexicographical comparison)

// Using strings package
import "strings"

contains := strings.Contains(s, "Go")      // true
hasPrefix := strings.HasPrefix(s, "Hello") // true
hasSuffix := strings.HasSuffix(s, "!")     // true
index := strings.Index(s, "Go")            // 7
count := strings.Count(s, "l")             // 2
replaced := strings.Replace(s, "Hello", "Hi", 1)  // "Hi, Go!"
upper := strings.ToUpper(s)                // "HELLO, GO!"
lower := strings.ToLower(s)                // "hello, go!"
trimmed := strings.TrimSpace(" Hello ")    // "Hello"
split := strings.Split("a,b,c", ",")       // ["a", "b", "c"]
joined := strings.Join([]string{"a", "b"}, "-") // "a-b"
`

// String runes
const StringRunes = `
// Working with runes (Unicode characters)
s := "Hello, 世界"  // String with non-ASCII characters

// Ranging over characters (as runes)
for i, r := range s {
    fmt.Printf("%d: %c (%d)\n", i, r, r)
}
// Output:
// 0: H (72)
// 1: e (101)
// ...
// 7: 世 (19990)  // Note the index jumps because this is a 3-byte character
// 10: 界 (30028)

// Converting between strings and runes
runes := []rune(s)
runeCount := len(runes)     // 9 (actual character count)
backToString := string(runes)

// Getting a specific character safely
thirdRune := []rune(s)[2]   // 'l'

// UTF-8 Package
import "unicode/utf8"

runeCount = utf8.RuneCountInString(s)  // 9
`

// Byte examples
const ByteExamples = `
// Byte declarations (byte is alias for uint8)
var a byte = 65            // 'A' in ASCII
var b byte = 'A'           // Same as above
var c []byte = []byte("Hello")  // Byte slice from string

// Zero value
var d byte                 // 0

// Numerical operations (same as uint8)
e := a + 1                 // 66 ('B')
f := a * 2                 // 130
`

// Byte operations
const ByteOperations = `
// Working with byte slices
bs := []byte("Hello, Go!")

// Modify individual bytes
bs[0] = 'J'                // Changes "Hello" to "Jello"

// Convert back to string
s := string(bs)            // "Jello, Go!"

// Using bytes package
import "bytes"

joined := bytes.Join([][]byte{[]byte("Hello"), []byte("Go")}, []byte(", "))
contains := bytes.Contains(bs, []byte("Go"))   // true
index := bytes.Index(bs, []byte("Go"))         // 7
count := bytes.Count(bs, []byte("o"))          // 2
replaced := bytes.Replace(bs, []byte("Go"), []byte("World"), 1)

// Reading/writing files
import "os"
import "io/ioutil"

data, err := ioutil.ReadFile("file.txt")  // Returns []byte
err = ioutil.WriteFile("file.txt", []byte("Hello"), 0644)
`

// Rune examples
const RuneExamples = `
// Rune declarations (rune is alias for int32)
var a rune = 65            // 'A' in Unicode
var b rune = 'A'           // Same as above
var c rune = '世'           // Unicode character (CJK ideograph)

// Zero value
var d rune                 // 0

// Creating rune from integer
e := rune(65)              // 'A'
`

// Rune operations
const RuneOperations = `
// Working with runes
r1 := '世'
r2 := '界'

// Numerical operations (same as int32)
sum := r1 + 1              // 19991 (next Unicode code point)

// String of one rune
s1 := string(r1)           // "世"

// Character classification with unicode package
import "unicode"

isLetter := unicode.IsLetter(r1)    // true
isDigit := unicode.IsDigit('9')     // true
isSpace := unicode.IsSpace(' ')     // true
isUpper := unicode.IsUpper('A')     // true
isLower := unicode.IsLower('a')     // true
toUpper := unicode.ToUpper('a')     // 'A'
toLower := unicode.ToLower('A')     // 'a'

// Iterating through string as runes
for _, r := range "Hello, 世界" {
    fmt.Printf("%c ", r)
}
// Output: H e l l o ,   世 界
`

// Complex examples
const ComplexExamples = `
// Complex number declarations
var a complex64 = 1 + 2i
var b complex128 = 1.5 + 3.1i

// Short declaration (defaults to complex128)
c := 2 + 3i             // complex128
d := complex(2, 3)      // complex128, same as 2+3i
e := complex64(1 + 2i)  // explicitly complex64
`

// Complex operations
const ComplexOperations = `
// Basic operations
a := 2 + 3i
b := 1 + 2i

sum := a + b           // 3+5i
difference := a - b    // 1+1i
product := a * b       // -4+7i  (2+3i)*(1+2i) = 2+4i+3i-6 = -4+7i
quotient := a / b      // 1.6+0.2i

// Extract real and imaginary parts
real := real(a)        // 2.0
imag := imag(a)        // 3.0

// Complex math functions
import "math/cmplx"

abs := cmplx.Abs(a)    // 3.6055 (magnitude)
phase := cmplx.Phase(a) // 0.9828 (phase angle in radians)
sqrt := cmplx.Sqrt(-1)  // 0+1i (square root of negative number)
exp := cmplx.Exp(1i * math.Pi) // -1+0i (Euler's formula)
log := cmplx.Log(1i)   // 0+1.5708i
pow := cmplx.Pow(a, b) // -0.015-0.179i
`

// Array examples
const ArrayExamples = `
// Array declarations with explicit size
var a [5]int                          // Array of 5 integers, initialized to zero values
var b [3]string = [3]string{"a", "b", "c"}
var c [2]bool = [2]bool{true, false}

// Short declarations with initialization
d := [5]int{1, 2, 3, 4, 5}
e := [3]float64{1.1, 2.2, 3.3}

// Using ... to let compiler determine size
f := [...]int{1, 2, 3, 4}             // Length is 4

// Sparse arrays (with specific positions initialized)
g := [5]int{1: 10, 3: 30}             // [0, 10, 0, 30, 0]

// Multi-dimensional arrays
h := [2][3]int{{1, 2, 3}, {4, 5, 6}}
`

// Array operations
const ArrayOperations = `
// Accessing elements (zero-indexed)
arr := [5]int{10, 20, 30, 40, 50}
first := arr[0]       // 10
third := arr[2]       // 30

// Modifying elements
arr[1] = 25           // [10, 25, 30, 40, 50]

// Getting array length
length := len(arr)    // 5

// Iterating over arrays
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

// Using range
for index, value := range arr {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Comparing arrays (arrays of same type and length can be compared)
arr1 := [3]int{1, 2, 3}
arr2 := [3]int{1, 2, 3}
arr3 := [3]int{3, 2, 1}

fmt.Println(arr1 == arr2)  // true
fmt.Println(arr1 == arr3)  // false

// Copying arrays (creates a new copy)
arr4 := arr1
arr4[0] = 99      // Only modifies arr4, not arr1
`

// Slice examples
const SliceExamples = `
// Slice declarations
var a []int                  // nil slice, has no backing array yet
var b []string = []string{"a", "b", "c"}
var c []bool = []bool{true, false, true}

// Short declarations with initialization
d := []int{1, 2, 3, 4, 5}    // Slice literal
e := []float64{1.1, 2.2, 3.3}

// Creating slices with make (allocates zeroed array and returns slice)
f := make([]int, 5)          // Slice of length 5, capacity 5
g := make([]int, 3, 5)       // Slice of length 3, capacity 5

// Creating slices from arrays
arr := [5]int{10, 20, 30, 40, 50}
h := arr[1:4]                // [20, 30, 40] (elements 1 through 3)

// Creating slices from other slices
i := d[2:4]                  // [3, 4] (elements 2 through 3 of d)

// Empty slices (non-nil but length 0)
j := []int{}
k := make([]int, 0)
`

// Slice operations
const SliceOperations = `
// Creating and accessing slices
s := []int{10, 20, 30, 40, 50}

// Accessing elements (like arrays)
first := s[0]                // 10
third := s[2]                // 30

// Getting slice length and capacity
length := len(s)             // 5
capacity := cap(s)           // 5 (can be larger than length)

// Modifying elements
s[1] = 25                    // [10, 25, 30, 40, 50]

// Check if slice is nil
isNil := s == nil            // false
`

// Slice manipulation
const SliceManipulation = `
// Slicing operations
s := []int{10, 20, 30, 40, 50}

// Slice expression syntax: slice[low:high]
// - low is inclusive, high is exclusive
// - defaults: low=0, high=len(slice)
s1 := s[1:4]                 // [20, 30, 40]
s2 := s[:3]                  // [10, 20, 30]
s3 := s[2:]                  // [30, 40, 50]
s4 := s[:]                   // [10, 20, 30, 40, 50] (full slice)

// Appending elements (key slice operation)
s = append(s, 60)            // [10, 20, 30, 40, 50, 60]
s = append(s, 70, 80)        // [10, 20, 30, 40, 50, 60, 70, 80]

// Appending one slice to another
t := []int{90, 100}
s = append(s, t...)          // [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]

// Copying slices
dest := make([]int, len(s))
copied := copy(dest, s)      // copied = number of elements copied
                             // dest now contains same elements as s

// Deleting elements (no built-in delete function)
// Option 1: Using append to skip element at index i
i := 2
s = append(s[:i], s[i+1:]...) // Removes element at index 2

// Option 2: Using copy to avoid memory leaks
s = s[:i+copy(s[i:], s[i+1:])]

// Clearing a slice
s = s[:0]                    // Length becomes 0, capacity unchanged
`

// Map examples
const MapExamples = `
// Map declarations
var a map[string]int                      // nil map (can't be used yet)
var b map[string]bool = map[string]bool{  // Initialized map
    "A": true,
    "B": false,
}

// Short declarations with initialization
c := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}

// Creating maps with make
d := make(map[string]float64)             // Empty map, ready to use
`

// Map operations
const MapOperations = `
// Working with maps
m := map[string]int{
    "apple":  5,
    "banana": 8,
    "orange": 7,
}

// Accessing values
count := m["apple"]            // 5

// Adding or updating values
m["pear"] = 3                  // Add new key-value
m["apple"] = 6                 // Update existing value

// Deleting keys
delete(m, "banana")            // Removes "banana" entry

// Getting map length
size := len(m)                 // 2 (after deletion)
`

// Map key checking
const MapKeyCheck = `
// Checking if a key exists
m := map[string]int{"a": 1, "b": 2}

// The "comma ok" idiom
value, exists := m["a"]        // value=1, exists=true
v, ok := m["z"]                // v=0 (zero value), ok=false

// Using it in conditionals
if val, exists := m["c"]; exists {
    fmt.Println("Value:", val)
} else {
    fmt.Println("Key 'c' not found")
}

// Safe value access with default
func getWithDefault(m map[string]int, key string, defaultVal int) int {
    if val, ok := m[key]; ok {
        return val
    }
    return defaultVal
}

result := getWithDefault(m, "z", -1)  // Returns -1 because "z" doesn't exist
`

// Map iteration
const MapIteration = `
// Iterating over maps
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

// Using range to iterate (order is not guaranteed)
for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}

// Iterating over just keys
for key := range m {
    fmt.Println("Key:", key)
}

// Iterating in a specific order
import "sort"

// Get all keys
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}

// Sort keys
sort.Strings(keys)

// Iterate in sorted order
for _, k := range keys {
    fmt.Printf("Key: %s, Value: %d\n", k, m[k])
}
`

// Struct examples
const StructExamples = `
// Basic struct definition
type Person struct {
    FirstName string
    LastName  string
    Age       int
    Email     string
}

// Struct declarations
var a Person                           // Zero-valued struct
var b Person = Person{                 // Initialized struct
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
    Email:     "john@example.com",
}

// Short declarations with field names
c := Person{
    FirstName: "Alice",
    LastName:  "Smith",
    Age:       25,
    Email:     "alice@example.com",
}

// Short declarations without field names (order matters)
d := Person{"Bob", "Johnson", 35, "bob@example.com"}

// Nested structs
type Address struct {
    Street  string
    City    string
    ZipCode string
}

type Employee struct {
    Person  Person
    Address Address
    Salary  float64
    Title   string
}

// Struct with tags (used for serialization and reflection)
type Product struct {
    ID    int    // has json:"id" tag
    Name  string // has json:"name" tag
    Price int    // has json:"price,omitempty" tag
}
`

// Struct operations
const StructOperations = `
// Working with structs
p := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

// Accessing fields
name := p.FirstName            // "John"
age := p.Age                   // 30

// Modifying fields
p.Age = 31
p.Email = "john.doe@example.com"

// Struct comparison
p1 := Person{"John", "Doe", 30, ""}
p2 := Person{"John", "Doe", 30, ""}
p3 := Person{"Jane", "Doe", 28, ""}

equal := p1 == p2              // true (all fields are equal)
notEqual := p1 == p3           // false (different values)

// Structs as function parameters (passed by value)
func birthday(p Person) Person {
    p.Age++
    return p
}

older := birthday(p)           // Creates a new struct with age incremented
// p.Age is still the original value
`

// Struct methods
const StructMethods = `
// Methods with value receivers
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Using methods
rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()                // 50

// Value receivers don't modify the original
rect.Area()                        // rect is unchanged

// Pointer receivers modify the original
rect.Scale(2)                      // rect is now {Width: 20, Height: 10}
`

// Struct embedding
const StructEmbedding = `
// Embedding for composition
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

// Embedding Person in Employee
type Employee struct {
    Person                // Embedded struct
    Company string
    Salary  float64
}

// Usage of embedded fields
e := Employee{
    Person:  Person{Name: "John", Age: 30},
    Company: "Acme Inc",
    Salary:  50000,
}

// Fields and methods are "promoted" to the containing struct
name := e.Name            // Access embedded field directly
greeting := e.Greet()     // Call embedded method directly

// Override behavior
func (e Employee) Greet() string {
    return "Hello, I'm " + e.Name + " and I work at " + e.Company
}

// Now e.Greet() calls the Employee method
`

// Pointer examples
const PointerExamples = `
// Pointer declarations
var a *int                // nil pointer to int
var b *string             // nil pointer to string

// Creating pointers with new
c := new(int)             // pointer to zero value int
*c = 42                   // set the value at the pointer

// Creating pointers to existing variables
x := 10
p := &x                   // p points to x
*p = 20                   // x is now 20

// Zero value of pointers is nil
var q *int
isNil := q == nil         // true

// Checking for nil before dereferencing
if q != nil {
    value := *q           // Safe, no panic
}
`

// Struct pointers
const StructPointers = `
// Creating struct pointers
type Person struct {
    Name string
    Age  int
}

// Method 1: Using &
p1 := &Person{Name: "Alice", Age: 30}

// Method 2: Using new
p2 := new(Person)
p2.Name = "Bob"           // Syntactic sugar for (*p2).Name
p2.Age = 25               // Same as (*p2).Age

// Field access with pointers
fmt.Println(p1.Name)      // Alice (Go allows p1.Name instead of (*p1).Name)
fmt.Println((*p1).Age)    // 30 (explicit dereferencing)

// Modifying struct through pointer
p1.Age = 31               // Changes the original struct
`

// Pointer patterns
const PointerPatterns = `
// Common pointer use case: modifying function parameters
func increment(val *int) {
    *val++                // Increments the value at the pointer
}

x := 10
increment(&x)             // x is now 11

// Avoid unnecessary copying of large structs
type LargeStruct struct {
    Data [1024]int
    // ... many fields
}

// More efficient with pointer
func process(s *LargeStruct) {
    s.Data[0] = 42        // No copying of the large struct
}

// Implement data structures like linked lists
type Node struct {
    Value int
    Next  *Node           // Pointer to next node
}

// Create a linked list
head := &Node{Value: 1}
head.Next = &Node{Value: 2}
head.Next.Next = &Node{Value: 3}
`

// Function examples
const FunctionExamples = `
// Basic function declaration
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return                // Returns x and y implicitly
}

// Variadic functions
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

total := sum(1, 2, 3, 4)  // 10
`

// Function values
const FunctionValues = `
// Functions as values
add := func(a, b int) int {
    return a + b
}
result := add(3, 4)       // 7

// Functions as types
type MathFunc func(int, int) int

var operations map[string]MathFunc = map[string]MathFunc{
    "add":      func(a, b int) int { return a + b },
    "subtract": func(a, b int) int { return a - b },
    "multiply": func(a, b int) int { return a * b },
    "divide":   func(a, b int) int { return a / b },
}

result := operations["add"](5, 3)  // 8
`

// Higher order functions
const HigherOrderFunctions = `
// Function that returns another function
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

add5 := makeAdder(5)
result := add5(3)         // 8

// Function that takes a function as an argument
func applyTwice(f func(int) int, x int) int {
    return f(f(x))
}

double := func(x int) int {
    return x * 2
}

result := applyTwice(double, 3)  // double(double(3)) = double(6) = 12

// Transforming a list with a function
func transform(numbers []int, f func(int) int) []int {
    result := make([]int, len(numbers))
    for i, v := range numbers {
        result[i] = f(v)
    }
    return result
}

nums := []int{1, 2, 3, 4}
squared := transform(nums, func(x int) int {
    return x * x
})  // [1, 4, 9, 16]
`

// Closures
const Closures = `
// Creating a counter with closure
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

counter := makeCounter()
fmt.Println(counter())    // 1
fmt.Println(counter())    // 2
fmt.Println(counter())    // 3

// Another counter is independent
counter2 := makeCounter()
fmt.Println(counter2())   // 1

// Capturing loop variables
funcs := make([]func(), 3)

// Wrong way (all funcs will use last value of i)
for i := 0; i < 3; i++ {
    funcs[i] = func() { fmt.Println(i) }
}
// All will print 3

// Correct way (creating a new i for each iteration)
for i := 0; i < 3; i++ {
    i := i  // Create new variable in inner scope
    funcs[i] = func() { fmt.Println(i) }
}
// Will print 0, 1, 2
`

// Channel basics
const ChannelBasics = `
// Channel declarations
var a chan int            // nil channel (unusable until initialized)
b := make(chan string)    // Create a channel of strings

// Sending and receiving
go func() {
    b <- "hello"          // Send "hello" to channel b
}()

msg := <-b                // Receive from channel b, blocks until data is available
fmt.Println(msg)          // "hello"

// Closing channels
close(b)                  // Signals that no more values will be sent

// Check if channel is closed
val, ok := <-b            // ok is false if channel is closed and empty
if !ok {
    fmt.Println("Channel closed")
}

// Range over a channel
c := make(chan int)
go func() {
    for i := 1; i <= 5; i++ {
        c <- i
    }
    close(c)              // Must close for the range to complete
}()

for num := range c {      // Iterates until channel is closed
    fmt.Println(num)
}
`

// Buffered channels
const BufferedChannels = `
// Creating buffered channels
c := make(chan int, 3)    // Buffer size 3

// Send without blocking (until buffer is full)
c <- 1                    // These don't block because buffer isn't full
c <- 2
c <- 3

// Next send would block until space is available
// c <- 4                 // This would block until something is received

// Receive values
first := <-c              // 1
second := <-c             // 2

// Now buffer has space again
c <- 4                    // Doesn't block

// Check buffer status
fmt.Println(len(c))       // 2 (number of items in buffer)
fmt.Println(cap(c))       // 3 (buffer capacity)
`

// Channel directions
const ChannelDirections = `
// Channel direction constraints
func send(ch chan<- int, value int) {
    ch <- value           // Can only send to this channel
}

func receive(ch <-chan int) int {
    return <-ch           // Can only receive from this channel
}

func pipe(in <-chan int, out chan<- int) {
    value := <-in
    out <- value
}

// Bidirectional channel can be passed to constrained functions
ch := make(chan int)
go send(ch, 42)           // ch as send-only
value := receive(ch)      // ch as receive-only

// But not vice versa
sendCh := make(chan<- int)
// value := receive(sendCh)  // Compile error
`

// Channel select
const ChannelSelect = `
// Select statement for multiple channels
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    time.Sleep(1 * time.Second)
    ch1 <- "message from ch1"
}()

go func() {
    time.Sleep(2 * time.Second)
    ch2 <- "message from ch2"
}()

// Select first channel that's ready
select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
}

// Non-blocking receive with default
select {
case msg := <-ch1:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No message available")
}

// Timeout pattern
select {
case msg := <-ch1:
    fmt.Println("Received:", msg)
case <-time.After(500 * time.Millisecond):
    fmt.Println("Timeout")
}
`

// Interface definition
const InterfaceDefinitionExample = `
// Interface definition
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Composing interfaces
type ReadWriter interface {
    Reader  // Embedding interfaces
    Writer
}

// Empty interface (matches any type)
type Any interface{}
`

// Interface implementation
const InterfaceImplementation = `
// Implementing an interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
    Width, Height float64
}

// Methods that satisfy the Shape interface
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}

// Circle also implements Shape
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Function that uses the interface
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Usage
rect := Rectangle{Width: 5, Height: 10}
circle := Circle{Radius: 7}

PrintShapeInfo(rect)      // Works for Rectangle
PrintShapeInfo(circle)    // Works for Circle
`

// Interface values
const InterfaceValues = `
// Interface values
var s Shape
s = Rectangle{Width: 5, Height: 10}
fmt.Println(s.Area())     // 50

s = Circle{Radius: 7}
fmt.Println(s.Area())     // ~153.94

// Interface values have both a type and a value
// nil interface values
var w Writer
// w.Write([]byte("hello"))  // Panic: nil pointer dereference

// Interface with nil concrete value
type NullWriter struct{}
func (NullWriter) Write(p []byte) (int, error) {
    return len(p), nil
}

var nw *NullWriter       // nil pointer
w = nw                   // Interface value with nil concrete value
len, _ := w.Write([]byte("hello"))  // No panic, calls the method with nil receiver
`

// Empty interface
const EmptyInterface = `
// Using the empty interface
func PrintAny(v interface{}) {
    fmt.Println(v)
}

PrintAny(42)            // Works with int
PrintAny("hello")       // Works with string
PrintAny(struct{}{})    // Works with any struct

// Storing mixed types in a slice
mixedSlice := []interface{}{
    42,
    "hello",
    true,
    []int{1, 2, 3},
}

// Map with any type of value
data := make(map[string]interface{})
data["name"] = "Alice"
data["age"] = 30
data["married"] = false
data["scores"] = []int{95, 89, 92}
`

// Type assertions
const TypeAssertions = `
// Type assertions
var i interface{} = "hello"

// Type assertion to access the underlying string
s, ok := i.(string)
if ok {
    fmt.Println(s)        // "hello"
}

// Type assertion without check (will panic if wrong type)
s = i.(string)            // Safe because i is actually a string
// n := i.(int)           // Would panic

// Type switch
func describe(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Printf("String of length %d: %s\n", len(v), v)
    case int:
        fmt.Printf("Integer: %d\n", v)
    case bool:
        fmt.Printf("Boolean: %v\n", v)
    case []interface{}:
        fmt.Printf("Slice of interface{} with %d elements\n", len(v))
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

describe("hello")         // String of length 5: hello
describe(42)              // Integer: 42
describe(true)            // Boolean: true
describe([]interface{}{1, 2, 3})  // Slice of interface{} with 3 elements
`

// Basic conversion
const BasicConversion = `
// Basic type conversions
var i int = 42
var f float64 = float64(i)     // int to float64
var u uint = uint(f)           // float64 to uint

// Character conversion
var r rune = 'A'
var b byte = byte(r)           // rune to byte (if it fits in a byte)

// Slice to array and back
s := []int{1, 2, 3, 4}
var a [4]int
copy(a[:], s)                  // Copy slice to array
s2 := a[:]                     // Array to slice
`

// Numeric conversion
const NumericConversion = `
// Numeric conversions
i := 42
f := float64(i)               // int to float64: 42.0

f = 3.14
i = int(f)                    // float64 to int: 3 (truncated, not rounded)

// Possible data loss
var big int64 = 1<<40
small := int32(big)           // 1<<40 doesn't fit in int32, data loss

// Rune/byte conversions
r := '世'                      // Unicode code point 19990
b := byte(r)                  // Only keeps the lower 8 bits, data loss

// Safe handling of potential overflow
func safeConvertToInt8(x int64) (int8, error) {
    if x < math.MinInt8 || x > math.MaxInt8 {
        return 0, errors.New("value out of range for int8")
    }
    return int8(x), nil
}

val, err := safeConvertToInt8(1000)  // Returns error
`

// String conversion
const StringConversion = `
// String conversions
i := 42
s := strconv.Itoa(i)          // int to string: "42"

s = "3.14"
f, err := strconv.ParseFloat(s, 64)  // string to float64: 3.14

s = "true"
b, err := strconv.ParseBool(s)       // string to bool: true

s = "42"
i, err = strconv.Atoi(s)             // string to int: 42

// Using fmt package
s = fmt.Sprintf("%d", 42)            // int to string: "42"
s = fmt.Sprintf("%.2f", 3.14159)     // float to string with formatting: "3.14"

// Rune/byte slice conversions
bytes := []byte("hello")              // string to []byte
s = string(bytes)                     // []byte to string

runes := []rune("hello 世界")           // string to []rune
s = string(runes)                     // []rune to string
`

// Interface conversion
const InterfaceConversion = `
// Type assertions for interface conversion
var i interface{} = "hello"

// Safely check if interface holds a string
if s, ok := i.(string); ok {
    fmt.Println(s)            // "hello"
} else {
    fmt.Println("not a string")
}

// Type switch for multiple possibilities
switch v := i.(type) {
case string:
    // v has type string
    fmt.Println("string:", v)
case int:
    // v has type int
    fmt.Println("int:", v)
default:
    // no match
    fmt.Println("neither string nor int")
}

// Interface-to-interface conversion
type Stringer interface {
    String() string
}

type Printer interface {
    String() string
    Print()
}

var s Stringer
// p := Printer(s)  // Compile error unless s implements Print() too
`

// Struct conversion
const StructConversion = `
// Conversion between similar structs
type User struct {
    ID   int
    Name string
    Age  int
}

type UserDTO struct {
    ID   int
    Name string
    Age  int
}

// Converting between similar structs requires explicit field mapping
user := User{ID: 1, Name: "Alice", Age: 30}
userDTO := UserDTO{
    ID:   user.ID,
    Name: user.Name,
    Age:  user.Age,
}

// Using reflection for generic conversions
func convertStruct(src, dst interface{}) error {
    srcVal := reflect.ValueOf(src)
    dstVal := reflect.ValueOf(dst)
    
    // Ensure dst is a pointer
    if dstVal.Kind() != reflect.Ptr {
        return errors.New("destination must be a pointer")
    }
    
    dstVal = dstVal.Elem()
    srcType := srcVal.Type()
    dstType := dstVal.Type()
    
    for i := 0; i < srcType.NumField(); i++ {
        srcField := srcType.Field(i)
        
        // Try to find a corresponding field in dst
        if dstField, found := dstType.FieldByName(srcField.Name); found {
            if srcField.Type == dstField.Type {
                dstVal.FieldByName(srcField.Name).Set(srcVal.Field(i))
            }
        }
    }
    
    return nil
}

// Usage
user := User{ID: 1, Name: "Alice", Age: 30}
var dto UserDTO
convertStruct(user, &dto)
`

// Strconv example
const StrconvExample = `
// Using strconv package for string conversions
import "strconv"

// String to int
s := "42"
i, err := strconv.Atoi(s)  // i = 42

// Int to string
s = strconv.Itoa(i)        // s = "42"

// String to int64 with base
s = "FF"
i64, err := strconv.ParseInt(s, 16, 64)  // i64 = 255 (hex FF)

// Int64 to string with base
s = strconv.FormatInt(255, 16)  // s = "ff"

// String to float
s = "3.14159"
f, err := strconv.ParseFloat(s, 64)  // f = 3.14159

// Float to string
s = strconv.FormatFloat(f, 'f', 2, 64)  // s = "3.14"

// String to bool
s = "true"
b, err := strconv.ParseBool(s)  // b = true

// Bool to string
s = strconv.FormatBool(false)  // s = "false"
`

// Conversion pitfalls
const ConversionPitfalls = `
// Common conversion pitfalls

// Integer overflow
var i int8 = 127
i++  // Wraps around to -128

// Float to int truncation (not rounding)
f := 9.99
i := int(f)  // i = 9, not 10

// Character conversion loss
r := '世'  // Unicode code point 19990
b := byte(r)  // b = 102, lost most of the data

// Map key type must be comparable
type Point struct{ x, y float64 }
// m := map[Point]string{}  // Compiles
type SliceContainer struct{ slice []int }
// m := map[SliceContainer]string{}  // Compiler error: slice cannot be map key

// Interface nil vs typed nil
var p *int = nil
var i interface{} = p
fmt.Println(i == nil)  // false, because i has a type *int with value nil

// Type assertion on nil interface
var n interface{}
// fmt.Println(*n.(*int))  // Panic: interface conversion: interface is nil, not *int

// Safe approach
ptr, ok := i.(*int)
if !ok || ptr == nil {
    fmt.Println("Not a valid int pointer")
}
`
