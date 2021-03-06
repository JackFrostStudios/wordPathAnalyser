package wordPathAnalyser

//Structs used to hold the mocked input.
type aStarAnalyseMockInput struct {
	StartWord, EndWord, FileLocation, Delimiter string
	PathFound                                   bool
	ResultPathLength                            int
}
type aStarReadFileMockInput struct {
	StartWord, EndWord, FileLocation, Delimiter string
	ResultList                                  []*aStarWordNode
}
type aStarCalculateNodeCostMockInput struct {
	StartWord, EndWord string
	Result             int
}
type aStarGenerateNodeChildren struct {
	InputNode                                              *aStarWordNode
	InputDictionary, ResultChildrenNodes, ResultDictionary []*aStarWordNode
}
type aStarGetResultPath struct {
	EndNode    aStarWordNode
	ResultList []string
}
