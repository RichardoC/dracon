// Package context offers a set of methods which permit components to
//
//	add information about the context of each vulnerability
//	this information is highly depended on the actual vulnerability and the component.
//	For example for SAST components, context can be a call graph or
//	a few lines of code before and after the line that triggered the finding.
//	For DAST components it can be a serialised request/response.
package context

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	v1 "github.com/ocurity/dracon/api/proto/v1"
)

// DefaultLineRange controls how many lines of code context will be returned by default
const DefaultLineRange = 10

// ExtractCode takes a path and line (or line range) for a vulnerable code segment and
// returns the DefaultLineRange lines above and the DefaultLineRange lines below the vulnerability.
// It does not take into account end of function.
func ExtractCode(finding *v1.Issue) (string, error) {
	path := ""
	lineRange := ""
	lineFrom := 0
	lineTo := 0

	split := strings.Split(finding.Target, ":")
	if len(split) < 2 {
		path = finding.Target
		lineRange = "0"
	} else {
		path = split[0]
		lineRange = split[1]
	}
	if strings.Contains(lineRange, "-") {
		lines := strings.Split(lineRange, "-")
		lf, err := strconv.Atoi(lines[0])
		if err != nil {
			return "", err
		}
		lineFrom = lf
		lt, err := strconv.Atoi(lines[1])
		if err != nil {
			return "", err
		}
		lineTo = lt

	} else {
		lf, err := strconv.Atoi(lineRange)
		if err != nil {
			return "", err
		}
		lineFrom = lf
		lineTo = lf
	}
	if lineFrom < DefaultLineRange {
		lineFrom = 0
	} else {
		lineFrom = lineFrom - DefaultLineRange
	}
	lineTo = lineTo + DefaultLineRange
	handle, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("context pkg could not open file in path %s, err:%v", path, err)
	}
	sc := bufio.NewScanner(handle)
	pos := 0
	lines := []string{}
	for sc.Scan() {
		if lineFrom <= pos && pos < lineTo {
			lines = append(lines, sc.Text())
		}
		pos++
	}
	return strings.Join(lines, "\n"), nil
}
