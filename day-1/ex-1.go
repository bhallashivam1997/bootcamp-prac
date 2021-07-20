package main

import "fmt"

type Matrix struct {
	row int
	col int
	grid [][]int
}

func getRow(mat Matrix) int{
	return mat.row
}

func getCol(mat Matrix) int{
	return mat.col
}

func setElement(i int , j int , val int , mat Matrix) {
	mat.grid[i][j] = val
}

func printMatrix(mat Matrix){
	fmt.Printf("{ \n")
	for _, i := range mat.grid {
		fmt.Printf("[")
		for _, j := range i {
			fmt.Printf("%d , ", j)
		}
		fmt.Printf("]")
		fmt.Println()
	}
	fmt.Printf("} \n")
}

func addMatrix(mat1 Matrix , mat2 Matrix) [][]int{
	var n1,m1 = mat1.row,mat1.col
	grid1 := make([][]int, n1)
	for i := 0; i < n1; i++ {
		grid1[i] = make([]int, m1)
	}
	for i:=0;i<n1;i++{
		for j:=0;j<m1;j++{
			grid1[i][j] = mat1.grid[i][j] + mat2.grid[i][j]
		}
	}
	return grid1
}



func main() {

	// initializing 2D MATRIX using the struct
	var n1,m1 = 2,3
	grid1 := make([][]int, n1)
	for i := 0; i < n1; i++ {
		grid1[i] = make([]int, m1)
	}
	var mat1 = Matrix{n1,m1,grid1}

	var n2,m2 = 2,3
	grid2 := make([][]int, n2)
	for i := 0; i < n1; i++ {
		grid2[i] = make([]int, m2)
	}
	//var mat2 = Matrix{n1,m1,grid1}
	//mat3 := addMatrix(mat1,mat2)
	//fmt.Printf("%d \n",getRow(mat1))
	//fmt.Printf("%d \n",getCol(mat1))
	//setElement(0,0,5,mat1)
	printMatrix(mat1)
}
