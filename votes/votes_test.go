package votes_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	rankedchoice "github.com/mvndaai/ranked_choice"
	"github.com/mvndaai/ranked_choice/multi"
	"github.com/mvndaai/ranked_choice/votes"
	"github.com/stretchr/testify/require"
)

func saveTestResults(t *testing.T, a any) {
	filename := filepath.Join(".", t.Name()+".txt")
	f, err := os.Create(filename)
	require.NoError(t, err)
	defer f.Close()

	b, err := json.MarshalIndent(a, "", "\t")
	require.NoError(t, err)

	_, err = f.Write(b)
	require.NoError(t, err)
}

func TestHardCodedRankedChoiceLowest(t *testing.T) {
	vs := votes.HardCoded()
	r, err := rankedchoice.CalculateDropLowest(vs)
	require.NoError(t, err)
	saveTestResults(t, r)
}

func TestHardCodedMulti(t *testing.T) {
	vs := votes.HardCoded()
	r, err := multi.Calculate(vs)
	require.NoError(t, err)
	saveTestResults(t, r)
}
