package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"<p>Hey there</p>")
	fmt.Fprintf(w,"<h1>OMG</h1>")
	fmt.Fprintf(w,"<p>You %s even add %s</p>","can","<strong>variables</strong>")

	fmt.Fprintf(w, `
		<p>Multiline HTML!</p>
		<h1>OMG 2!</h1>
		`)
}
 
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}