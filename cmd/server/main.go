package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"go-academy/pkg/calculator"
)

func apiCalculateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<form method="post">
		<input name="equation" required> <input type="submit" value="Calculate">
		</form>
		{{ if .Equation }}<h1>{{ .Equation }}</h1>{{ end }}`

	// set the encoding
	w.Header().Add("Content-type", "text/html")

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	equation := ""
	if r.FormValue("equation") != "" {
		equation = calculator.Calculate(r.FormValue("equation"))
	}


	data := struct {
		Equation string
	}{
		Equation: equation,
	}

	// parse the template
	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		fmt.Println("Failed to parse template;", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func startServer(address string) {
	http.HandleFunc("/api/calculate", apiCalculateHandler)

	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, nil)
}

func main() {
	var addr = flag.String("addr", "", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	startServer(*addr)
}