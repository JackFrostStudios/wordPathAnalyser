package wordPathAnalyser

type aStarWordNode struct {
	//fScore - Total estimated number of steps to goal.
	//gscore - Total cost of current path.
	//hScore - Predicted number of steps to goal.
	FScore, GScore, HScore int
	ParentNode             *aStarWordNode
	Word                   string
}

func newAStarWordNode(word string) aStarWordNode {
	return aStarWordNode{
		FScore: 0,
		GScore: 0,
		HScore: 0,
		Word:   word,
	}
}
