package votes_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	rankedchoice "github.com/mvndaai/ranked_choice"
	"github.com/mvndaai/ranked_choice/votes"
	"github.com/stretchr/testify/require"
)

func saveTestResults(t *testing.T, b []byte) {
	filename := filepath.Join(".", t.Name()+".txt")
	f, err := os.Create(filename)
	require.NoError(t, err)
	defer f.Close()

	_, err = f.Write(b)
	require.NoError(t, err)
}

func TestHardCoded(t *testing.T) {
	vs := votes.HardCoded()
	r, err := rankedchoice.CalculateDropLowest(vs)
	require.NoError(t, err)

	b, err := json.MarshalIndent(r, "", "\t")
	require.NoError(t, err)
	saveTestResults(t, b)
	t.Log(string(b))
}
