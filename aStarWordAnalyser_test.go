package wordPathAnalyser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

//Test for the main function in AstarWordAnalyser
func TestAStarAnalyseFile(t *testing.T) {
	fmt.Println("Testing Main Analyse File method: 'AStarAnalyseFile'....")

	//Arrange

	testInputs := readTestFile("../Answers.txt", " -> ")

	for _, input := range testInputs {
		input.FileLocation = "../WordList.txt"
		input.Delimiter = ""
		input.PathFound = true
	}

	customInput1 := aStarAnalyseMockInput{StartWord: "test", EndWord: "most", FileLocation: "./testInput.txt", Delimiter: "", PathFound: true, ResultPathLength: 4}
	customInput2 := aStarAnalyseMockInput{StartWord: "pest", EndWord: "post", FileLocation: "./testInput.txt", Delimiter: "", PathFound: true, ResultPathLength: 2}
	customInput3 := aStarAnalyseMockInput{StartWord: "test", EndWord: "fail", FileLocation: "./testInput.txt", Delimiter: "", PathFound: false, ResultPathLength: 0}
	customInput4 := aStarAnalyseMockInput{StartWord: "test", EndWord: "most", FileLocation: "./testInputDelimited.txt", Delimiter: ",", PathFound: true, ResultPathLength: 4}
	customInput5 := aStarAnalyseMockInput{StartWord: "pest", EndWord: "post", FileLocation: "./testInputDelimited.txt", Delimiter: ",", PathFound: true, ResultPathLength: 2}
	customInput6 := aStarAnalyseMockInput{StartWord: "test", EndWord: "fail", FileLocation: "./testInputDelimited.txt", Delimiter: ",", PathFound: false, ResultPathLength: 0}

	testInputs = append(testInputs, &customInput1)
	testInputs = append(testInputs, &customInput2)
	testInputs = append(testInputs, &customInput3)
	testInputs = append(testInputs, &customInput4)
	testInputs = append(testInputs, &customInput5)
	testInputs = append(testInputs, &customInput6)

	totalTests := len(testInputs)

	var waitGroup sync.WaitGroup
	waitGroup.Add(totalTests)
	testResultOutput := make(chan string, 100)

	for i, input := range testInputs {
		go runAstarTestConcurrently(input, i+1, totalTests, &waitGroup, testResultOutput)
	}
	go dealWithTestOutputConcurrelty(testResultOutput, t)

	waitGroup.Wait()
	close(testResultOutput)
	fmt.Print("\n")
}

func runAstarTestConcurrently(testInput *aStarAnalyseMockInput, testNumber, totalTests int, wg *sync.WaitGroup, outputChannel chan string) {
	defer wg.Done()
	//Act
	pathFound, resultPath := AStarAnalyseFile(testInput.StartWord, testInput.EndWord, testInput.FileLocation, testInput.Delimiter)

	//Assert
	if pathFound != testInput.PathFound || testInput.ResultPathLength != len(resultPath) {
		outputText := fmt.Sprint("Test number ", testNumber, "\n", "Given the inputs:\n",
			"start word = ", testInput.StartWord, "\n",
			"end word = ", testInput.EndWord, "\n",
			"file location = ", testInput.FileLocation, "\n",
			"delimiter = ", testInput.Delimiter, "\n",
			"Expected results to be:\n",
			"Path Found = ", testInput.PathFound, "\n",
			"Result Path = ", testInput.ResultPathLength, "\n",
			"Actual results were:\n",
			"Path Found = ", pathFound, "\n",
			"Result Path = ", resultPath, "\n")
		outputChannel <- outputText
		fmt.Println("Test ", testNumber, " of ", totalTests, "- failed.")
	} else {
		fmt.Println("Test ", testNumber, " of ", totalTests, " - passed.")
	}
}

func dealWithTestOutputConcurrelty(inputChannel chan string, t *testing.T) {
	for inputString := range inputChannel {
		t.Error(inputString)
	}
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
				"Test number ", i+1, "\n",
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
				"Test number ", i+1, "\n",
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
	resultDict1 := []*aStarWordNode{&wordNodeBeat, &wordNodeBrat, &wordNodeBrag}

	inputNode2 := &wordNodeBest
	inputDict2 := []*aStarWordNode{&wordNodeTest, &wordNodePest, &wordNodeBeat, &wordNodeBrat, &wordNodeBrag}
	resultChildren2 := []*aStarWordNode{&wordNodeTest, &wordNodePest, &wordNodeBeat}
	resultDict2 := []*aStarWordNode{&wordNodeBrat, &wordNodeBrag}

	inputNode3 := &wordNodeBrag
	inputDict3 := []*aStarWordNode{&wordNodeTest, &wordNodeBest, &wordNodePest, &wordNodeBeat, &wordNodeBrat}
	resultChildren3 := []*aStarWordNode{&wordNodeBrat}
	resultDict3 := []*aStarWordNode{&wordNodeTest, &wordNodeBest, &wordNodePest, &wordNodeBeat}

	testInputs := []aStarGenerateNodeChildren{
		{InputNode: inputNode1,
			InputDictionary:     inputDict1,
			ResultChildrenNodes: resultChildren1,
			ResultDictionary:    resultDict1},
		{InputNode: inputNode2,
			InputDictionary:     inputDict2,
			ResultChildrenNodes: resultChildren2,
			ResultDictionary:    resultDict2},
		{InputNode: inputNode3,
			InputDictionary:     inputDict3,
			ResultChildrenNodes: resultChildren3,
			ResultDictionary:    resultDict3},
	}

	for i, input := range testInputs {
		fmt.Print("Test ", i+1, " of ", len(testInputs))
		//Act
		resultChildren, resultDictionary := generateNodeChildren(input.InputNode, input.InputDictionary)

		//Assert
		if !doNodePointerArraysMatch(input.ResultChildrenNodes, resultChildren) || !doNodePointerArraysMatch(input.ResultDictionary, resultDictionary) {
			t.Error(
				"Test number ", i+1, "\n",
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
				"Test number ", i+1, "\n",
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

func readTestFile(fileLocation, delimiter string) []*aStarAnalyseMockInput {
	//Open the file and log an error if there is one.
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	//Defer file.close to the end of this function.
	defer file.Close()

	//Array to store wordNodes.
	result := make([]*aStarAnalyseMockInput, 0)

	//create scanner for the file opened.
	scanner := bufio.NewScanner(file)
	//While there are still lines in the file:
	for scanner.Scan() {
		//Check the text is not the start or end word (these are dealt with seperately) and if not then create word node and add it to the array.
		words := strings.Split(scanner.Text(), delimiter)
		var lineInput aStarAnalyseMockInput
		lineInput.StartWord = words[0]
		lineInput.EndWord = words[len(words)-1]
		lineInput.ResultPathLength = len(words)
		result = append(result, &lineInput)
	}

	return result
}
