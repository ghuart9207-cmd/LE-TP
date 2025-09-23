package main

import (
	"fmt"
	"strings"
)

type Voiture struct {
	Marque    string
	nom       string
	Puissance int
	Couleur   string
	coffre    []objet
}

type TrunkItem struct {
	name     string
	quantity int
}

type objet struct {
	nom      string
	quantité int
}

func (c Car) displayCar() {
	fmt.Println("===== Mon véhicule =====")
	fmt.Printf("\tModèle : %s\n", c.name)
	fmt.Printf("\tMarque : %s\n", c.brand)
	fmt.Printf("\tPuissance : %dCH\n", c.power)
	fmt.Printf("\tCouleur : %s\n", c.color)
	fmt.Println("=== Coffre du véhicule ===")
	if len(c.trunk) < 1 {
		fmt.Printf("\tCoffre du véhicule vide...\n")
		return
	}
	for _, item := range c.trunk {
		fmt.Printf("\t- %s x%d\n", item.name, item.quantity)
	}
}

func (c Car) changeColor(color string) {
	if strings.EqualFold(c.color, color) {
		fmt.Println("Oupss le véhicule est déjà de cette couleur....")
	} else {
		c.color = strings.ToLower(color)
		fmt.Printf("La couleur du véhicule à changé : %s\n", c.color)
	}
}

func (cCar) addItemToTrunck(item TrunkItem) {
	isFind := false
	for i := 0; i < len(c.trunk); i++ {
		if strings.EqualFold(c.trunk[i].name, item.name) {
			c.trunk[i].quantity += item.quantity
			isFind = true
			break
		}
	}
	if !isFind {
		item.name = strings.ToLower(item.name)
		c.trunk = append(c.trunk, item)
	}
	fmt.Println(fmt.Sprintf("+%d %s ajouté dans le coffre du véhicule", item.quantity, item.name))
}

func main() {
	item01 := TrunkItem{"sac de course", 1}
	item02 := TrunkItem{"triangle de signalisation", 1}

	car := Car{"RENAULT", "Clio III", "black", 130, []TrunkItem{item01, item02}}

	car.displayCar()
	car.changeColor("black")
	car.changeColor("red")
	car.addItemToTrunck(item01)
	car.addItemToTrunck(TrunkItem{"Sac de course", 5})
	car.addItemToTrunck(TrunkItem{"ordinateur portable", 1})
	car.displayCar()
}
