package arraytools

// Mult erwartet ein Array und einen skalaren Faktor.
// Multipliziert jedes Element des Arrays mit dem Faktor.
func ScalarMult(a []float64, factor float64) {
	for i, el := range a {
		a[i] = el * factor
	}
}

// Add erwartet zwei Arrays der gleichen Länge.
// Addiert die Elemente der Arrays paarweise.
func Add(a, b []float64) {
	for i := 0; i < len(a); i++ {
		a[i] += b[i]
	}
}
