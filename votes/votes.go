package votes

import (
	rankedchoice "github.com/mvndaai/ranked_choice"
)

func HardCoded() []rankedchoice.Vote {
	apple := rankedchoice.Candidate("Apple")
	orange := rankedchoice.Candidate("Orange")
	banana := rankedchoice.Candidate("Banana")
	kiwi := rankedchoice.Candidate("Kiwi")

	return []rankedchoice.Vote{
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},
		{Choices: []rankedchoice.Candidate{apple, banana}},

		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},
		{Choices: []rankedchoice.Candidate{orange, banana}},

		{Choices: []rankedchoice.Candidate{banana, orange}},
		{Choices: []rankedchoice.Candidate{banana, orange}},
		{Choices: []rankedchoice.Candidate{banana, orange}},
		{Choices: []rankedchoice.Candidate{banana, orange}},
		{Choices: []rankedchoice.Candidate{banana, orange}},
		{Choices: []rankedchoice.Candidate{banana, apple}},
		{Choices: []rankedchoice.Candidate{banana, apple}},
		{Choices: []rankedchoice.Candidate{banana, apple}},
		{Choices: []rankedchoice.Candidate{banana, apple}},
		{Choices: []rankedchoice.Candidate{banana, apple}},

		{Choices: []rankedchoice.Candidate{kiwi, banana, apple}},
		{Choices: []rankedchoice.Candidate{kiwi, banana, apple}},
		{Choices: []rankedchoice.Candidate{kiwi, banana, orange}},
		{Choices: []rankedchoice.Candidate{kiwi, banana, orange}},
		{Choices: []rankedchoice.Candidate{kiwi, banana}},
		{Choices: []rankedchoice.Candidate{kiwi, banana}},
		{Choices: []rankedchoice.Candidate{kiwi}},
		{Choices: []rankedchoice.Candidate{kiwi}},
		{Choices: []rankedchoice.Candidate{kiwi}},
	}
}
