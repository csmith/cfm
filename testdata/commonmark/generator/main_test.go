package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeFixtureTree(t *testing.T, files map[string]string) string {
	t.Helper()
	dir := t.TempDir()
	for name, content := range files {
		path := filepath.Join(dir, filepath.FromSlash(name))
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
		require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
	}
	return dir
}

func plainExample() []specExample {
	return []specExample{{
		Markdown:  "hello\n",
		HTML:      "<p>hello</p>\n",
		Example:   1,
		Section:   "Paragraphs",
		StartLine: 1,
		EndLine:   1,
	}}
}

func strikethroughLossExample() []specExample {
	return []specExample{{
		Markdown:  "~~struck away~~\n",
		HTML:      "<p>~~struck away~~</p>\n",
		Example:   1,
		Section:   "Strikethrough",
		StartLine: 1,
		EndLine:   1,
	}}
}

func TestLoadOverridesReadsCorpusOverridesDir(t *testing.T) {
	dir := writeFixtureTree(t, map[string]string{
		"overrides/042.html": "<p>hand authored</p>\n",
	})
	overrides, err := loadOverrides(dir)
	require.NoError(t, err)
	assert.Equal(t, map[int]string{42: "<p>hand authored</p>\n"}, overrides)
}

func TestLoadOverridesMissingDirectoryIsEmpty(t *testing.T) {
	overrides, err := loadOverrides(filepath.Join(t.TempDir(), "overrides"))
	require.NoError(t, err)
	assert.Empty(t, overrides)
}

func TestLoadOverridesRejectsNonCanonicalNames(t *testing.T) {
	for _, name := range []string{"42.html", "0042.html", "042.htm", "spec-042.html"} {
		t.Run(name, func(t *testing.T) {
			dir := writeFixtureTree(t, map[string]string{
				"overrides/" + name: "<p>x</p>\n",
			})
			_, err := loadOverrides(dir)
			assert.ErrorContains(t, err, "bad override file name")
		})
	}
}

func TestGenerateWritesIncludedExample(t *testing.T) {
	dir := t.TempDir()
	m, audit, err := generate(dir, plainExample())
	require.NoError(t, err)
	assert.Empty(t, audit)
	require.Len(t, m.Examples, 1)
	entry := m.Examples[0]
	assert.True(t, entry.Included)
	assert.Equal(t, "spec", entry.ExpectedSource)
	assert.Equal(t, "examples/001.md", entry.MarkdownFile)
	assert.Equal(t, 1, m.Total)
	assert.Equal(t, 1, m.Included)
	assert.Equal(t, 0, m.Adapted)
	for _, name := range []string{"examples/001.md", "examples/001.html", "manifest.json"} {
		_, statErr := os.Stat(filepath.Join(dir, filepath.FromSlash(name)))
		assert.NoError(t, statErr)
	}
}

func TestGenerateFailsOnContentLossWithoutOverride(t *testing.T) {
	_, _, err := generate(t.TempDir(), strikethroughLossExample())
	require.Error(t, err)
	require.ErrorContains(t, err, "example 1 (Strikethrough)")
	require.ErrorContains(t, err, "missing content")
	assert.ErrorContains(t, err, "overrides/001.html")
}

func TestGenerateAcceptsContentLossWithOverride(t *testing.T) {
	dir := writeFixtureTree(t, map[string]string{
		"overrides/001.html": "<p>struck away</p>\n",
	})
	m, audit, err := generate(dir, strikethroughLossExample())
	require.NoError(t, err)
	require.Len(t, m.Examples, 1)
	assert.Equal(t, "hand-authored", m.Examples[0].ExpectedSource)
	assert.True(t, m.Examples[0].Adapted)
	assert.Equal(t, "examples/001.adapted.html", m.Examples[0].AdaptedFile)
	written, err := os.ReadFile(filepath.Join(dir, "examples", "001.adapted.html"))
	require.NoError(t, err)
	assert.Equal(t, "<p>struck away</p>\n", string(written))
	require.Len(t, audit, 1)
	assert.Contains(t, audit[0], "DATA LOSS example 1 (Strikethrough)")
}

func TestGenerateRejectsUnusedOverrides(t *testing.T) {
	dir := writeFixtureTree(t, map[string]string{
		"overrides/002.html": "<p>orphan</p>\n",
	})
	_, _, err := generate(dir, plainExample())
	require.Error(t, err)
	assert.ErrorContains(t, err, "unused overrides: 002.html")
}

func TestGenerateOverrideMatchingSpecIsNotAdapted(t *testing.T) {
	dir := writeFixtureTree(t, map[string]string{
		"overrides/001.html": "<p>hello</p>\n",
	})
	m, _, err := generate(dir, plainExample())
	require.NoError(t, err)
	require.Len(t, m.Examples, 1)
	assert.Equal(t, "hand-authored", m.Examples[0].ExpectedSource)
	assert.False(t, m.Examples[0].Adapted)
	assert.Empty(t, m.Examples[0].AdaptedFile)
}
