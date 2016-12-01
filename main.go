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
	fmt.Println("===========================================")

	if _, ok := Origen[origen]; ok {
		for k, v := range Destino {
			fmt.Printf("Codigo: %v  -  Ciudad: %v\n", k, v)
		}
		fmt.Print("Por favor, selecione el Destino: ")
		destino, _ := reader.ReadString('\n')
		destino = strings.TrimSpace(destino)
		if _, ok := Destino[destino]; ok {
			fmt.Print("Por favor, ingrese fecha de Salida: ")
			fecha_salida, _ := reader.ReadString('\n')
			fecha_salida = strings.TrimSpace(fecha_salida)
			fmt.Print("Por favor, ingrese fecha de Regreso: ")
			fecha_regreso, _ := reader.ReadString('\n')
			fecha_regreso = strings.TrimSpace(fecha_regreso)
			FormData.Add("B_LOCATION_1", origen)
			FormData.Add("E_LOCATION_1", destino)
			FormData.Add("B_LOCATION_2", destino)
			FormData.Add("E_LOCATION_2", origen)
			FormData.Add("ORIGEN", origen)
			FormData.Add("DESTINO", destino)
			FormData.Add("B_DATE_1", fecha_salida+"1200")
			FormData.Add("B_DATE_2", fecha_regreso+"0000")
			//Solicitud Form Post
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
		}
	}
}
