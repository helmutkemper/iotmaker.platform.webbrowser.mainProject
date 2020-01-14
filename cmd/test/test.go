package main

import (
	"fmt"
	"math"
	"math/rand"
)

const kUIdSize = 5

type funcList struct {
	id string
	f  string
}

var funcListToDraw []funcList
var kUIdCharList []string

func main() {
	var zero = AddToDraw("zero")
	var um = AddToDraw("um")
	var dois = AddToDraw("dois")
	var tres = AddToDraw("tres")
	var quatro = AddToDraw("quatro")

	fmt.Printf("0: %v\n", zero)
	fmt.Printf("1: %v\n", um)
	fmt.Printf("2: %v\n", dois)
	fmt.Printf("3: %v\n", tres)
	fmt.Printf("4: %v\n", quatro)

	ToFrontOneLevel(um)

	fmt.Printf("lista: %v\n", funcListToDraw)
}

func init() {
	kUIdCharList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_", "!", "@",
		"#", "$", "%", "&", "*", "(", ")", "-", "_", "+", "=", "[", "{", "}", "]", "/", "?", ";", ":", ".", ",", "<", ">",
		"|"}
}

func getUId() string {
	var UId = ""
	for i := 0; i != kUIdSize; i += 1 {
		UId += kUIdCharList[rand.Intn(len(kUIdCharList)-1)]
	}

	return UId
}

func AddToDraw(runnerFunc string) string {
	UId := getUId()
	funcListToDraw = append(funcListToDraw, funcList{id: UId, f: runnerFunc})

	return UId
}

func DeleteFromDraw(UId string) {
	for k, runner := range funcListToDraw {
		if runner.id == UId {
			funcListToDraw = append(funcListToDraw[:k], funcListToDraw[k+1:]...)
			break
		}
	}
}

func SetZIndex(UId string, index int) int {
	var function funcList
	var pass = false
	var length = len(funcListToDraw)

	if index < 0 || index > length-1 {
		return math.MaxInt32
	}

	for k, runner := range funcListToDraw {
		if runner.id == UId {
			pass = true
			function = runner
			funcListToDraw = append(funcListToDraw[:k], funcListToDraw[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	if index == 0 {

		funcListToDraw = append([]funcList{function}, funcListToDraw...)

	} else if index == length-1 {

		funcListToDraw = append(funcListToDraw, function)

	} else {

		firstPart := make([]funcList, len(funcListToDraw[:index]))
		lastPart := make([]funcList, len(funcListToDraw[index:]))

		copy(firstPart, funcListToDraw[:index])
		copy(lastPart, funcListToDraw[index:])

		firstPart = append(firstPart, function)

		funcListToDraw = make([]funcList, 0)
		funcListToDraw = append(firstPart, lastPart...)
	}

	return index
}

func ToFront(UId string) int {
	var function funcList
	var pass = false
	for k, runner := range funcListToDraw {
		if runner.id == UId {
			pass = true
			function = runner
			funcListToDraw = append(funcListToDraw[:k], funcListToDraw[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	funcListToDraw = append(funcListToDraw, function)

	return 0
}

func ToBack(UId string) int {
	var function funcList
	var pass = false
	for k, runner := range funcListToDraw {
		if runner.id == UId {
			pass = true
			function = runner
			funcListToDraw = append(funcListToDraw[:k], funcListToDraw[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	funcListToDraw = append([]funcList{function}, funcListToDraw...)

	return len(funcListToDraw) - 1
}
