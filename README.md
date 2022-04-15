# Go practice
small cheatsheet to help me migrate from .NET language to Go.


## Structure
The go apps are composed by modules and pacakges.

A package is all the content inside a  folder, and variables, functions, etc, are accessible between them.

Does an "early file" has access to a "later" one? (like f# does)

Similar idea of a namespace in .net

go mod init command to initialize a new module like the next:
`go mod init url.com/subdomain` the url subdomain has to be an actual url used in the future (this part is werid, investigate more)

and then create folders inside

the extension of the files is `.go`

the first line should contain the `pacakge` and when `main` is specified it compiles as an app not as a shared library.


the first that runs is `func main()` inside the pacakge main, so it is basically the entry point and they are mandatory for the app to run.

### Packages
To import pagackes we specify `import "PackageName"`
- where is this name defined?

- does this add them in all the files oronly in this one?.
- Does goIDE (jetbrains one) import this automatically?

to print something just `fmp.Println("text")`

multiples pacakges can be imported in one go (hehe):

````go
import (
	"fmt"
)
````

### Running and compiling

We can run, build and install the apps.

- run: `go run file`
- build: `go build file` | `go run .` -> compiles as executable with the name as the pacakge.
- install `go install` compiles the app in `bin` folder of the user (it can be changed, i dont think i will be using this.)

** how does it compiles it?  for your machine or how? can this be shared between windows/mac/linux? because it does generate an `.exe` file.

## variables
it is strongly typed

they have to be declared **outside** of the functions.

- does this affects to all the variables or only the "shared ones"
  ** "private" variables can be specified inside each function. it is done with the shortassigment `variable := 123`

this is how it looks:
````go
package main

import (
	"fmt"
	"reflect"
)

var (
	name   string
	course string
	module int
	clip   int
)

func main() {
	fmt.Println("name is the type", reflect.TypeOf(name))
}

````

so when a variable starts by **upercase** it can be accessed from outside the package. but if it is camel case only internally. (similar as public/private properties.)

if the values are not assigned, they are assigned with "default" values (similar as .net).

to print the variables, you have to pass them like `"this is a sentecence and a variable:", variablename, " with variables"`

basically commas to "interpolate" strings, or something like in C:
````go
fmt.Println("your name is", name, "and your age is", age)

formatedString := fmt.Sprintf("your name is %s and your age is %d", name, age)
fmt.Println(formatedString)
````




go supports type inference (like .net)
`variable := 12` maps to integer automatically.


To declare a variable without assigning any value we should do it like the next `var edad int`


- are they value types or reference types?
  - when using an asterisk is a reference type (sending the pointer)
  - when not using the asterisk is the value type

## Functions
functions can return more than one type, and seems very common that they return a value and then another parameter with the error.

For example the function `.Atoi` returns two values, `converted, err`.

if you don't use (to validate or use) any of the values, it's gonna return an error, to prevent this error you can discard that variable using `_` like in C#.

-Note: There is no extension methods in go, like there is in C#.


### Conversion
if we use the package `"strconv"` is like string conversion [Link](https://pkg.go.dev/strconv)

to convert an `int` to `string` you have to use `strconv.Itoa(age)`.

To convert an `string` to an integer we use `strconv.Atoi(stringage)`


### User input
we can request user input if we use `fmt.scanf()` inside you have to specify the type and the variable to be assigned (same as c)
```
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

like `fmt.Scanf("%d\n", &age)` also `\n` must be specified so when you click enter so the enter character does not get recorded.

And the variable specified must contain an ampersand `&` to get the assigned value.

## Conditional sentences

### If
Pretty much the same as everywhere, just it does not have brackets (they are optional); And the brackets need to be in the same line.

````go
if x > y {
    fmt.Println("x is bigger than Y")
} else if x == y {
    fmt.Println("y and X are equal")
} else {
    fmt.Println(" y is bigger")
}
````
the else has to be in the same line as the `}`

to add multiple conditions, like `and` (`&&`), `or` (`||`) finally `distinct` (`!=`)

### for
follows the same logic as other languages
````go
for i := 0; i < 10; i++ {
    fmt.Println("iteration number: ", i+1)
}
````

but with for you can simulate a while if you do
````go
i := 0
for  i < 10 {
    fmt.Println("iteration number: ", i+1)
    i++
}
````
or even an infinite loop if it just contains `for`, if you need to stop it you can use `break`
````go
for {
	//code
	//more logic with a break
}
````

## Arrays
normal creation of arrays, defining the size of it. `var testArray [10]string`


it can be also instantiated with values : `var testArray = [10]string{"first", "second", "third"}`

It does also support not specifying the size if the initialization is done `var testArray = []string{"first", "second", "third"}` (with a size of 3)


as in any other language we can use an indexer to access the position `testArray[0]`.

- `cap(array)` max size of the array.

- what about lists like in c#



### Matrix
normal as in other languages `var matrix[2][2]` and the same to access `matrix[][]`

### Slice
Similar to ranges in .net, we can specify an slice or a part of an array to store it in another variable.
````go
slice := testArray2[1:3]
fmt.Println(slice)
````
A slice stores the pointer, the length and the capacity.

There is no such things as .NET lists in golang, but you can append to an slice
````go
testArray2 = append(testArray2, "newone")
fmt.Println(testArray)
````



## API
use gin, added it as a dependency, to get it you have to use `go get .`


and just run it for a simple example.# Go practice
small cheatsheet to help me migrate from .NET language to Go.


## Structure
The go apps are composed by modules and pacakges.

A package is all the content inside a  folder, and variables, functions, etc, are accessible between them.

Does an "early file" has access to a "later" one? (like f# does)

Similar idea of a namespace in .net

go mod init command to initialize a new module like the next:
`go mod init url.com/subdomain` the url subdomain has to be an actual url used in the future (this part is werid, investigate more)

and then create folders inside

the extension of the files is `.go`

the first line should contain the `pacakge` and when `main` is specified it compiles as an app not as a shared library.


the first that runs is `func main()` inside the pacakge main, so it is basically the entry point and they are mandatory for the app to run.

### Packages
To import pagackes we specify `import "PackageName"`
- where is this name defined?

- does this add them in all the files oronly in this one?.
- Does goIDE (jetbrains one) import this automatically?

to print something just `fmp.Println("text")`

multiples pacakges can be imported in one go (hehe):

````go
import (
	"fmt"
)
````

### Running and compiling

We can run, build and install the apps.

- run: `go run file`
- build: `go build file` | `go run .` -> compiles as executable with the name as the pacakge.
- install `go install` compiles the app in `bin` folder of the user (it can be changed, i dont think i will be using this.)

** how does it compiles it?  for your machine or how? can this be shared between windows/mac/linux? because it does generate an `.exe` file.

## variables
it is strongly typed

they have to be declared **outside** of the functions.

- does this affects to all the variables or only the "shared ones"
  ** "private" variables can be specified inside each function. it is done with the shortassigment `variable := 123`

this is how it looks:
````go
package main

import (
	"fmt"
	"reflect"
)

var (
	name   string
	course string
	module int
	clip   int
)

func main() {
	fmt.Println("name is the type", reflect.TypeOf(name))
}

````

so when a variable starts by **upercase** it can be accessed from outside the package. but if it is camel case only internally. (similar as public/private properties.)

if the values are not assigned, they are assigned with "default" values (similar as .net).

to print the variables, you have to pass them like `"this is a sentecence and a variable:", variablename, " with variables"`

basically commas to "interpolate" strings, or something like in C:
````go
fmt.Println("your name is", name, "and your age is", age)

formatedString := fmt.Sprintf("your name is %s and your age is %d", name, age)
fmt.Println(formatedString)
````




go supports type inference (like .net)
`variable := 12` maps to integer automatically.


To declare a variable without assigning any value we should do it like the next `var edad int`


- are they value types or reference types?

## Functions
functions can return more than one type, and seems very common that they return a value and then another parameter with the error.

For example the function `.Atoi` returns two values, `converted, err`.

if you don't use (to validate or use) any of the values, it's gonna return an error, to prevent this error you can discard that variable using `_` like in C#.

-Note: There is no extension methods in go, like there is in C#.


### Conversion
if we use the package `"strconv"` is like string conversion [Link](https://pkg.go.dev/strconv)

to convert an `int` to `string` you have to use `strconv.Itoa(age)`.

To convert an `string` to an integer we use `strconv.Atoi(stringage)`


### User input
we can request user input if we use `fmt.scanf()` inside you have to specify the type and the variable to be assigned (same as c)
```
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

like `fmt.Scanf("%d\n", &age)` also `\n` must be specified so when you click enter so the enter character does not get recorded.

And the variable specified must contain an ampersand `&` to get the assigned value.

## Conditional sentences

### If
Pretty much the same as everywhere, just it does not have brackets (they are optional); And the brackets need to be in the same line.

````go
if x > y {
    fmt.Println("x is bigger than Y")
} else if x == y {
    fmt.Println("y and X are equal")
} else {
    fmt.Println(" y is bigger")
}
````
the else has to be in the same line as the `}`

to add multiple conditions, like `and` (`&&`), `or` (`||`) finally `distinct` (`!=`)

### for
follows the same logic as other languages
````go
for i := 0; i < 10; i++ {
    fmt.Println("iteration number: ", i+1)
}
````

but with for you can simulate a while if you do
````go
i := 0
for  i < 10 {
    fmt.Println("iteration number: ", i+1)
    i++
}
````
or even an infinite loop if it just contains `for`, if you need to stop it you can use `break`
````go
for {
	//code
	//more logic with a break
}
````

## Arrays
normal creation of arrays, defining the size of it. `var testArray [10]string`


it can be also instantiated with values : `var testArray = [10]string{"first", "second", "third"}`

It does also support not specifying the size if the initialization is done `var testArray = []string{"first", "second", "third"}` (with a size of 3)


as in any other language we can use an indexer to access the position `testArray[0]`.

- `cap(array)` max size of the array.

- what about lists like in c#



### Matrix
normal as in other languages `var matrix[2][2]` and the same to access `matrix[][]`

### Slice
Similar to ranges in .net, we can specify an slice or a part of an array to store it in another variable.
````go
slice := testArray2[1:3]
fmt.Println(slice)
````
A slice stores the pointer, the length and the capacity.

There is no such things as .NET lists in golang, but you can append to an slice
````go
testArray2 = append(testArray2, "newone")
fmt.Println(testArray)
````


## structs 
TODO: fill the documentation, but basically objects.


## API
in here I realized the next:
- simple Get/Create Api Endpoints
- What I think is DI in Golang (this part still not clear, I can't find proper stuff on internet)
- Mysql for data persistence
- Unit testing 
- Swagger (pending)

use gin, added it as a dependency, to get it you have to use `go get .`

- investigate if there is any other way than "gin" to run apis.

But it seems quite simple.

and just run it for a simple example.



### how I organized it (structure)

I created a similar structure of what .NET apis contains.

The main package and the main function contains the "startup", so the endpoints and the configuration (which repo is using, etc)
- NOTE: the "inMemory" is commented, and to run the database one it is using docker.


I tried to implement Dependency Injection (not sure if it is done correctly), I created a file called `albumService` which contains the interface and the "logic" is quite simple so a single service is being used.

this `AlbumService` is the one I will create the unit testing later on (mocking the information).

Eventually I would like to do also integration testing, but it seems that i can do it with docker on an easy way. 

And then I did the repos. I created an interface, so I can switch between the `InMemory` and the `Sql` one.

### Alerts/break
the way of breaking the code, like an exception is using `panic("implement me")`

### Connect to SQL
To connect to MySQL i had to import

````go

"database/sql"
_ "github.com/go-sql-driver/mysql"
````
- Investigate what the `_` character is for.

then just connect: 
````go
//TODO: find a way to define this in a file like the appsettings
const (
	username = "gouser"
	password = "gopass"
	hostname = "127.0.0.1:3307"
	dbname   = "testgo"
)

func GetDb() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", username+":"+password+"@("+hostname+")/"+dbname)
	return
}

````
- Investigate how to specify that configuration in files like in C# or any other technique.

### Query
To query, we have to type the query manually 
- Investigate if there is orm like EF or Dapper,
- what about slq injection.
- Investigate about dispose the connections 

the query returns the `row` and the `error`
````go
rows, error := repo.Db.Query("select * from album where Id = ?")
````

then we should iterate rows, using `.Next()` and map into an object. Even if there is only one element.

- When an item does not exist, for example by id, it does return a `Default` object. Is there something similar as `nullable` in C#?

### Insert
To insert an element the idea is the same, the query and then `.Exec()`
````go
func (repo SqlRepository) AddAlbum(album Dtos.Album) {
	query := "INSERT INTO Album(Title, Artist, Price) " +
		"VALUES (?, ?, ?)"
	repo.Db.Exec(query, album.Title, album.Artist, album.Price)

}
````
#### Get last inserted Id
to get the last inserted id we can retrieve the value of `Exec()`
````go
result, err :=repo.Db.Exec(qu...)
````
and then result has a function called `.LastInsertedId()`


### Update
update is the same story, but is boring to have to do it manually

## Testing
The name of the file must end in `_test` and the test name must start for `Test`.
wtf.

I am trying to figure it out if there is any library like `moq` in c#; all examples I find on internet are with fakes.


the assertion is doing an if and trigger a failure if it fails.