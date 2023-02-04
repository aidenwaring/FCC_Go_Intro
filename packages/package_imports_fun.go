package packages

import (
	"fmt"
	"github.com/aidenwaring/FCC_Go_Intro/packages/castle"
	"github.com/aidenwaring/FCC_Go_Intro/packages/river"
)

func PackageGreetings() {
	var myNum int = 27
	fmt.Println(myNum)

	castle.KnightGreeting()
	river.FishGreeting()
}
