//Program 5 Advanced from CSC101
//An attempt to learn basics in go by translating a python project
//I made into golang
//Logan Pinel
//October 22, 2024

// setup program
package main

import (
	"fmt"     //for strings
	"math"    //for power
	"strconv" //will allow us to convert strings to ints
)

// create a dictionary (MAP IN GOLANG) for the operators and assign each operator a value
// based on PEMDAS with exponents being the largest value
var operatorsDictionary = map[string]int{ //map refers to dictionary //type string //key int
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

// getOperators function will trace through the eqaution and get the operators
// and the numbers
func getOperators(equation string) ([]string, []string) { //have to identify return types in golang, we are returning two lists of strings
	//create two lists to store the equations operators and numbers
	//their types need to be strings
	var operators []string
	var nums []string
	//blank string to hold the number if it is multiple digits
	numStr := ""
	//get the operators by looping through the equation
	for _, element := range equation { //range in golang returns two values: index and element so we use a blank identifier to not get the index
		//use boolean var "exists" ot check if the operator exits in the map
		if _, exists := operatorsDictionary[string(element)]; exists { //ignore the value of the key in the map so use _
			operators = append(operators, string(element)) //append the operator //append(list, value)
			//append the number to the list because operator is found
			nums = append(nums, string(numStr))
			//reset the numStr
			numStr = ""
		} else {
			numStr += string(element)
		}
	}
	//if loop is over and still a number because no more operators were found
	//append it to the nums list aswell
	if numStr != "" {
		nums = append(nums, numStr)
	}

	return operators, nums
}

// getGreatestOperators function will find the greatest operator according to PEMDAS
func getGreatestOperator(operators []string) string {
	//set 1st element as placeholder to check greater than
	greatestOperator := operators[0]
	//loop through the operators list
	for _, op := range operators {
		//find the greates operator and store it in greatestOperator
		if operatorsDictionary[op] > operatorsDictionary[greatestOperator] {
			greatestOperator = op
		}
	}
	return greatestOperator
}

// Golang doesnt have a build in indexOf function so lets build one
func indexOf(slice []string, value string) int { //slice is another form of list specific to golang
	for i, v := range slice {
		if v == value {
			return i //return index if found
		}
	}
	return -1 //if nothing is found
}

// calculateOneExpression function will calculate a single expression at a time
// from the equation by the rules of PEMDAS
func calculateOneExpression(operators []string, nums []string) ([]string, []string) {
	//call getGreatestOperator
	greatestOperator := getGreatestOperator(operators)
	//get the index of the greatest operator
	indexOfGreatestOp := indexOf(operators, greatestOperator)
	//get the two numbers to the left and right of the operator
	//need to cast the strings as float64 so we can use strconv.ParseFloat("string", type)
	num1, _ := strconv.ParseFloat(nums[indexOfGreatestOp], 64) //strconv.ParseFloat returns two variables but we dont need the second one so use _
	num2, _ := strconv.ParseFloat(nums[indexOfGreatestOp+1], 64)

	//Determine what the greatestOperator is and execute the correct calculation using if-statements
	var answer float64
	//all the operator executions
	if greatestOperator == "^" {
		answer = math.Pow(num1, num2)
	} else if greatestOperator == "*" {
		answer = num1 * num2
	} else if greatestOperator == "/" {
		answer = num1 / num2
	} else if greatestOperator == "+" {
		answer = num1 + num2
	} else if greatestOperator == "-" {
		answer = num1 - num2
	}

	//update the equation now that we solved one step
	//convert answer to a string using FormatFloat() which formats a float variable into a string
	//'f' represents float in decimal notation //-1 represents smallest number of digits with greatest precision //64 represents the bit value
	nums[indexOfGreatestOp] = strconv.FormatFloat(answer, 'f', -1, 64) //replace the left number with the answer
	//find the index to delete from the slice
	numsIndexToDelete := indexOfGreatestOp + 1
	//update nums list
	nums = append(nums[:numsIndexToDelete], nums[numsIndexToDelete+1:]...) //... is a variadic function and used to expand a slice
	//this gets the whole list past the indexToDelete and ... appends it one at a time
	//Remove the operator aswell
	operatorsIndexToDelete := indexOfGreatestOp
	//update operators list
	operators = append(operators[:operatorsIndexToDelete], operators[operatorsIndexToDelete+1:]...)

	return operators, nums
}

// getAnswer() function will continually call the calculateOneExpression function untill
// there is only one element left in the nums list i.e. the answer
func getAnswer(equation string) string {
	//call getOperatorsandNums to get the two lists from equation
	operators, nums := getOperators(equation)
	//while there are still operators in the list keep solving one equation at a time
	for len(operators) > 0 {
		operators, nums = calculateOneExpression(operators, nums)
	}

	//once loop is broken there should only be one element in nums
	return nums[0]
}

// main() method
func main() {
	//welcome message
	fmt.Println("Welcome to the calculator!")
	fmt.Println("This calculator can use +, -, *, /, ^ operators.")
	//input the equation from the user
	fmt.Print("Please enter an equation (no spaces): ")
	var equation string
	fmt.Scan(&equation)
	//call getAnswer() function
	answer := getAnswer(equation)
	//print answer
	fmt.Println(answer)
}
