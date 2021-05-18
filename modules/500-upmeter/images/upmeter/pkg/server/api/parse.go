package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"d8.io/upmeter/pkg/server/ranges"
)

// parseStepRange decodes 3 arguments
func parseStepRange(fromArg, toArg, stepArg string) (ranges.StepRange, error) {
	var (
		hasFrom = fromArg != ""
		hasTo   = toArg != ""
		hasStep = stepArg != ""
		err     error
	)

	rng := ranges.StepRange{Step: 300}

	if hasFrom {
		rng.From, err = parseTimestamp(fromArg)
		if err != nil {
			return rng, fmt.Errorf("from=%q is not timestamp: %v", fromArg, err)
		}
	} else {
		rng.From = time.Now().Truncate(5 * time.Minute).Add(-6 * time.Hour).Unix()
	}

	if hasTo {
		rng.To, err = parseTimestamp(toArg)
		if err != nil {
			return rng, fmt.Errorf("to=%q is not timestamp: %v", toArg, err)
		}
	} else {
		rng.To = time.Now().Truncate(5 * time.Minute).Unix()
	}

	if hasStep {
		rng.Step, err = parseDuration(stepArg)
		if err != nil {
			return rng, fmt.Errorf("step=%q is not duration: %v", stepArg, err)
		}
	} else {
		rng.Step = int64(300) // 5min
	}

	return rng, nil
}

func parseTimestamp(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func parseDuration(s string) (int64, error) {
	dur, err := time.ParseDuration(s)
	if err != nil {
		return parseTimestamp(s)
	}
	return int64(dur.Seconds()), nil
}

func parseDowntimeTypes(in string) []string {
	res := []string{}
	muteTypes := strings.Split(in, "!")
	for _, muteType := range muteTypes {
		switch muteType {
		case "Mnt":
			res = append(res, "Maintenance")
		case "Acd":
			res = append(res, "Accident")
		case "InfMnt":
			res = append(res, "InfrastructureMaintenance")
		case "InfAcd":
			res = append(res, "InfrastructureAccident")
		}
	}
	return res
}
