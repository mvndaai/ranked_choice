package multi

import (
	"log"

	rankedchoice "github.com/mvndaai/ranked_choice"
)

func Calculate(votes []rankedchoice.Vote) (rankedchoice.Result, error) {
	r := rankedchoice.Result{TotalVotes: len(votes)}

	step := rankedchoice.Step{
		NoChoiceLeft:     0,
		CandidateResults: map[rankedchoice.Candidate]int{},
	}

	for _, vote := range votes {
		for _, c := range vote.Choices {
			if _, ok := step.CandidateResults[c]; !ok {
				step.CandidateResults[c] = 0
			}
			// record vote
			step.CandidateResults[c]++
		}
	}

	r.Steps = append(r.Steps, step)
	log.Println("step results", step)

	high, err := step.Highest()
	if err != nil {
		return r, err
	}
	log.Println("The winner is", high)
	r.Winner = &high

	return r, nil
}
