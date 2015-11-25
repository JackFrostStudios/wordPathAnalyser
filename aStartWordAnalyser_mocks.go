package wordPathAnalyser

//Structs used to hold the mocked input.
type aStarAnalyseMockInput struct {
	StartWord, EndWord, FileLocation, Delimiter string
	PathFound                                   bool
	ResultPath                                  []string
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
	InputNode           *aStarWordNode
	InputDictionary     []*aStarWordNode
	ResultChildrenNodes []*aStarWordNode
}
type aStarGetResultPath struct {
	EndNode    aStarWordNode
	ResultList []string
}
