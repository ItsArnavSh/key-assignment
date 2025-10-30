package server

import (
	"fmt"
	"os"
	"path/filepath"
	patternmatcher "stack/src/core/pattern_matcher"
	sourceidentifier "stack/src/core/report"
	"stack/src/entity"
	"strings"
)

type Server struct {
	pm            *patternmatcher.AhoCorasick
	inventory_map map[string]entity.InventorySecret
	CurrentURL    string
}

func NewServer() Server {
	inventory, err := patternmatcher.Load_json("data/inventory.json")
	if err != nil {
		return Server{}
	}
	inven_map := make(map[string]entity.InventorySecret)
	for _, inventory_item := range inventory {
		inven_map[inventory_item.Value] = inventory_item
	}
	values := make([]string, 0, len(inventory))
	for _, item := range inventory {
		values = append(values, item.Value)
	}
	pm := patternmatcher.NewAhoCorasick(values)

	return Server{pm: pm, inventory_map: inven_map}
}
func (s *Server) StartServer() {

}

func (s *Server) Scan(src string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read dir %s: %v", src, err)
	}

	for _, entry := range entries {
		path := filepath.Join(src, entry.Name())

		if entry.IsDir() {
			// recursively scan subdirectory
			if err := s.Scan(path); err != nil {
				return err
			}
		} else {
			// scan file contents
			if err := s.ScanFile(path); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Server) ScanFile(path string) error {
	data, err := os.ReadFile(path)
	reporter := sourceidentifier.Report{}
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", path, err)
	}

	content := string(data)
	res := s.pm.Search(content) // suppose this returns map[string][]int or similar
	if len(res) == 0 {
		return nil
	}

	lines := strings.Split(content, "\n")

	for secret := range res {
		// find which line contains the secret
		for i, line := range lines {
			if strings.Contains(line, secret) {
				start := max(i-5, 0)
				end := min(i+5, len(lines))
				context := strings.Join(lines[start:end], "\n")

				secretFound := s.inventory_map[secret]
				report := entity.BasicReport{
					Source:  path,
					Secret:  secretFound,
					Context: context,
				}
				reporter.GenerateReport(report, s.CurrentURL)
				break
			}
		}
	}

	return nil
}
func (s *Server) ScanCodebase(loc string, url string) {
	s.CurrentURL = url
	s.Scan(loc)
}
