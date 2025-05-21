// go run main.go

/* 
Slice maintains 3 things
Slice is part of array which has pointer , length and capacity;

-	Pointer : denote the staring   memory location adddress of the slice
-	Length: full length of the slice =3
-	Capacity: Start slice to last array. here, 5


*/
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


// previous things

func example1(){

	// a:=[3]int{1,2,3}; // This is an array
	s1:=[]int{1,2,3} // it becomes slice and its called slice literal
	fmt.Println("s1:",s1,", length:",len(s1),", capacity",cap(s1))


	s2:=make([]int ,3)
	fmt.Println("s2:",s2,", length:",len(s2),", capacity",cap(s2)) // s2: [0 0 0] , length: 3 , capacity 3
	s2[1]=5;
	fmt.Println("s2:",s2,", length:",len(s2),", capacity",cap(s2)) 


	s3 :=make([]int, 3,5);
	fmt.Println("s3:",s3,", length:",len(s3),", capacity",cap(s3)) // s3: [0 0 0] , length: 3 , capacity 5

	var s4 []int // nil slice or emapty slice
	fmt.Println((s4))
	fmt.Println("s4:",s4,", length:",len(s4),", capacity",cap(s4)) 

	s4 = append(s4,66) // s4: [] , length: 0 , capacity 0
	s4 = append(s4,40)
	fmt.Println("s4:",s4,", length:",len(s4),", capacity",cap(s4))  // s4: [66 40] , length: 2 , capacity 2
}

func slice_exectution2(){



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

	

}

func example3(){
	var x []int 
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y:=x // [1,2,3]

	x = append(x, 4)
	y = append(y, 5) // // [1,2,3,5]

	x[0]=10

	fmt.Println(x)
	fmt.Println(y) // [10,2,3,5]

	
}


func changeSlice(p []int) []int{
	p[0]=10 // [5,6,7] => [10,6,7]
	p = append(p, 11) // [10,6,7,101]
	return p  // [10,6,7,101]
}

func example4(){

	x:=[]int {1,2,3,4,5}
	x = append(x, 6)
	x = append(x, 7) // [1,2,3,4,5,6,7]

	a:=x[4:] // start from 4 index // [5,6,7]
	y:=changeSlice(a) // [10,6,7,101]

	fmt.Println(x) // [1,2,3,4,5,67] => [1,2,3,4,10,6,7] directly mutated the 
	fmt.Println(y) // [10,6,7,101]

	fmt.Println(x[0:8]) // [1 2 3 4 10 6 7 11]  start index 0 and print till index 8. where length 7

}



