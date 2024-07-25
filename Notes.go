// ---------------------------- INTRO ---------------------------------

// => What is GoLang or Go ?
// -> Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007.
// -> It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.
// -> Later in 2009 it was open sourced.

// => Use Cases of Go (What yet another language is needed?) -
// - Since, infrastructure changed a lot and Cloud infrastructure, Big networked computations clustors and Multicore Processors came into picture. And due to these changes, the infrastructure became more scalable, distributed, dynamic and have more capacity.
// - Due to these changes, the existing programming languages wasn't able to take full advantages of these infrastructures. They were capable of doing only 1 task at a time. Parallelism(Multithreading) and concurrency was not possible.
// - Languages that supports Multithreading -
// 		- C++, Java -> complex code writing and expensive int these languages
// 		- Go(This is the problem that Go solves.
// - Languages that does not supports Multithreading - Python, Nodejs

// => Main use cases -
// 1. For writing PERFOMANT Applications.
// 2. Running on scaled, distributed systems (typically on a cloud platform.)
// 3. It is now being used for writing automation task and cli applications for DevOps and SRE(Site Reliability Engineer) tasks as well.

// => Characteristics of GO -
// 1. Simple and readable syntax of dynamically typed language like Python.
// 2. Efficieny and safety of lower-level, statically typed language like C++.
// 3. Go is used on the server side applications like Microservices, web applications and database applications.
// 4. Some majors applications written in Go are - Docker, Hashicorps Vault., CockroachDB, Kubenetes
// 5. Faster build time, startup and run time.
// 6. Very Resource efficient (requires fewer resources like cpu, ram etc.).
// 7. It compiles into a single binary(machine code) and then it can be run across all platforms.

// ------------------------- INSTALLATION ----------------------------
// https://go.dev/doc/install

// -------------------------- FIRST PROGRAM ---------------------------
// package main
// import "fmt"

// func main () {
// 	fmt.Print("Hello World!")
// }

// OR
// package main

// func main () {
// 	print("Hello World!")
// }

// To execute this file -
// go run [filename]

// ----------------------------- BASICS -------------------------------
// => A. Packages -
// - package "package_name"
// - Package is used to define a package and also it shows that the current file belongs to that package.
// - This package system is used to organise our code and files in better and systematic manner.
// - Multiple files can belong to the same package.
// - A project can have multiple packages as well.

// ---------------------------------
// => B. Imports -
// import "package/file_name";
// or
// import ("package_name1" "package_name2" "package_name3")
// - "package main" is a special package name that tells go that this package is the main entry point of the application we are building.

// ---------------------------------
// => C. Comments -
// 1. Single line comment -
// This is a single line comment.

// 2. Multi line comment -
/**
 * This
 * is
 * a
 * multi
 * line
 * comment
 */

// ---------------------------------
// => D. Main Function -
// - This is the starting point of the program in go.
// - We cannot name this function as anything else other than "main".
// ex -
// package main
// import "fmt" // "fmt" is a part of go standard library

// func main () {
// 	fmt.Print("Print inside a main function!")
// }

// ---------------------------------
// => E. Taking using input -
// package main

// import (
// 	"fmt"
// 	"bufio"
// 	"os"
// 	"log"
// )

// var print = fmt.Println;

// func main () {
// 	print("What is your name");
// 	reader := bufio.NewReader(os.Stdin);

//  name, _ := reader.ReadString('\n');
// 	name, err := reader.ReadString('\n');
// 	if err == nil {
// 		print("Hello ", name)
// 	} else {
// 		log.Fatal(err)
// 	}
// }

// ---------------------------------
// => F. Building And Running A Simple Go Application Program -
// 1. Create a file "helloWorld.go" and a simple program as below
// package main

// import "fmt"

// func main() {
// 	fmt.Print("Hello world!")
// }

// 2. Now, we have to create a module of our application. A module is simply a go application that can contain mulitple packages. We can create a module with using the command - go mod init [module_name]
// 3. Now after we run the command our go application module file will be created and then we can simply build our application using the following command - "go build"
// 4. After our application is build, a [module_name].exe file will be created which we can simple run as a application by executing "./module_name" the command

// OR
// 1. We can also run a single file using the command  => go run filename.go

// ---------------------------------
// => G. Variables -

// 1. Declaring Variables using "var" keyword -
// var var_name var_data_type = var_value / literal
// or
// var var_name = var_value / literal
// eg - var amount float64 = 1000.50

// 2. Decalaring Variables without using the "var" keyword -
// var_name := var_value / literal
// eg - amount := 1000.50
// Note: In this method, variable data type will be inferred by the go from the declared variable value. We can not explicitly declare the type of the variable.

// 3. Declaring constant variables using the "const" keyword -
// const var_name var_data_type = var_value / literal
// eg - const gravity float64 = 9.81 