package main

/*
slice é um array de tamanho variável.
"type" neste caso é qualquer tipo de dado que o array tem.
[5]type == array 
[5 5] == multidimensional array
[]int == slice (sem tamanho fixo)
*/

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

type SitemapIndex struct {
	Locations []Location  `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close() 

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}

