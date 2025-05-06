# Programming Workshop Assignments

This repository contains a series of programming workshops designed to test and improve your skills across various programming languages and concepts.

## Workshop Overview

| Workshop | Topic | Description |
|----------|-------|-------------|
| Workshop 1 | Problem Solving | Practical programming exercises in Python, JavaScript, Go, and Shell |
| Workshop 2 | Requirement Analysis | Learning to analyze and document requirements |
| Workshop 3 | Frontend Development | Creating a landing page using preferred technologies |
| Workshop 4 | Backend Development | API development with database modeling and documentation |
| Workshop 5 | Testing | Backend testing based on features and test cases |
| Workshop 6 | Code Review | Identifying and fixing issues in existing code |

## Workshop 1: Problem Solving

### Task 1: Matrix Rotation (Python)
**Objective**: Implement a function that rotates a 2D matrix 90 degrees clockwise.

**Target File**: `workshop-1/task-1/matrix_rotation.py`

**Requirements**:
- Create a Python function named `rotate_matrix`
- The function should accept a 2D matrix (list of lists) as input
- It should rotate the matrix 90 degrees to the right (clockwise)
- Return the rotated matrix

**Test Command**:
```shell
cd workshop-1/task-1
python -m unittest matrix_rotation_test.py
```

### Task 2: FizzBuzz++ (JavaScript)
**Objective**: Implement an enhanced version of the FizzBuzz algorithm.

**Target File**: `workshop-1/task-2/fizzbuzz.js`

**Rules**:
- If the number is divisible by 3, include "Fizz"
- If the number is divisible by 5, include "Buzz"
- If the number is divisible by 7, include "Woof"
- If divisible by multiple numbers, concatenate the words (e.g., divisible by 3 and 5 = "FizzBuzz")
- If not divisible by 3, 5, or 7, return the number itself

**Test Command**:
```shell
cd workshop-1/task-2
npm install
npm run test
```

### Task 3: Custom Sorting (Go)
**Objective**: Implement a custom sorting function for a slice of Person structs.

**Target File**: `workshop-1/task-3/custom_sort.go`

**Requirements**:
- Create a function named `SortPeopleByAge`
- The function should accept a slice of Person structs (`[]Person`)
- Sort the slice in-place by the `Age` field in ascending order

**Test Command**:
```shell
cd workshop-1/task-3
go test
```

### Task 4: Binary Tree Visualization (Shell Script)
**Objective**: Create a shell script that displays a binary tree structure.

**Target File**: `workshop-1/task-4/binary.sh`

**Requirements**:
- Implement a shell script that represents a binary tree
- Use a left-right structure format

**Execution Command**:
```shell
cd workshop-1/task-4
chmod +x binary.sh
./binary.sh
```

## Workshop 2: Requirement Analysis
Exercise focused on analyzing requirements and documenting specifications.

**Deliverables**:
- User stories
- Acceptance criteria
- POC (Proof of Concept)

## Workshop 3: Frontend Development
**Objective**: Create an attractive landing page using your preferred tools and technologies.

## Workshop 4: Backend Development
**Objective**: Develop APIs to support the provided UI design.

**Deliverables**:
- Working API implementation
- Database diagram
- API documentation (Swagger)

## Workshop 5: Testing
**Objective**: Implement backend tests according to specified features and test cases.

## Workshop 6: Code Review
**Objective**: Identify and fix problems in the provided code samples.
