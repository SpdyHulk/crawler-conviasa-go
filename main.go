package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

const (
	//URL Endpoint path
	URL = "http://portal.conviasa.aero/"
)

func main() {

	fmt.Println("====== Crawler Conviasa By Pablo Cariel ======")

	//New entrada
	reader := bufio.NewReader(os.Stdin)
	//New Browser
	browser := surf.NewBrowser()
	//Each Origen
	for k, v := range Origen {
		fmt.Printf("Codigo: %v  -  Ciudad: %v\n", k, v)
	}
	fmt.Print("Por favor, selecione el Origen: ")
	//Leer entrada
	origen, _ := reader.ReadString('\n')
	//Formatear entrada, eliminar espacios
	origen = strings.TrimSpace(origen)

	browser.PostForm("https://wftc1.e-travel.com/plnext/Conviasa/Override.action", FormData)

	browser.Find("form").Each(func(_ int, s *goquery.Selection) {
		name, ok := s.Attr("name")
		if ok {
			if name == "form_select_fare" {
				fmt.Println("Hay Disponibilidad, verifique la pagina")
				return
			}
		}
	})
	fmt.Println("No hay disponibilidad")
}
