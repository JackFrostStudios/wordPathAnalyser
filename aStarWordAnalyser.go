package wordPathAnalyser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//AStarAnalyseFile uses the A* Graphing Algorythm to find the shorted path between two words of the same length when changing one letter at a time.
//It will read in the list of words to be used
//INPUTS: startword, endword, filelocation, delimiter (strings) (**If delimiter is not to be used enter ""**)
//OUTPUT: path found result (Boolean), path from end word to start word ([]string) (if no path is found emtpy array is returned)
func AStarAnalyseFile(sW, eW, fL, dL string) (foundResult bool, resultPath []string) {
	//List of all words that can possibly be used.
	wordDictionary := readFile(sW, eW, fL, dL)
	//List of words that have been assigned a partentNode and are still to be analyzed
	openList := make([]*aStarWordNode, 0)
	//List of words that have been analyzed.
	closedList := make([]*aStarWordNode, 0)
	//The aStarWordNode that relates to the start word selected.
	startNode := newAStarWordNode(sW)
	//The aStarWordNode that relates to the end word seletected.
	endNode := newAStarWordNode(eW)
	//List used to store the children nodes that relate to the current node being checked.
	childenNodes := make([]*aStarWordNode, 0)
	//Boolean used to indicate if a path has been found.
	foundResult = false

	//Calculate the estimated minimum cost from start to end word.
	startNode.HScore = calculateNodeCost(startNode.Word, endNode.Word)
	startNode.FScore = startNode.HScore
	//Add startWord to openList
	openList = append(openList, &startNode)
	//Add endword to the list of words to be analyzed.
	wordDictionary = append(wordDictionary, &endNode)

	//While there are still elements in openList continue analysis
	for len(openList) != 0 {
		//The current node being analyzed.
		var currentNode *aStarWordNode
		//int used to indicate the best potential score of all nodes in the open list
		bestFScore := -1
		//int used to indicate the best potential score of all nodes in the open list
		bestGScore := -1
		//Int used to indicate the position of the best score.
		index := 0

		//Find the best scored node in current openList
		for i, node := range openList {
			//Special case if the bestFScore is -1 then this is first pass through analysis.
			if bestFScore >= node.FScore || bestFScore == -1 {
				//Special case if the bestGScore is -1 then this is first pass through analysis.
				//We check for the lowest GScore after the best Fscore to make sure that children nodes are attached at the earliest point possible.
				if bestGScore >= node.GScore || bestGScore == -1 {
					//Store details of the best scored node in open list.
					bestFScore = node.FScore
					bestGScore = node.GScore
					currentNode = node
					index = i
				}
			}
		}

		//remove current node from openList
		openList = append(openList[:index], openList[index+1:]...)

		//If true we have found the solution
		if currentNode.Word == endNode.Word {
			foundResult = true
			break
		}

		//Get all the word nodes that are 1 step from the current node, update word dictionary so that all words found are removed from list to be analyzed.
		childenNodes, wordDictionary = generateNodeChildren(currentNode, wordDictionary)
		//G score (cost of path to this point) will always be current gscore + 1 for children as they are 1 step from the previous node.
		tempGScore := currentNode.GScore + 1

		//For each child node update the scores and add the node to the open list
		for _, cN := range childenNodes {
			if tempGScore < cN.GScore || cN.GScore == 0 {
				cN.GScore = tempGScore
				cN.HScore = calculateNodeCost(cN.Word, endNode.Word)
				cN.FScore = cN.GScore + cN.HScore
				cN.ParentNode = currentNode
				openList = append(openList, cN)
			}
		}

		//Append current node to closed list as it has now been analysed
		closedList = append(closedList, currentNode)
	}

	if foundResult {
		resultPath = getResultPath(endNode)
	} else {
		resultPath = []string{}
	}

	return
}

//Function to read in the word file and create a list of wordNodes from the data.
func readFile(startWord, endWord, fileLocation, delimiter string) []*aStarWordNode {
	//Open the file and log an error if there is one.
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	//Defer file.close to the end of this function.
	defer file.Close()

	//Array to store wordNodes.
	wD := make([]*aStarWordNode, 0)

	//create scanner for the file opened.
	scanner := bufio.NewScanner(file)
	//While there are still lines in the file:
	for scanner.Scan() {
		//Check the text is not the start or end word (these are dealt with seperately) and if not then create word node and add it to the array.
		if delimiter == "" {
			if scanner.Text() != startWord && scanner.Text() != endWord {
				aStarWordNode := newAStarWordNode(scanner.Text())
				wD = append(wD, &aStarWordNode)
			}
		} else {
			words := strings.Split(scanner.Text(), delimiter)
			for _, word := range words {
				if word != startWord && word != endWord {
					aStarWordNode := newAStarWordNode(word)
					wD = append(wD, &aStarWordNode)
				}
			}
		}

	}

	return wD
}

//Calculate the minimum potential cost from one word to another.
func calculateNodeCost(s, e string) int {
	//The maximum result will be if every letter is different in the two words (steps would be length).
	result := len(s)
	//length of the word starting from index 0
	wordLength := result - 1

	//For each letter in the word check if they match. If they are minus 1 required step from path cost result.
	for i := 0; i <= wordLength; i++ {
		if s[i] == e[i] {
			result = result - 1
		}
	}

	return result
}

//Generate all the children nodes when given a starting node and a list of potential nodes.
func generateNodeChildren(node *aStarWordNode, dict []*aStarWordNode) (childrenNodes, newDict []*aStarWordNode) {
	//The array to store the children nodes (maximum potential size / cap is length of aStarWordNode dictionary)
	childrenNodes = make([]*aStarWordNode, 0, len(dict))

	newDict = make([]*aStarWordNode, 0, len(dict))
	//Length of word being checked from an index of 0
	wordLength := len(node.Word) - 1
	//Number of letters that are the same from current node and the potential children.
	matchingLetters := 0

	//For each potential word calculate the number of matching letters, update the node and lists as required based on matching letters.
	for _, dictNode := range dict {
		for i := 0; i <= wordLength; i++ {
			if node.Word[i] == dictNode.Word[i] {
				matchingLetters++
			}
		}
		//This means there is only 1 letter different (matching are length-1) so is will be a child of the current node.
		if matchingLetters == wordLength {
			//Add the node to the childrenNode list.
			childrenNodes = append(childrenNodes, dictNode)
		} else {
			newDict = append(newDict, dictNode)
		}

		matchingLetters = 0
	}

	return
}

//For a given end node calculate the path to the first parent node.
func getResultPath(endNode aStarWordNode) []string {
	//Current Node in path
	currentNode := &endNode
	//Array holding the path of nodes.
	result := make([]string, 0)

	//While the current node has a parent node then continue to loop and add the current loop to the array.
	for currentNode.ParentNode != nil {
		result = append(result, currentNode.Word)
		currentNode = currentNode.ParentNode
	}
	//Add the current node to the array as the final node would not be included in the above loop as its parentNode == nil
	result = append(result, currentNode.Word)

	return result
}
