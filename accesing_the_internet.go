package main

import ("fmt"
		"net/http"
		"io/ioutil")

func main() {
	// Get HTML
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// Convert to String and Print
	string_body := string(bytes)
	fmt.Println(string_body)
	// Close Connection
	resp.Body.Close() 
}