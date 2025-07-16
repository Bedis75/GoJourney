package main

import (
	"fmt"
)

func add(a int , b int){
	fmt.Println("the sum of a and b is ",a+b)
}
func remove(a int , b int){
	fmt.Println("after remove b from a we have ",a-b)
}

func equation(age []int, action func(int,int)){
	for i:=0;i<len(age)-1;i++{
		fmt.Printf("the value after action for %d and %d is: ", age[i], age[i+1])
		action(age[i], age[i+1])
	}
}

func main() {
	/*name := "Bedis"
	number1 := 10
	number2 := 20
	number3 := number1 + number2
	ch1 := fmt.Sprintf("hello my name is %s and the sum of %d and %d is %d",name,number1,number2,number3)
	fmt.Printf("the saved string is %q \n",ch1)
	var ages  = []int{10,20,30,40,50}
	fmt.Printf("the ages are %v \n",ages)

	ages = append(ages,60)
	fmt.Printf("the ages are %v \n",ages)

	ages[1] = 25
	fmt.Printf("the ages are %v \n",ages)
	fmt.Println("the length of the ages slice is", len(ages))
	ages1 := ages[1:3]
	ages1 := ages[:3]
	fmt.Printf("the ages1 slice is %v \n",ages1)
	name := "Hello my name is Bedis Bensaid and im 21 years old"
	fmt.Println(strings.Contains(name, "21"))
	fmt.Println(strings.ReplaceAll(name, "21", "22"))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Index(name,"Bedis"))
	i:=0
	for i<5{
		fmt.Println("la valeur de i est :",i)
		i++
	}
	for j:=0;j<5;j++{
		fmt.Println("la valeur de j est :",j)
	}
	for index , value :=range ages {
		fmt.Println("the value at index ",index," is ",value)
	}*/
	var ages  = []int{10,20,30,40,50}
	add(10,20)
	remove(40,15)
	equation(ages,add)
	equation(ages,remove)


}
