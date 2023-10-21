// Реализовать паттерн «адаптер» на любом примере.


package main



import (
	"fmt"
)

// OldSystem имеет метод RequestOld.
type OldSystem struct{}

func (os *OldSystem) RequestOld() string {
	return "Старая система делает запрос"
}

// NewSystem интерфейс с методом RequestNew.
type NewSystem interface {
	RequestNew() string
}

// OldSystemAdapter адаптирует OldSystem к интерфейсу NewSystem.
type OldSystemAdapter struct {
	OldSystem
}

func (a *OldSystemAdapter) RequestNew() string {
	return a.RequestOld()
}

// Новая система реализует интерфейс NewSystem.
type NewSystemImpl struct{}

func (ns *NewSystemImpl) RequestNew() string {
	return "Новая система делает новый запрос"
}

func Exercise_21() {
	newSystem := &NewSystemImpl{}
	oldSystemAdapter := &OldSystemAdapter{}

	fmt.Println(newSystem.RequestNew()) // Используем новую систему
	fmt.Println(oldSystemAdapter.RequestNew()) // Используем адаптированную старую систему
}