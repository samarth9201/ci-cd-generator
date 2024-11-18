package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/samarth9201/ci-cd-generator/utils"
)

func main() {

	dir := flag.String("dir", ".", "Directory To Scan")
	flag.Parse()

	files := 0

	if *dir == "" {
		*dir = utils.AskQuestion("Specify the repository to scan for configurations: ")
	}

	valid, err := utils.IsValidDirectory(*dir)
	if err != nil {
		log.Fatalf("Error Reading Directory : %v", err)
		return
	}
	if !valid {
		log.Fatalf("Invalid Directory : %v", *dir)
		return
	}

	triggers := utils.SetupTriggers()
	runners, count := utils.SetupRunners()
	files += count

	config := utils.Config{
		Triggers: triggers,
		Runners:  runners,
		Path:     *dir,
		Count:    files,
	}

	pipeline := utils.GeneratePipeline(config)

	fmt.Println(pipeline)
}
