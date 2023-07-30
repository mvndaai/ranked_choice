package rankedchoice

import "log"

type Candidate struct {
	Name        string
	Affiliation string
}

type Vote struct {
	Choices []Candidate
}

type Step struct {
	NoChoiceLeft     int
	CandidateResults map[Candidate]int
}

//type

type Result struct {
	TotalVotes int
	Winner     *Candidate
	Steps      []Step
}

func CalculateDropLowest(votes []Vote) Result {
	r := Result{}
	lowest := map[Candidate]bool{}

	// Unit we find a winner
	for true {

		step := Step{
			NoChoiceLeft:     0,
			CandidateResults: map[Candidate]int{},
		}

		for _, vote := range votes {
			c, ok := vote.TopAvaliableChoice(lowest)
			// note if no choices avaliable
			if !ok {
				step.NoChoiceLeft++
				continue
			}
			// ensure the candiate exists in map
			if _, ok := step.CandidateResults[c]; !ok {
				step.CandidateResults[c] = 0
			}
			// record vote
			step.CandidateResults[c]++
		}
		r.Steps = append(r.Steps, step)

		winner, ok := step.Winner()
		if ok {
			r.Winner = &winner
			break
		}

		if len(votes) == step.NoChoiceLeft {
			log.Println("No winner?")
			break
		}
	}

	return r
}

func (v Vote) TopAvaliableChoice(removed map[Candidate]bool) (Candidate, bool) {
	for _, choice := range v.Choices {
		if removed[choice] {
			continue
		}
		return choice, true
	}
	return Candidate{}, false
}

func (s Step) Winner() (Candidate, bool) {
	var total int
	for _, votes := range s.CandidateResults {
		total += votes
	}

	for candidate, votes := range s.CandidateResults {
		percent := votes / total
		if percent > 50 {
			return candidate, true
		}
	}
	return Candidate{}, false
}
