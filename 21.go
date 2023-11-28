// Реализовать паттерн «адаптер» на любом примере.


package main



import (
	"fmt"
)


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