package main

import "fmt"

func main(){
	var accountBalance = 1000.0
 fmt.Println("Go Bank")
 fmt.Println("What do you want to do ?")
 fmt.Println("1. Check balance")
 fmt.Println("2. Deposit Money")
 fmt.Println("3. Withdraw Money")
 fmt.Println("4. Exit")

var choice int
fmt.Println("Enter choice : ")
fmt.Scan(&choice)


if choice == 1 {
  fmt.Println("Balance :" ,accountBalance)
  
}else if choice ==2 {
	fmt.Println("Your deposit : ")
	var depositAmount float64
	fmt.Scan(&depositAmount)
	if depositAmount <=0{
		fmt.Println("Invalid Amount. Should be greater than 0")
		return
	}
	accountBalance += depositAmount
	fmt.Println("Total Bal after deposit:",accountBalance)

}else if choice == 3{
	fmt.Println("Your withdraw Amt : ")
	var withdrawAmount float64
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <=0{
		fmt.Println("Invalid Amount. Must be greater than 0")
		return
	}

	if( withdrawAmount> accountBalance){
		fmt.Println("Should not be greater than your account balance")
		return
	}
	accountBalance -= withdrawAmount
	fmt.Println("Total Bal after withdraw:",accountBalance)
}else{
    fmt.Println("Wrong option! Check again")
}
}