package main
/* File with 'main' function, only executing test input data 

	available test inputs:

	sudokuInput_easy
	
	sudokuInput_medium

	sudokuInput_hard
	sudokuInput_hard2
	sudokuInput_hard3

	sudokuInput_GOD
	sudokuInput_GOD2
*/


func main() {
    a := NewSudoku(sudokuInput_hard2)
    if a != nil {
      a.resolve()  
    }  
}