// Реализовать паттерн «адаптер» на любом примере.


package main



import (
	"fmt"
)

// OldSystem имеет метод RequestOld.
// type OldSystem struct{}

// func (os *OldSystem) RequestOld() string {
// 	return "Старая система делает запрос"
// }

// // NewSystem интерфейс с методом RequestNew.
// type NewSystem interface {
// 	RequestNew() string
// }

// // OldSystemAdapter адаптирует OldSystem к интерфейсу NewSystem.
// type OldSystemAdapter struct {
// 	OldSystem
// }

// func (a *OldSystemAdapter) RequestNew() string {
// 	return a.RequestOld()
// }

// // Новая система реализует интерфейс NewSystem.
// type NewSystemImpl struct{}

// func (ns *NewSystemImpl) RequestNew() string {
// 	return "Новая система делает новый запрос"
// }

// func main() {
// 	newSystem := &NewSystemImpl{}
// 	oldSystemAdapter := &OldSystemAdapter{}

// 	fmt.Println(newSystem.RequestNew()) // Используем новую систему
// 	fmt.Println(oldSystemAdapter.RequestNew()) // Используем адаптированную старую систему
// }


type Dog struct{}

func (d *Dog) Wof() {
	fmt.Println("Wof")
}
type Cat struct{}

func (c *Cat) Meow() {
	fmt.Println("Meow")
}

type AnimalAdapter interface { //Общий интерфейс
    Reaction()
}

type DogAdapter struct{
    *Dog
}

func (adapter *DogAdapter) Reaction() {
    adapter.Wof()
}

type CatAdapter struct{
    *Cat
}

func (adapter *CatAdapter) Reaction() {// Тут можно было сделать один адаптер и преобразовать Meow в Wof
    adapter.Meow()
}

func NewCatAdapter(cat *Cat) AnimalAdapter { //создаем адаптер для Cat
    return &CatAdapter{cat}
}

func NewDogAdapter(dog *Dog) AnimalAdapter { //создаем адаптер для Dog
    return &DogAdapter{dog}
}

func main() {
	var d Dog
	var c Cat
	new_c := NewCatAdapter(&c)
	new_d := NewDogAdapter(&d)
	new_c.Reaction()
	new_d.Reaction() //Теперь оба объекта реализуют метод Reaction а значит могу реализовывать один интерфейс 
}