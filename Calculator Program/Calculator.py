#Calculator program takes the foundational code from my CSC101 courses advanced Program 5
#to create a text based calculator.
#Logan Pinel
#October 25, 2024

#create a dictionary for the operators and assign each operator a value 
#based on PEMDAS with exponents being the largest value
operatorsDictionary = {
    "+": 1,
    "-": 1,
    "*": 2,
    "/": 2,
    "^": 3
}

#getOperators subroutine will trace through the eqaution and get the operators
#and the numbers
def getOperatorsandNums(equation):
    #create two lists to store the equations operators and numbers
    operators = []
    nums = []
    #blank string to hold the number if it is multiple digits
    numStr = ""
    #get the operators by looping through the equation
    for i in equation:
        if i in operatorsDictionary:
            operators.append(i)
            #append the number to the list because operator is found
            nums.append(numStr)
            #reset the numStr
            numStr = ""
        #keep adding the digits to the number string
        else:
            numStr += i
        
    #if loop is over and still a number because no more operators were found
    #append it to the nums list aswell
    if numStr:
        nums.append(numStr)

    return operators, nums

#getGreatestOperators subroutine will find the greatest operator according to PEMDAS
def getGreatestOperator(operators):
    #set 1st element as placeholder to check greater than
    greatestOperator = operators[0]
    #loop through the operators list
    for op in operators:
        #find the greates operator and store it in greatestOperator
        if operatorsDictionary[op] > operatorsDictionary[greatestOperator]:
            greatestOperator = op

    return greatestOperator

#calculateOneExpression subroutine will calculate a single expression at a time
#from the equation by the rules of PEMDAS
def calculateOneExpression(operators, nums):
    #call getGreatestOperator
    greatestOperator = getGreatestOperator(operators)
    #get the index of the greatest operator
    indexOfGreatestOp = operators.index(greatestOperator)
    #get the two numbers to the left and right of the operator
    num1 = int(nums[indexOfGreatestOp]) #the indexOfGreatesOp corresponds to the left number in expression since nums are stored in a different list
    num2 = int(nums[indexOfGreatestOp + 1]) #+1 to get the right number in expression

    #Determine what the greatestOperator is and execute the correct calculation using if-statements
    answer = 0
    #power
    if greatestOperator == "^":
        answer = num1 ** num2
    #multiplication
    elif greatestOperator == "*":
        answer = num1 * num2
    #division
    elif greatestOperator == "/":
        answer = num1 / num2
    #addition
    elif greatestOperator == "+":
        answer = num1 + num2
    #subtraction
    elif greatestOperator == "-":
        answer = num1 - num2
    
    #update the equation now that we solved one step
    nums[indexOfGreatestOp] = answer #replace the left number with the answer
    del nums[indexOfGreatestOp+1] #delete the right number since we equated the expression
    #Remove the operator aswell
    del operators[indexOfGreatestOp]

    return operators, nums

#getAnswer() subroutine will continually call the calculateOneExpression subroutine untill 
#there is only one element left in the nums list i.e. the answer
def getAnswer(equation):
    #call getOperatorsandNums to get the two lists from equation
    operators, nums = getOperatorsandNums(equation)
    #while there are still operators in the list keep solving one equation at a time
    while operators:
        operators, nums = calculateOneExpression(operators, nums)
    
    #once loop is broken there should only be one element in nums
    return nums[0]

#welcome message
print("Welcome to the calculator!")
print("This calculator can use +, -, *, /, ^ operators.")
#input the equation from the user
equation = input("Please enter an equation (no spaces): ")
#call getAnswer()
answer = getAnswer(equation)
#print the answer
print(answer)

