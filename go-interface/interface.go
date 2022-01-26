package interface

type AwesomeVehicle interface {
	popASickWheelie()
}

type GrandmasScooter struct {
	modelNumber string,
	age int
}

func (gs GrandmasScooter) popASickWheelie() {
	fmt.Printf("Why does a %i year old look so fly right now?\n", gs.age);
}
