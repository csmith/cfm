package cfm

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSiteFixtures(t *testing.T) {
	entries, err := os.ReadDir("testdata/site")
	require.NoError(t, err)
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".md") || name == "README.md" {
			continue
		}
		base := strings.TrimSuffix(name, ".md")
		t.Run(
			base,
			func(t *testing.T) {
				input, err := os.ReadFile(filepath.Join("testdata", "site", name))
				require.NoError(t, err)
				want, err := os.ReadFile(filepath.Join("testdata", "site", base+".html"))
				require.NoError(t, err)
				assert.Equal(
					t,
					string(want),
					Markdown(string(input)),
					"golden mismatch for testdata/site/%s against testdata/site/%s.html",
					name,
					base,
				)
			},
		)
	}
}

func TestFixturePairs(t *testing.T) {
	entries, err := os.ReadDir("testdata/site")
	require.NoError(t, err)
	inputs := map[string]bool{}
	goldens := map[string]bool{}
	for _, entry := range entries {
		name := entry.Name()
		switch {
		case entry.IsDir(), name == "README.md", name == "generate.go":
			continue
		case strings.HasSuffix(name, ".md"):
			inputs[strings.TrimSuffix(name, ".md")] = true
		case strings.HasSuffix(name, ".html"):
			goldens[strings.TrimSuffix(name, ".html")] = true
		default:
			t.Errorf("unexpected file testdata/site/%s", name)
		}
	}
	for base := range inputs {
		assert.True(t, goldens[base], "missing golden testdata/site/%s.html", base)
	}
	for base := range goldens {
		assert.True(t, inputs[base], "missing input testdata/site/%s.md", base)
	}
}
