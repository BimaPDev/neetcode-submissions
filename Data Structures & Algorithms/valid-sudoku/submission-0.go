func isValidSudoku(board [][]byte) bool {
	seen := make(map[string]bool)
	for row := 0; row < 9; row++ {
		for colum := 0; colum < 9; colum++ {
			cell := board[row][colum]
			if cell == '.' { continue }
			box := (row/3)*3 + (colum/3)
			keyRow := fmt.Sprintf("r%d-%c", row, cell)
			keyCol := fmt.Sprintf("c%d-%c", colum, cell)
			keyBox := fmt.Sprintf("b%d-%c", box, cell)

			check := []string{keyRow, keyCol, keyBox}
			for _, k := range check {
				if seen[k] {
					return false
				}
				seen[k] = true
 			}
 		}
	}
	return true
}