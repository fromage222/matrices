package matrixtools

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {
	return m[i]
}

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	col := make([]float64, len(m))
	for i := 0; i < len(m); i++ {
		col[i] = m[i][j]
	}
	return col
}

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {
	for k := 0; k < len(m[0]); k++ {
		m[i][k] += m[j][k]
	}
}

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {
	for n := 0; n < len(m[0]); n++ {
		m[i][n] = m[i][n] * factor
	}
}

// Transpose erwartet eine Matrix und liefert ihre Transponierte.
// D.h. alle Zeilen der ersten Matrix werden zu Spalten der Transponierten und umgekehrt.
func Transposed(m [][]float64) [][]float64 {
	transposed := make([][]float64, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		transposed[i] = GetCol(m, i)
	}
	return transposed
}
