package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile ="balance.txt"
func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
    os.WriteFile(accountBalanceFile,[]byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error){
	data, err :=os.ReadFile(accountBalanceFile)
	if(err!= nil){
		return 1000, errors.New("Failed ro read file")
	}
	balanceText := string(data)
	balance, err :=strconv.ParseFloat(balanceText,64)
	if err != nil{
		return 1000, errors.New("Failed to parse the value")
	}
	return balance, nil
}

func main(){
 var accountBalance,err = getBalanceFromFile()

 if err != nil{
	fmt.Println("ERROR")
	fmt.Println(err)
	fmt.Println("------")
	panic("Can't continue. Please re-run ")
 }
 fmt.Println("Go Bank")
 for i:=0; i <4 ; i++ {
   
 

 fmt.Println("What do you want to do ?")
 fmt.Println("1. Check balance")
 fmt.Println("2. Deposit Money")
 fmt.Println("3. Withdraw Money")
 fmt.Println("4. Exit")

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
	writeBalanceToFile(accountBalance)

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
	writeBalanceToFile(accountBalance)

default:
	fmt.Println("Wrong option! Check again")
	fmt.Println("Thanks for choosing GO Bank")
	return
}




}


}