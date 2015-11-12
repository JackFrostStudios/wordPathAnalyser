package wordPathAnalyser

import (
	"fmt"
	"testing"
)

//Struct used to hold the mocked input.
type aStarAnalyseMockInput struct {
	startWord, endWord, fileLocation, delimiter string
	pathFound                                   bool
	resultPath                                  []string
}
type aStarReadFileMockInput struct {
	startWord, endWord, fileLocation, delimiter string
	resultList                                  []aStarWordNode
}

//Test for the main function in AstarWordAnalyser
func TestAStarAnalyseFile(t *testing.T) {
	fmt.Println("Testing Main Analyse File method: 'AStarAnalyseFile'....")

	//Arrange
	testInputs := []aStarAnalyseMockInput{
		{"test", "most", "./testInput.txt", "", true, []string{"most", "post", "pest", "test"}},
		{"pest", "post", "./testInput.txt", "", true, []string{"post", "pest"}},
		{"test", "fail", "./testInput.txt", "", false, []string{}},
		{"test", "most", "./testInputDelimited.txt", ",", true, []string{"most", "post", "pest", "test"}},
		{"pest", "post", "./testInputDelimited.txt", ",", true, []string{"post", "pest"}},
		{"test", "fail", "./testInputDelimited.txt", ",", false, []string{}},
	}

	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		pathFound, resultPath := AStarAnalyseFile(input.startWord, input.endWord, input.fileLocation, input.delimiter)

		//Assert
		if pathFound != input.pathFound || !doArraysMatch(input.resultPath, resultPath) {
			t.Error(
				"Given the inputs:\n",
				"start word = ", input.startWord, "\n",
				"end word = ", input.endWord, "\n",
				" file location = ", input.fileLocation, "\n",
				" delimiter = ", input.delimiter, "\n",
				"Expected results to be:\n",
				"Path Found = ", input.pathFound, "\n",
				"Result Path = ", input.resultPath, "\n",
				"Actual results were:\n",
				"Path Found = ", pathFound, "\n",
				"Result Path = ", resultPath, "\n",
			)
			fmt.Println("- failed.")
		} else {
			fmt.Println(" - success.")
		}
	}
}

func TestReadFile(t *testing.T) {
	fmt.Println("Testing read in file method: 'readFile'....")

	//Arrange
	testInputs := []aStarReadFileMockInput{
		{"test", "most", "./testInput.txt", "",
			[]aStarWordNode{
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "pest"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "post"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "fail"}}},

		{"pest", "post", "./testInput.txt", "",
			[]aStarWordNode{
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "test"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "most"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "fail"}}},
		{"test", "most", "./testInputDelimited.txt", ",",
			[]aStarWordNode{
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "pest"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "post"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "fail"}}},

		{"pest", "post", "./testInputDelimited.txt", ",",
			[]aStarWordNode{
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "test"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "most"},
				{fScore: 0,
					gScore: 0,
					hScore: 0,
					word:   "fail"}}},
	}

	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		resultList := readFile(input.startWord, input.endWord, input.fileLocation, input.delimiter)

		//Assert
		if !doNodeArraysMatch(input.resultList, resultList) {
			t.Error(
				"Given the inputs:\n",
				"start word = ", input.startWord, "\n",
				"end word = ", input.endWord, "\n",
				" file location = ", input.fileLocation, "\n",
				" delimiter = ", input.delimiter, "\n",
				"Expected results to be:\n",
				"Result List = ", input.resultList, "\n",
				"Actual results were:\n",
				"Result List = ", convertNodePointersToNodes(resultList), "\n",
			)
			fmt.Println(" - failed.")
		} else {
			fmt.Println(" - passed.")
		}
	}
}

func doArraysMatch(expected, actual []string) (inputsMatch bool) {
	if len(expected) != len(actual) {
		inputsMatch = false
		return
	}
	for i, word := range expected {
		if word != actual[i] {
			inputsMatch = false
			return
		}
	}
	inputsMatch = true
	return
}

func doNodeArraysMatch(expected []aStarWordNode, actual []*aStarWordNode) (inputsMatch bool) {
	if len(expected) != len(actual) {
		inputsMatch = false
		return
	}
	for i, node := range expected {
		if node.fScore != actual[i].fScore ||
			node.gScore != actual[i].gScore ||
			node.hScore != actual[i].hScore ||
			node.parentNode != actual[i].parentNode ||
			node.word != actual[i].word {
			inputsMatch = false
			return
		}
	}
	inputsMatch = true
	return
}

func convertNodePointersToNodes(nodes []*aStarWordNode) (outputNodes []aStarWordNode) {
	for _, node := range nodes {
		currentNode := newAStarWordNode(node.word)
		outputNodes = append(outputNodes, currentNode)
	}
	return outputNodes
}
