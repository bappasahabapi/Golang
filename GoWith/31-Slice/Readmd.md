## go run main.go

## Interview Topic asked? 

- 01. Make Slice from Array
- 02. Make Slice from slic
- 03. Slice Literal
- 04. Declear Slice useing Make()
- 05. Declear Slice useing Make() function with len and capacity
- 06. var s4 []int //` nil slice or empty slice`
- 07. variadic function //` send unlimited argument in slice` len and cap is equal here


## Explaination:

- How memory store in slice 

```go
	var s []int ; // [], ptr=nill/empty, len=0, cap=0
	fmt.Println("s:",s,", length:",len(s),", capacity",cap(s)) 

	s = append(s, 10) // [10] , ptr=address-1024, len=0+1 =1 ,cap =1
	fmt.Println("s:",s,", length:",len(s),", capacity",cap(s)) 

	s = append(s, 20) // [10,20] , ptr=address-1028 (1024+4byte) , len=1+1 =2 ,cap =1*2=2	
	fmt.Println("s:",s,", length:",len(s),", capacity",cap(s)) 

	s = append(s, 30) // [10,20,30] , ptr=address-1032 (1028+4byte) , len=2+1 =3 ,cap =2*2=4	
	fmt.Println("s:",s,", length:",len(s),", capacity",cap(s)) 
	
	s = append(s, 40) // [10,20,30,40] , ptr=address-1036 (1032+4byte) , len=3+1 =4 ,cap =2*2=4	
	fmt.Println("s:",s,", length:",len(s),", capacity",cap(s)) 

	s = append(s, 50) // [10,20,30,40,50] , ptr=address-1040 (1036+4byte) , len=4+1 =5 ,cap =4*2=4	
	fmt.Println("s:",s,", length:",len(s),", capacity",cap(s)) 

```
```shell
❯ go run main.go
s: [] , length: 0 , capacity 0
s: [10] , length: 1 , capacity 1
s: [10 20] , length: 2 , capacity 2
s: [10 20 30] , length: 3 , capacity 4
s: [10 20 30 40] , length: 4 , capacity 4
s: [10 20 30 40 50] , length: 5 , capacity 8
```

### 01

<p> Slice maintains 3 things <p/>
Slice is part of array which has pointer , length and capacity; 

	-	`Pointer` : denote the staring   memory location adddress of the slice
	-	`Length:` full length of the slice =3
	-	`Capacity:` Start slice to last array. here, 5

--- 
- 01. Make Slice from Array

```go
package main

import "fmt"

func main(){
	//arr index start   0    1    2   3      4          5 
	arr :=[6]string {"This","is","a","Go","interview","question"}
	// fmt.Println("Array length is: ",len(arr))
	fmt.Println(arr) // [This is a Go interview question]

	slice :=arr[1:4] 
	fmt.Println(slice) //[is a Go]
}

```
- 02. Make Slice from slice
```go
package main

import "fmt"

func main(){

	// arr :=[6]string {"This","is","a","Go","interview","question"}
	// fmt.Println(arr) 

	//todo : make slice from an array
	// slice :=arr[1:4] 
	// fmt.Println(slice) //[is a Go]

	//todo: make slice from slice
	s1 :=slice[1:3]
	fmt.Println("s1:",s1,", length:",len(s1),", capacity",cap(s1)) // s1: [a Go] , length: 2 , capacity 4
}

```

- 03. Slice Literal
```go
a:=[3]int{1,2,3}; // This is an array
s:=[]int{1,2,3} // it becomes slice and its called slice literal
```

- 04. Make()
```go
s2:=make([]int ,3)
fmt.Println("s2:",s2,", length:",len(s2),", capacity",cap(s2))
```
- 05. 

```go
s2:=make([]int ,3)
fmt.Println("s2:",s2,", length:",len(s2),", capacity",cap(s2)) // s2: [0 0 0] , length: 3 , capacity 3

s3 :=make([]int, 3,5);
fmt.Println("s3:",s3,", length:",len(s3),", capacity",cap(s3)) // s3: [0 0 0] , length: 3 , capacity 5

var s4 []int // nil slice or emapty slice
```
- 06. 

```go
var s4 []int // nil slice or emapty slice
fmt.Println((s4))
fmt.Println("s4:",s4,", length:",len(s4),", capacity",cap(s4)) 

s4 = append(s4,66) // s4: [] , length: 0 , capacity 0
s4 = append(s4,40)
s4 = append(s4,40)
fmt.Println("s4:",s4,", length:",len(s4),", capacity",cap(s4))  // s4: [66 40] , length: 2 , capacity 2
```
- 07.  variadic function

```go
package main

import "fmt"



//variadic funciotn : we can send unlimited argument like this syntex
func variadic(numbers ...int){
	fmt.Println("slice:",numbers)
	fmt.Println("len:",len(numbers))
	fmt.Println("cap:",cap(numbers))
}

func main(){

	variadic(3,5,7,8)
}
```

```shell
❯ go run main.go

slice: [3 5 7 8]
len: 4
cap: 4

```