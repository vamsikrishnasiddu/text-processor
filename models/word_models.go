package models

type WordPair struct {
	Word      string `json:"word"`
	Occurence int    `json:"occurence"`
}

type WordPairList []WordPair

func (p WordPairList) Len() int           { return len(p) }
func (p WordPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p WordPairList) Less(i, j int) bool { return p[i].Occurence > p[j].Occurence }
