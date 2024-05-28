package main

import (
	"math/rand"
	"os"
	"time"

	"scheduler-framework-sample1/pkg"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Register custom filter to the scheduler framework.
	// Later they can consist of scheduler profile(s) and hence
	// used by various kinds of workloads.
	command := app.NewSchedulerCommand(
		app.WithPlugin(pkg.Name, pkg.New),
	)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
