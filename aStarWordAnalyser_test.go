package wordPathAnalyser

import (
	"fmt"
	"testing"
)

//Test for the main function in AstarWordAnalyser
func TestAStarAnalyseFile(t *testing.T) {
	fmt.Println("Testing Main Analyse File method: 'AStarAnalyseFile'....")

	//Arrange
	testInputs := []aStarAnalyseMockInput{
		{StartWord: "test", EndWord: "most", FileLocation: "./testInput.txt", Delimiter: "", PathFound: true, ResultPath: []string{"most", "post", "pest", "test"}},
		{StartWord: "pest", EndWord: "post", FileLocation: "./testInput.txt", Delimiter: "", PathFound: true, ResultPath: []string{"post", "pest"}},
		{StartWord: "test", EndWord: "fail", FileLocation: "./testInput.txt", Delimiter: "", PathFound: false, ResultPath: []string{}},
		{StartWord: "test", EndWord: "most", FileLocation: "./testInputDelimited.txt", Delimiter: ",", PathFound: true, ResultPath: []string{"most", "post", "pest", "test"}},
		{StartWord: "pest", EndWord: "post", FileLocation: "./testInputDelimited.txt", Delimiter: ",", PathFound: true, ResultPath: []string{"post", "pest"}},
		{StartWord: "test", EndWord: "fail", FileLocation: "./testInputDelimited.txt", Delimiter: ",", PathFound: false, ResultPath: []string{}},
	}

	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		pathFound, resultPath := AStarAnalyseFile(input.StartWord, input.EndWord, input.FileLocation, input.Delimiter)

		//Assert
		if pathFound != input.PathFound || !doArraysMatch(input.ResultPath, resultPath) {
			t.Error(
				"Given the inputs:\n",
				"start word = ", input.StartWord, "\n",
				"end word = ", input.EndWord, "\n",
				" file location = ", input.FileLocation, "\n",
				" delimiter = ", input.Delimiter, "\n",
				"Expected results to be:\n",
				"Path Found = ", input.PathFound, "\n",
				"Result Path = ", input.ResultPath, "\n",
				"Actual results were:\n",
				"Path Found = ", pathFound, "\n",
				"Result Path = ", resultPath, "\n",
			)
			fmt.Println("- failed.")
		} else {
			fmt.Println(" - success.")
		}
	}
	fmt.Print("\n")
}

//Test the read file function will return an array of the word nodes for a given file.
func TestReadFile(t *testing.T) {
	fmt.Println("Testing read in file method: 'readFile'....")

	//Arrange
	wordNodePest := newAStarWordNode("pest")
	wordNodePost := newAStarWordNode("post")
	wordNodeFail := newAStarWordNode("fail")
	wordNodeTest := newAStarWordNode("test")
	wordNodeMost := newAStarWordNode("most")
	testInputs := []aStarReadFileMockInput{
		{StartWord: "test",
			EndWord:      "most",
			FileLocation: "./testInput.txt",
			Delimiter:    "",
			ResultList:   []*aStarWordNode{&wordNodePest, &wordNodePost, &wordNodeFail}},

		{StartWord: "pest",
			EndWord:      "post",
			FileLocation: "./testInput.txt",
			Delimiter:    "",
			ResultList:   []*aStarWordNode{&wordNodeTest, &wordNodeMost, &wordNodeFail}},
		{StartWord: "test",
			EndWord:      "most",
			FileLocation: "./testInputDelimited.txt",
			Delimiter:    ",",
			ResultList:   []*aStarWordNode{&wordNodePest, &wordNodePost, &wordNodeFail}},

		{StartWord: "pest",
			EndWord:      "post",
			FileLocation: "./testInputDelimited.txt",
			Delimiter:    ",",
			ResultList:   []*aStarWordNode{&wordNodeTest, &wordNodeMost, &wordNodeFail}},
	}

	//Loop through all test cases.
	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		resultList := readFile(input.StartWord, input.EndWord, input.FileLocation, input.Delimiter)

		//Assert
		if !doNodePointerArraysMatchOnValue(input.ResultList, resultList) {
			t.Error(
				"Given the inputs:\n",
				"start word = ", input.StartWord, "\n",
				"end word = ", input.EndWord, "\n",
				" file location = ", input.FileLocation, "\n",
				" delimiter = ", input.Delimiter, "\n",
				"Expected results to be:\n",
				"Result List = ", convertNodePointersToNodes(input.ResultList), "\n",
				"Actual results were:\n",
				"Result List = ", convertNodePointersToNodes(resultList), "\n",
			)
			fmt.Println(" - failed.")
		} else {
			fmt.Println(" - passed.")
		}
	}
	fmt.Print("\n")
}

//Test that the calculate node function will return the correct cost for two given words.
func TestCalculateNodeCost(t *testing.T) {
	fmt.Println("Testing node cost calculation method: 'calculateNodeCost'....")

	//Arrange
	testInputs := []aStarCalculateNodeCostMockInput{
		{StartWord: "test",
			EndWord: "test",
			Result:  0},
		{StartWord: "test",
			EndWord: "best",
			Result:  1},
		{StartWord: "test",
			EndWord: "beat",
			Result:  2},
		{StartWord: "test",
			EndWord: "brat",
			Result:  3},
		{StartWord: "test",
			EndWord: "brag",
			Result:  4},
	}

	//Loop through all test cases
	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		result := calculateNodeCost(input.StartWord, input.EndWord)

		//Assert
		if input.Result != result {
			t.Error(
				"Given the inputs:\n",
				"start word = ", input.StartWord, "\n",
				"end word = ", input.EndWord, "\n",
				"Expected result to be:\n",
				"Result = ", input.Result, "\n",
				"Actual result was:\n",
				"Result = ", result, "\n",
			)
			fmt.Println(" - failed.")
		} else {
			fmt.Println(" - passed.")
		}
	}
	fmt.Print("\n")
}

//Test the generate node children function will return the correct result for a given input
func TestGenerateNodeChildren(t *testing.T) {
	fmt.Println("Testing generating children nodes method: 'generateNodeChildren'....")

	//Arrange
	wordNodeTest := newAStarWordNode("test")
	wordNodePest := newAStarWordNode("pest")
	wordNodeBest := newAStarWordNode("best")
	wordNodeBeat := newAStarWordNode("beat")
	wordNodeBrat := newAStarWordNode("brat")
	wordNodeBrag := newAStarWordNode("brag")

	inputNode1 := &wordNodeTest
	inputDict1 := []*aStarWordNode{&wordNodePest, &wordNodeBest, &wordNodeBeat, &wordNodeBrat, &wordNodeBrag}
	resultChildren1 := []*aStarWordNode{&wordNodePest, &wordNodeBest}

	inputNode2 := &wordNodeBest
	inputDict2 := []*aStarWordNode{&wordNodeTest, &wordNodePest, &wordNodeBeat, &wordNodeBrat, &wordNodeBrag}
	resultChildren2 := []*aStarWordNode{&wordNodeTest, &wordNodePest, &wordNodeBeat}

	inputNode3 := &wordNodeBrag
	inputDict3 := []*aStarWordNode{&wordNodeTest, &wordNodeBest, &wordNodePest, &wordNodeBeat, &wordNodeBrat}
	resultChildren3 := []*aStarWordNode{&wordNodeBrat}
	testInputs := []aStarGenerateNodeChildren{
		{InputNode: inputNode1,
			InputDictionary:     inputDict1,
			ResultChildrenNodes: resultChildren1},
		{InputNode: inputNode2,
			InputDictionary:     inputDict2,
			ResultChildrenNodes: resultChildren2},
		{InputNode: inputNode3,
			InputDictionary:     inputDict3,
			ResultChildrenNodes: resultChildren3},
	}

	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		resultChildren := generateNodeChildren(input.InputNode, input.InputDictionary)

		//Assert
		if !doNodePointerArraysMatch(input.ResultChildrenNodes, resultChildren) {
			t.Error(
				"Given the inputs:\n",
				"start node = ", input.InputNode.Word, "\n",
				"Input Dictionary = ", convertNodePointersToNodes(input.InputDictionary), "\n",
				"Expected result to be:\n",
				"Result Children = ", convertNodePointersToNodes(input.ResultChildrenNodes), "\n",
				"Actual result was:\n",
				"Result Children = ", convertNodePointersToNodes(resultChildren), "\n",
			)
			fmt.Println(" - failed.")
		} else {
			fmt.Println(" - passed.")
		}
	}
	fmt.Print("\n")
}

func TestGetResultPath(t *testing.T) {
	fmt.Println("Testing get result path method: 'getResultPath'....")
	//Arrange
	wordNodeTest := newAStarWordNode("test")
	wordNodeBest := newAStarWordNode("best")
	wordNodeBest.ParentNode = &wordNodeTest
	wordNodeBeat := newAStarWordNode("beat")
	wordNodeBeat.ParentNode = &wordNodeBest
	wordNodeBrat := newAStarWordNode("brat")
	wordNodeBrat.ParentNode = &wordNodeBeat
	wordNodeBrag := newAStarWordNode("brag")
	wordNodeBrag.ParentNode = &wordNodeBrat

	testInputs := []aStarGetResultPath{
		{EndNode: wordNodeBrag,
			ResultList: []string{"brag", "brat", "beat", "best", "test"}},
		{EndNode: wordNodeBeat,
			ResultList: []string{"beat", "best", "test"}},
		{EndNode: wordNodeTest,
			ResultList: []string{"test"}},
	}

	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		result := getResultPath(input.EndNode)

		//Assert
		if !doArraysMatch(input.ResultList, result) {
			t.Error(
				"Given the inputs:\n",
				"End node = ", input.EndNode.Word, "\n",
				"Expected result to be:\n",
				"Result Path = ", input.ResultList, "\n",
				"Actual result was:\n",
				"Result Path = ", result, "\n",
			)
			fmt.Println(" - failed.")
		} else {
			fmt.Println(" - passed.")
		}
	}
	fmt.Print("\n")
}

//-----------INTERNAL FUNCTIONS-----------\\
//Check if the results from AStarAnalyseFile match those expected.
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

//Check if the results from readFile match those expected.
func doNodePointerArraysMatchOnValue(expected []*aStarWordNode, actual []*aStarWordNode) (inputsMatch bool) {
	if len(expected) != len(actual) {
		inputsMatch = false
		return
	}
	for i, node := range expected {
		if node.FScore != actual[i].FScore ||
			node.GScore != actual[i].GScore ||
			node.HScore != actual[i].HScore ||
			node.ParentNode != actual[i].ParentNode ||
			node.Word != actual[i].Word {
			inputsMatch = false
			return
		}
	}
	inputsMatch = true
	return
}

//Check if the results from readFile match those expected.
func doNodePointerArraysMatch(expected []*aStarWordNode, actual []*aStarWordNode) (inputsMatch bool) {
	if len(expected) != len(actual) {
		inputsMatch = false
		return
	}
	for i, node := range expected {
		if node != actual[i] {
			inputsMatch = false
			return
		}
	}
	inputsMatch = true
	return
}

//Used to convert the array of pointers returned by the readFile function to actual objects so they can be printed.
func convertNodePointersToNodes(nodes []*aStarWordNode) (outputNodes []aStarWordNode) {
	for _, node := range nodes {
		currentNode := newAStarWordNode(node.Word)
		outputNodes = append(outputNodes, currentNode)
	}
	return outputNodes
}
