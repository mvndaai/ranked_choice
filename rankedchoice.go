package rankedchoice

import (
	"fmt"
	"log"
	"sort"
)

type Candidate string

type Vote struct {
	Choices []Candidate
}

type Step struct {
	NoChoiceLeft     int
	CandidateResults map[Candidate]int
}

type Result struct {
	TotalVotes int
	Winner     *Candidate
	Steps      []Step
}

func CalculateDropLowest(votes []Vote) (Result, error) {
	r := Result{TotalVotes: len(votes)}
	lowest := map[Candidate]bool{}

	// Unit we find a winner
	for {
		log.Println("starting step")
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
		log.Println("step results", step)

		winner, ok := step.Winner()
		if ok {
			log.Println("The winner is", winner)
			r.Winner = &winner
			break
		}

		low, err := step.Lowest()
		if err != nil {
			return r, err
		}
		lowest[low] = true

		if len(votes) == step.NoChoiceLeft {
			return r, fmt.Errorf("no winner")
		}
	}

	return r, nil
}

func (v Vote) TopAvaliableChoice(removed map[Candidate]bool) (Candidate, bool) {
	for _, choice := range v.Choices {
		if removed[choice] {
			continue
		}
		return choice, true
	}
	return "", false
}

func (s Step) Winner() (Candidate, bool) {
	var total int
	for _, votes := range s.CandidateResults {
		total += votes
	}

	for candidate, votes := range s.CandidateResults {
		percent := int(float64(votes/total) * 100)
		log.Println("candiate %", candidate, votes, total, percent)
		if percent > 50 {
			return candidate, true
		}
	}
	return "", false
}

func (s Step) Lowest() (Candidate, error) {
	if len(s.CandidateResults) == 0 {
		return "", fmt.Errorf("no candiates with votes")
	}

	keys := make([]Candidate, 0, len(s.CandidateResults))
	for key := range s.CandidateResults {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return s.CandidateResults[keys[i]] < s.CandidateResults[keys[j]]
	})

	if len(keys) > 1 {
		if s.CandidateResults[keys[0]] == s.CandidateResults[keys[1]] {
			// This error don't not show for 3 way ties
			return "", fmt.Errorf("tie between %v and %v", keys[0], keys[1])
		}
	}
	return keys[0], nil
}

func (s Step) Highest() (Candidate, error) {
	if len(s.CandidateResults) == 0 {
		return "", fmt.Errorf("no candiates with votes")
	}

	keys := make([]Candidate, 0, len(s.CandidateResults))
	for key := range s.CandidateResults {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return s.CandidateResults[keys[i]] > s.CandidateResults[keys[j]]
	})

	if len(keys) > 1 {
		if s.CandidateResults[keys[0]] == s.CandidateResults[keys[1]] {
			// This error don't not show for 3 way ties
			return "", fmt.Errorf("tie between %v and %v", keys[0], keys[1])
		}
	}
	return keys[0], nil
}
