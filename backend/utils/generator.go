package utils

import (
	"fmt"
	"strings"
)

type Config struct {
	Triggers map[string]interface{} `json:"triggers"`
	Runners  []string               `json:"operating_systems"`
	Path     string                 `json:"directory_path"`
	Count    int                    `json:"file_count"`
}

func SetupTriggers() map[string]interface{} {
	triggers := make(map[string]interface{})

	// Basic triggers
	if AskYesNo("Do you want the pipeline to run on every push to 'main'?") {
		triggers["push"] = map[string][]string{
			"branches": {"main"},
		}
	}

	if AskYesNo("Do you want the pipeline to run on pull requests targeting 'main'?") {
		triggers["pull_request"] = map[string][]string{
			"branches": {"main"},
		}
	}

	// Advanced triggers
	if AskYesNo("Do you want to schedule the pipeline?") {
		schedule := AskQuestion("Specify the cron schedule (e.g., '0 0 * * 1' for every Monday):")
		triggers["schedule"] = []map[string]string{
			{"cron": schedule},
		}
	}

	if AskYesNo("Do you want to allow manual triggering?") {
		triggers["workflow_dispatch"] = struct{}{}
	}

	return triggers
}

func SetupRunners() ([]string, int) {
	var runners []string
	count := 0
	fmt.Println("Select the runner(s) you want to use: see https://docs.github.com/en/actions/using-github-hosted-runners/using-github-hosted-runners/about-github-hosted-runners. (Type 'done' when done)")
	for {
		osChoice := AskQuestion("Add new Label (e.g., ubuntu-latest, windows-latest, macos-latest) : ")
		if strings.ToLower(osChoice) == "done" {
			break
		}
		runners = append(runners, osChoice)
		count += 1
	}
	return runners, count
}

func GeneratePipeline(config Config) string {
	pipeline := "name: CI Pipeline\n\non:\n"

	for trigger, details := range config.Triggers {
		pipeline += fmt.Sprintf("  %s:\n", trigger)
		switch v := details.(type) {
		case map[string][]string:
			for key, value := range v {
				pipeline += fmt.Sprintf("    %s:\n", key)
				for _, value := range value {
					pipeline += fmt.Sprintf("      - '%s'\n", value)
				}
			}
		case []map[string]string:
			for _, schedule := range v {
				for key, value := range schedule {
					pipeline += fmt.Sprintf("    %s: '%s'\n", key, value)
				}
			}
		default:
			pipeline += "    {}\n"
		}
	}

	// Add jobs section
	pipeline += "\njobs:\n  build:\n    runs-on: "
	if len(config.Runners) > 1 {
		pipeline += "\n      matrix:\n        os:\n"
		for _, os := range config.Runners {
			pipeline += fmt.Sprintf("          - %s\n", os)
		}
		pipeline += "    strategy:\n      fail-fast: false\n"
	} else {
		pipeline += config.Runners[0] + "\n"
	}

	// Placeholder for steps
	pipeline += `
    steps:
      - name: Checkout code
        uses: actions/checkout@v2
`

	return pipeline
}
