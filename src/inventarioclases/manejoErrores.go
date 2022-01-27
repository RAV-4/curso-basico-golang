package inventarioclases

func isError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Ejemplo de uso
func creandoUnError() {
	_, err := strconv.Atoi("53a")
	isError(err)
}
