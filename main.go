package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	// convertjson "github.com/ethic-quiz-maze-generator/convertJson"
)

type Fact struct {
	Statement string
	Correct   bool
}

type WrongFact struct {
	Fact Fact
	used bool
}

type ConnectedFact struct {
	OwnFact       Fact
	ConnectedFact Fact
	Path          bool
}

type FactCorrect []Fact

func main() {
	// convertjson.ConvertJson()
	data, _ := os.ReadFile("blarg.json")
	var facts []Fact
	json.Unmarshal(data, &facts)
	leftFacts := createFactIntList(len(facts))
	factCorrectList := []int{}
	factCorrectList, leftFacts = connectFact(factCorrectList, leftFacts)
	fmt.Print(factCorrectList)
	fmt.Print(createWrongFactList(facts))
}

func connectFact(factCorrectList []int, leftFacts []int) ([]int, []int) {
	factItemInt, factInt := getRandomFact(leftFacts)
	factCorrectList = append(factCorrectList, factItemInt)
	leftFacts = remove(leftFacts, factInt)
	factsLen := len(leftFacts)
	if factsLen > 0 {
		return connectFact(factCorrectList, leftFacts)
	}
	return factCorrectList, leftFacts
}

func remove(slice []int, i int) []int {
	sliceLen := len(slice)
	if i < sliceLen {
		slice[i] = slice[sliceLen-1]
	}
	return slice[:sliceLen-1]
}

func getRandomFact(facts []int) (int, int) {
	len := len(facts)
	randInt := rand.Intn(len)
	return facts[randInt], randInt
}

func createFactIntList(len int) []int {
	factList := make([]int, len)
	for i := 0; i < len; i++ {
		factList[i] = i
	}
	return factList
}
