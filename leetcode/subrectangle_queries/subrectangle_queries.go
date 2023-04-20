package main

func main() {

}

type SubrectangleQueries struct {
	rec [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rec: rectangle,
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	// 0:0 0:1 0:2
	// 1:0 1:1 1:2
	// 2:0 2:1 2:2
	// 3:0 3:1 3:3
	tRow := len(this.rec)
	tCol := len(this.rec[0])

	for i := 0; i < tRow; i++ {
		for j := 0; j < tCol; j++ {
			if i >= row1 && j >= col1 {
				if i <= row2 && j <= col2 {
					this.rec[i][j] = newValue
				}
			}
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {

	tRow := len(this.rec)
	tCol := len(this.rec[0])

	for i := 0; i < tRow; i++ {
		for j := 0; j < tCol; j++ {
			if i == row && j == col {
				return this.rec[i][j]
			}
		}
	}
	return 0
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
