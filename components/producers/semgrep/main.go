package main

import (
	"fmt"
	"log"

	v1 "github.com/ocurity/dracon/api/proto/v1"
	"github.com/ocurity/dracon/components/producers/semgrep/types"
	"github.com/ocurity/dracon/pkg/context"

	"github.com/ocurity/dracon/components/producers"
)

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	inFile, err := producers.ReadInFile()
	if err != nil {
		log.Fatal(err)
	}

	var results types.SemgrepResults
	if err := producers.ParseJSON(inFile, &results); err != nil {
		log.Fatal(err)
	}

	issues, err := parseIssues(results)
	if err != nil {
		log.Fatal(err)
	}
	if err := producers.WriteDraconOut(
		"semgrep",
		issues,
	); err != nil {
		log.Fatal(err)
	}
}

func parseIssues(out types.SemgrepResults) ([]*v1.Issue, error) {
	issues := []*v1.Issue{}

	results := out.Results

	for _, r := range results {

		// Map the semgrep severity levels to dracon severity levels
		severityMap := map[string]v1.Severity{
			"INFO":    v1.Severity_SEVERITY_INFO,
			"WARNING": v1.Severity_SEVERITY_MEDIUM,
			"ERROR":   v1.Severity_SEVERITY_HIGH,
		}

		sev := severityMap[r.Extra.Severity]
		iss := &v1.Issue{
			Target:      fmt.Sprintf("%s:%v-%v", r.Path, r.Start.Line, r.End.Line),
			Type:        r.Extra.Message,
			Title:       r.CheckID,
			Severity:    sev,
			Cvss:        0.0,
			Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
			Description: fmt.Sprintf("%s\n extra lines: %s", r.Extra.Message, r.Extra.Lines),
		}
		cs, err := context.ExtractCode(iss)
		if err != nil {
			return nil, err
		}
		iss.ContextSegment = &cs
		issues = append(issues, iss)
	}
	return issues, nil
}
