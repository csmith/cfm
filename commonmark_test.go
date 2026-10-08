package cfm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type commonMarkEntry struct {
	Example        int      `json:"example"`
	Section        string   `json:"section"`
	Included       bool     `json:"included"`
	ExcludedReason []string `json:"excluded_reason"`
	MarkdownFile   string   `json:"markdown_file"`
	HTMLFile       string   `json:"html_file"`
	Adapted        bool     `json:"adapted_html"`
	AdaptedFile    string   `json:"adapted_html_file"`
	ExpectedSource string   `json:"expected_source"`
}

type commonMarkManifest struct {
	Total         int               `json:"total_examples"`
	IncludedCount int               `json:"included_examples"`
	AdaptedCount  int               `json:"adapted_examples"`
	Examples      []commonMarkEntry `json:"examples"`
}

func TestCommonMark(t *testing.T) {
	manifestBytes, err := os.ReadFile(filepath.Join("testdata", "commonmark", "manifest.json"))
	require.NoError(t, err, "read corpus manifest")
	manifest := commonMarkManifest{}
	require.NoError(t, json.Unmarshal(manifestBytes, &manifest), "parse corpus manifest")
	specBytes, err := os.ReadFile(filepath.Join("testdata", "commonmark", "spec.json"))
	require.NoError(t, err, "read upstream examples")
	var upstream []json.RawMessage
	require.NoError(t, json.Unmarshal(specBytes, &upstream), "parse upstream examples")
	require.NotEmpty(t, upstream)
	assert.Equal(t, len(upstream), manifest.Total, "the manifest must cover the upstream corpus")
	require.Positive(t, manifest.IncludedCount, "the corpus must exercise the parser")
	assert.Len(
		t,
		manifest.Examples,
		manifest.Total,
		"the manifest must account for every CommonMark example",
	)

	included, adapted := validateCommonMarkManifest(t, manifest)
	assert.Equal(
		t,
		manifest.IncludedCount,
		included,
		"included_examples must match the listed included entries",
	)
	assert.Equal(
		t,
		manifest.AdaptedCount,
		adapted,
		"adapted_examples must match the listed adapted entries",
	)

	corpus := filepath.Join("testdata", "commonmark")
	for _, entry := range manifest.Examples {
		if !entry.Included {
			continue
		}
		t.Run(
			fmt.Sprintf("%03d %s", entry.Example, entry.Section),
			func(t *testing.T) {
				runCommonMarkExample(t, corpus, entry)
			},
		)
	}
}

func validateCommonMarkManifest(t *testing.T, manifest commonMarkManifest) (included, adapted int) {
	seen := make(map[int]bool, len(manifest.Examples))
	for _, entry := range manifest.Examples {
		assert.False(t, seen[entry.Example], "example %d listed twice", entry.Example)
		seen[entry.Example] = true
		assert.GreaterOrEqual(t, entry.Example, 1, "example number out of range")
		assert.LessOrEqual(t, entry.Example, manifest.Total, "example number out of range")
		if !entry.Included {
			assert.NotEmpty(
				t,
				entry.ExcludedReason,
				"excluded example %d must record a reason",
				entry.Example,
			)
			assert.False(t, entry.Adapted, "adapted example %d must be included", entry.Example)
			continue
		}
		included++
		assert.Empty(
			t,
			entry.ExcludedReason,
			"included example %d must not carry an exclusion reason",
			entry.Example,
		)
		if entry.Adapted {
			adapted++
		}
	}
	return included, adapted
}

func runCommonMarkExample(t *testing.T, corpus string, entry commonMarkEntry) {
	require.NotEmpty(t, entry.MarkdownFile, "example %d lacks an input file", entry.Example)
	require.NotEmpty(t, entry.HTMLFile, "example %d lacks an expectation file", entry.Example)
	markdown, err := os.ReadFile(filepath.Join(corpus, entry.MarkdownFile))
	require.NoError(t, err, "read input")
	wantFile := entry.HTMLFile
	if entry.Adapted {
		require.NotEmpty(t, entry.AdaptedFile, "adapted example %d lacks an adapted file", entry.Example)
		wantFile = entry.AdaptedFile
	}
	want, err := os.ReadFile(filepath.Join(corpus, wantFile))
	require.NoError(t, err, "read expectation")
	assert.Equal(
		t,
		string(want),
		Markdown(string(markdown)),
		"spec example %d (%s, expectation from %s)",
		entry.Example,
		entry.Section,
		entry.ExpectedSource,
	)
}
