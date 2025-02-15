package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

func GetStatus() (Status, error) {
	text, err := GetPmsetOutput()
	if err != nil {
		return Status{}, err
	}
	return ParsePmsetOutput(text)
}

func GetPmsetOutput() (string, error) {
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var pmsetOutput = regexp.MustCompile("([0-9]+)%")

func ParsePmsetOutput(s string) (Status, error) {

	matches := pmsetOutput.FindStringSubmatch(s)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset. output: %v", s)
	}
	charge, e := strconv.Atoi(matches[1])
	if e != nil {
		return Status{}, fmt.Errorf("failed to parse charge. percent: %q", matches[1])
	}
	return Status{
		ChargePercent: charge,
	}, nil
}
