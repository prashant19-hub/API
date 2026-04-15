package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    x, _ := strconv.Atoi(scanner.Text())
    
    fmt.Println(x * 31)
}

//Question:

// New Year Resolution
// To start the year 
// 2026
// 2026 off, Chef made a resolution to exercise daily. He decided to do exactly 
// X
// X push-ups every day.

// If he sticks to his resolution, how many push-ups will he do in the month of January?

// Note that the month of January has 
// 31
// 31 days.

// Input Format
// The first and only line of input will contain a single integer 
// X
// X, denoting the number of push-ups Chef does every day.
// Output Format
// Output a single integer: the number of push-ups Chef will do in January.

// Constraints

// 1≤X≤100