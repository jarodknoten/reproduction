package producer

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type producer struct {
	Name   string
	Age    int
	Alive  bool
	MaxAge int
	//deathAge    int
	//sex         string
	//orientation string
}

/*New Creates a new producer */
func New(Name string, Age int, Alive bool, MaxAge int) producer {
	p := producer{Name, Age, Alive, MaxAge}
	return p
}

func (p producer) Dump() {
	fmt.Println()
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Alive: %s\n", strconv.FormatBool(p.Alive))
	fmt.Printf("Max Age: %d\n", p.MaxAge)
	fmt.Println()
}

func generateDeathAge() {
	rand.Seed(time.Now().Unix())
}
