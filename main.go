package main

import (
	"fmt"
	route2 "github.com/Luvicitousr/project-simulator/application/route"
)

func main()  {
	route := route2.Route{
		ID: 		"1",
		ClientID: 	"1",
	}
	route.LoadPositions()
	stringjson := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
	
}