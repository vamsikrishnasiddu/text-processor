package services

import (
	"sort"

	"github.com/vamsikrishnasiddu/text-processor/models"
	"github.com/vamsikrishnasiddu/text-processor/rest_errors"
	"github.com/vamsikrishnasiddu/text-processor/utils"
)

var (
	WordService wordServiceInterface = &wordService{}
)

type wordService struct{}

type wordServiceInterface interface {
	ProcessWordOccurence(text string) (models.WordPairList, rest_errors.RestErr)
}

// ProcessWordOccurence takes the text and Creates the Word pairs with their occurences.
// if no error it returns the list of wordPairs sorted on basis of highest repeated word.
func (w *wordService) ProcessWordOccurence(text string) (models.WordPairList, rest_errors.RestErr) {
	m := utils.WordCount(text)
	// if no text is passed we will return with an error
	if len(m) == 0 {
		err := rest_errors.NewBadRequestError("invalid request body")
		return nil, err
	}

	p := make(models.WordPairList, len(m))

	q := models.WordPairList{}

	i := 0
	for k, v := range m {
		p[i] = models.WordPair{Word: k, Occurence: v}
		i++
	}

	// sorting the Words based on the highest occurence
	sort.Sort(p)

	// appending the top 10 most used word pairs into a new wordPairlist
	for i := 0; i < len(p); i++ {
		if i < len(p) && i < 10 {
			q = append(q, p[i])
		}
	}

	return q, nil
}
