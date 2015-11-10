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
//OUTPUT: path found result (Boolean), path from end word to start word ([]string)
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
	startNode.hScore = calculateNodeCost(startNode.word, endNode.word)
	startNode.fScore = startNode.hScore
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
		//int used to indicate the best remaining step score of all nodes in the open list.
		bestHScore := -1
		//Int used to indicate the position of the best score.
		index := 0

		//Find the best scored node in current openList
		for i, node := range openList {
			//Special case if the bestFScore is -1 then this is first pass through analysis.
			if bestFScore >= node.fScore || bestFScore == -1 {
				//We may get nodes in the list that have the same overall score but are 1 step back.
				//Checking that the Hscore is the lowest of the best F scores will ensure the best path is checked first.
				if bestHScore > node.hScore || bestHScore == -1 {
					//Store details of the best scored node in open list.
					bestFScore = node.fScore
					currentNode = node
					index = i
				}
			}
		}

		//remove current node from openList
		openList = append(openList[:index], openList[index+1:]...)

		//If true we have found the solution
		if currentNode.word == endNode.word {
			foundResult = true
			break
		}

		//Get all the word nodes that are 1 step from the current node, update word dictionary so that all words found are removed from list to be analyzed.
		childenNodes, wordDictionary = generateNodeChildren(currentNode, wordDictionary)
		//G score (cost of path to this point) will always be current gscore + 1 for children as they are 1 step from the previous node.
		tempGScore := currentNode.gScore + 1

		//For each child node update the scores and add the node to the open list
		for _, cN := range childenNodes {
			cN.gScore = tempGScore
			cN.hScore = calculateNodeCost(cN.word, endNode.word)
			cN.fScore = cN.gScore + cN.hScore
			openList = append(openList, cN)
		}

		//Add the current node to the closed list as it has now been analyzed.
		closedList = append(closedList, currentNode)
	}

	resultPath = getResultPath(endNode)

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
		if scanner.Text() != startWord && scanner.Text() != endWord {
			if delimiter == "" {
				aStarWordNode := newAStarWordNode(scanner.Text())
				wD = append(wD, &aStarWordNode)
			} else {
				words := strings.Split(scanner.Text(), delimiter)
				for _, word := range words {
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
	for i := 0; i < wordLength; i++ {
		if s[i] == e[i] {
			result = result - 1
		}
	}

	return result
}

//Generate all the children nodes when given a starting node and a list of potential nodes.
func generateNodeChildren(node *aStarWordNode, dict []*aStarWordNode) (childrenNodes, newDict []*aStarWordNode) {
	//The node dictionary once the node children have been removed.
	newDict = make([]*aStarWordNode, 0, len(dict))
	//The array to store the children nodes (maximum potential size / cap is length of aStarWordNode dictionary)
	childrenNodes = make([]*aStarWordNode, 0, len(dict))

	//Length of word being checked from an index of 0
	wordLength := len(node.word) - 1
	//Number of letters that are the same from current node and the potential children.
	matchingLetters := 0

	//For each potential word calculate the number of matching letters, update the node and lists as required based on matching letters.
	for _, dictNode := range dict {
		for i := 0; i <= wordLength; i++ {
			if node.word[i] == dictNode.word[i] {
				matchingLetters++
			}
		}
		//This means there is only 1 letter different (matching are length-1) so is will be a child of the current node.
		if matchingLetters == wordLength {
			//Update the aStarWordNode with the parentNode
			dictNode.parentNode = node
			//Add the node to the childrenNode list.
			childrenNodes = append(childrenNodes, dictNode)

			//Otherwise if the matching letters is less then word length add it to the list of words to remain in the pool.
		} else if matchingLetters < wordLength {
			newDict = append(newDict, dictNode)
		}
		//If the word node has not matched on either of the above and gets here it means the word is the same as the current node so it can be removed from the list to be checked.

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
	for currentNode.parentNode != nil {
		result = append(result, currentNode.word)
		currentNode = currentNode.parentNode
	}
	//Add the current node to the array as the final node would not be included in the above loop as its parentNode == nil
	result = append(result, currentNode.word)

	return result
}
