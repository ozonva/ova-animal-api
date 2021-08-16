package domain

import (
	"fmt"
	"log"
)

type AnimalType uint8

const (
	CAT AnimalType = iota
	DOG
	FISH
	MOUSE
)

func (this AnimalType) String() string {
	switch this {
	case CAT:
		return "CAT"
	case DOG:
		return "DOG"
	case FISH:
		return "FISH"
	case MOUSE:
		return "MOUSE"
	default:
		panic(fmt.Sprintf("Unsupported animal type %d", this))
	}
}

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
