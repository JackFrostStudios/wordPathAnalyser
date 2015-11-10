package wordPathAnalyser

type aStarWordNode struct {
	//fScore - Total estimated number of steps to goal.
	//gscore - Total cost of current path.
	//hScore - Predicted number of steps to goal.
	fScore, gScore, hScore int
	parentNode             *aStarWordNode
	word                   string
}

func newAStarWordNode(word string) aStarWordNode {
	return aStarWordNode{
		fScore: 0,
		gScore: 0,
		hScore: 0,
		word:   word,
	}
}
