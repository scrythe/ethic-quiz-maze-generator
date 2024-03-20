package convertjson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Fact struct {
	Statement string
	Correct   bool
}

func ConvertJson() {
	facts := []Fact{}
	file, err := os.Open("questions.txt")
	if err != nil {
		fmt.Print(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)
		statement := line[0 : lineLen-1]
		correct := string(line[lineLen-1])
		correctBool := correct == "r"
		fact := Fact{Statement: statement, Correct: correctBool}
		facts = append(facts, fact)
	}
	factsJson, _ := json.MarshalIndent(facts, "", "  ")
	fmt.Print(string(factsJson))
	os.WriteFile("blarg.json", factsJson, os.ModePerm)
}
