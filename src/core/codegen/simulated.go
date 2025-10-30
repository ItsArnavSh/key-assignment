package codegen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/osdc/resrap"
)

func GenerateLargeCodebase() error {
	rs := resrap.NewResrap()

	// Load grammars
	grammars := map[string]string{
		"C":   "data/C.g4",
		"JS":  "data/JS.g4",
		"SQL": "data/SQL.g4",
	}
	for lang, path := range grammars {
		if err := rs.ParseGrammarFile(lang, path); err != nil {
			return fmt.Errorf("failed to parse %s grammar: %v", lang, err)
		}
	}

	// Create output folder
	if err := os.MkdirAll("gen", 0755); err != nil {
		return fmt.Errorf("failed to create gen dir: %v", err)
	}

	// Config: a few hundred files, each very large
	codeSpecs := map[string]struct {
		StartRule string
		FileCount int
		Chunks    int // how many snippets per file
		Ext       string
	}{
		"C":   {"program", 50, 1000, ".c"},
		"JS":  {"program", 50, 1000, ".js"},
		"SQL": {"program", 20, 500, ".sql"},
	}

	for lang, spec := range codeSpecs {
		for i := 1; i <= spec.FileCount; i++ {
			var combined string
			for j := 0; j < spec.Chunks; j++ {
				combined += rs.GenerateRandom(lang, spec.StartRule, 200) + "\n"
			}

			fileName := filepath.Join("gen", fmt.Sprintf("%s_%d%s", lang, i, spec.Ext))
			if err := os.WriteFile(fileName, []byte(combined), 0644); err != nil {
				return fmt.Errorf("failed to write %s: %v", fileName, err)
			}
			fmt.Printf("Generated large file: %s (approx %d snippets)\n", fileName, spec.Chunks)
		}
	}

	fmt.Println("Large sample codebase generated successfully in ./gen/")
	return nil
}

func main() {
	if err := GenerateLargeCodebase(); err != nil {
		fmt.Println("Error:", err)
	}
}
