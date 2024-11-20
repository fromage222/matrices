package matrixtools

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {
	// TODO
	return []float64{}
}

// HINT
// Die Matrix ist i.W. ein Array aus Arrays.
// Also eine Liste aus Zeilen, die i-te Zeile ist dann das i-te Element der Liste.

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	col := make([]float64, len(m))
	// TODO
	return col
}

// HINT
// Erzeugen Sie sich eine neue Liste, deren Länge die Anzahl der Zeilen der Matrix ist.
// Füllen Sie diese Liste mit den Elementen der j-ten Spalte, also mit m[0][j], m[1][j], ..., m[n][j].

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {
	// TODO
}

// HINT
// Iterieren Sie über die Elemente der Zeilen und addieren Sie sie paarweise.
// Schreiben Sie das Ergebnis zurück in die erste Zeile.
// Alternativ: Verwenden Sie die Funktion `Add` aus dem Paket `arraytools`.

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {
	// TODO
}

// HINT
// Iterieren Sie über die Elemente der Zeile und multiplizieren Sie sie mit dem Faktor.
// Alternativ: Verwenden Sie die Funktion `ScalarMult` aus dem Paket `arraytools`.

// Transpose erwartet eine Matrix und liefert ihre Transponierte.
// D.h. alle Zeilen der ersten Matrix werden zu Spalten der Transponierten und umgekehrt.
func Transposed(m [][]float64) [][]float64 {
	transposed := make([][]float64, len(m[0]))
	// TODO
	return transposed
}

// HINT
// Erzeugen Sie eine neue Matrix, die so viele Zeilen hat wie die Spalten der ursprünglichen Matrix.
// Verwenden Sie dann eine geschachtelte Schleife, die über alle Zeilen der neuen Matrix läuft.
// Die innere Schleife muss jeweils eine neue Zeile der transponierten Matrix erzeugen und dann
// die Elemente dieser Zeile mit den Elementen der Spalte der ursprünglichen Matrix füllen.
