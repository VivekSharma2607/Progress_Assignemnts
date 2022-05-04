package main

import "fmt"

type Anime struct {
	genre string
}

func (a Anime) getType() string {
	return a.genre
}

type NewGen struct{
	Anime
}

type OldGen struct{
	Anime
}

func main() {
	a1 := OldGen {
		Anime {
			genre: "Shonen",
		},
	}
	a2 := NewGen {
		Anime {
			genre: "Action",
		},
	}
	fmt.Println("The anime from structure Old Generation = " , a1)
	
	fmt.Println("The anime from structure New Generation = " , a2)
}