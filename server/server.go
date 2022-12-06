package server

/*
Per compilare (dovete avere installato git):

  go mod init server
  go mod tidy

(solo una volta)
*/

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"net/http"

	"github.com/holizz/terrapin"
)

func pippoFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
    <!doctype html>
    <title>Pagina di saluto</title>

    <h1>Titolo</h1>

    <p>Ciao, pinocchietto! Bentornato!

    <p><img src="https://www.napolike.it/wp-content/uploads/2022/09/Giorgia-Meloni-e1664186154697.png" style="width: 55vw; min-width: 330px;">
    `))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
    <!doctype html>
    <title>Buon Natale</title>

    <h1>Buon Natale</h1>

    <p>Buon Natale, pinocchietto! Bentornato!

    <p><img src="/mainImage.png">
    `))
}

func mainImage(w http.ResponseWriter, r *http.Request) {
	campo := image.NewRGBA(image.Rect(0, 0, 600, 400))
	t := terrapin.NewTerrapin(campo, terrapin.Position{250.0, 50.0})
	kochSnowflake(t, 4, 4)

	png.Encode(w, campo)
}

func kochSnowflake(t *terrapin.Terrapin, lung float64, liv int) {
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, lung, liv)
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, lung, liv)
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, lung, liv)
}

func kochCurve(t *terrapin.Terrapin, lung float64, liv int) {
	if liv == 0 {
		t.Forward(lung)
		return
	}
	kochCurve(t, lung, liv-1)
	t.Left(math.Pi / 3.0)
	kochCurve(t, lung, liv-1)
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, lung, liv-1)
	t.Left(math.Pi / 3.0)
	kochCurve(t, lung, liv-1)
}

func main() {
	fmt.Println("Listening on http://localhost:3000/")

	http.HandleFunc("/pippo", pippoFunc)
	http.HandleFunc("/pluto", pippoFunc)

	http.HandleFunc("/main", mainPage)
	http.HandleFunc("/mainImage.png", mainImage)

	http.ListenAndServe(":3000", nil)
}
