package main

import "fmt"
import "example.com/bank/fileoperation"

const accountBalanceFile ="balance.txt"

func main(){
 var accountBalance,err = fileoperation.GetFloatFromFile(accountBalanceFile)

 if err != nil{
	fmt.Println("ERROR")
	fmt.Println(err)
	fmt.Println("------")
	panic("Can't continue. Please re-run ")
 }
 fmt.Println("Go Bank")
 for i:=0; i <4 ; i++ {
   
 presentOptions()



var choice int
fmt.Println("Enter choice : ")
fmt.Scan(&choice)

switch choice{
case 1:
	fmt.Println("Balance :" ,accountBalance)
case 2:
	fmt.Println("Your deposit : ")
	var depositAmount float64
	fmt.Scan(&depositAmount)
	if depositAmount <=0{
		fmt.Println("Invalid Amount. Should be greater than 0")
		continue
	}
	accountBalance += depositAmount
	fmt.Println("Total Bal after deposit:",accountBalance)
	fileoperation.WriteFloatToFile(accountBalance,accountBalanceFile)

case 3:
	fmt.Println("Your withdraw Amt : ")
	var withdrawAmount float64
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <=0{
		fmt.Println("Invalid Amount. Must be greater than 0")
		continue
	}

	if( withdrawAmount> accountBalance){
		fmt.Println("Should not be greater than your account balance")
		continue
	}
	accountBalance -= withdrawAmount
	fmt.Println("Total Bal after withdraw:",accountBalance)
	fileoperation.writeFloatToFile(accountBalance,accountBalanceFile)

default:
	fmt.Println("Wrong option! Check again")
	fmt.Println("Thanks for choosing GO Bank")
	return
}




}


}

