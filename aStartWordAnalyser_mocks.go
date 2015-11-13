package wordPathAnalyser

//Structs used to hold the mocked input.
type aStarAnalyseMockInput struct {
	startWord, endWord, fileLocation, delimiter string
	pathFound                                   bool
	resultPath                                  []string
}
type aStarReadFileMockInput struct {
	startWord, endWord, fileLocation, delimiter string
	resultList                                  []aStarWordNode
}
