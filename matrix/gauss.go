package matrix

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {
	m.ScalarMultRow(col, 1/m[col][col]) //Diagonal muss durch sich selbst geteil werden
}

// EliminateBelow erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen unter der Zeile row mit -1/Matrix[i][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(row int) {

	for i := row + 1; i < len(m); i++ {
		m.ScalarMultRow(i, -1/m[i][row])
		m.AddRows(i, row)
	}

}

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(row int) {
	for i := row - 1; i >= 0; i-- {
		m.ScalarMultRow(i, -1/m[i][row])
		m.AddRows(i, row)
	}
}

// fixt -0 zu 0
func almostEqual(a, b float64) bool {
	return a-b < 1e-10 && b-a < 1e-10
}
func fixZeros(m Matrix) {
	for i := range m {
		for j := range m[i] {
			if almostEqual(m[i][j], 0) {
				m[i][j] = 0
			}
		}
	}
}

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {
	for i := 0; i < len(m); i++ {
		m.Normalize(i)
		m.EliminateBelow(i)

	}
	fixZeros(m)
}

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {
	for i := len(m) - 1; i >= 0; i-- {
		m.Normalize(i)
		m.EliminateAbove(i)

	}
	fixZeros(m)
}

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {
	m.UpperTriangular()
	m.LowerTriangular()
}
