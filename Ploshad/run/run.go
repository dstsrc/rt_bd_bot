package run
import()

type geometry interface{
FindS() int
}

func RunS(g geometry) int {
 return g.FindS()
}