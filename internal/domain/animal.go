package domain

import (
	"fmt"
	"log"
)

//go:generate enumer -type=AnimalType -json
type AnimalType uint8

const (
	CAT AnimalType = iota
	DOG
	FISH
	MOUSE
)

type Animal struct {
	Id     uint64
	UserId uint64
	Name   string
	Type   AnimalType
}

func (this Animal) Say() {
	switch this.Type {
	case CAT:
		fmt.Println("Meow!")
	case DOG:
		fmt.Println("Bow!")
	case MOUSE:
		fmt.Println("Beep!")
	default:
		log.Panicf("Don't know, how a %s tells!", this.Type)
	}
}

func (this Animal) String() string {
	return fmt.Sprintf(
		"%d %d %s %s",
		this.Id,
		this.UserId,
		this.Type,
		this.Name,
	)
}

func (this *Animal) Rename(newName string) {
	this.Name = newName
}

func (this Animal) IncorrectRename(newName string) {
	this.Name = newName
}
