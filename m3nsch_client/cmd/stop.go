// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"log"

	"github.com/m3db/m3nsch/coordinator"

	"github.com/m3db/m3x/instrument"
	"github.com/spf13/cobra"
)

var (
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stop the load generation",
		Long:  "Stop stops the load generation on each agent process",
		Run:   stopExec,
		Example: `# Stop the load generation process on various agents:
./m3nsch_client -e "<agent1_host:agent1_port>,...,<agentN_host>:<agentN_port>" stop`,
	}
)

func stopExec(_ *cobra.Command, _ []string) {
	if !gFlags.isValid() {
		log.Fatalf("unable to execute, invalid flags\n%s", M3nschCmd.UsageString())
	}

	var (
		iopts            = instrument.NewOptions()
		logger           = iopts.Logger()
		mOpts            = coordinator.NewOptions(iopts)
		coordinator, err = coordinator.New(mOpts, gFlags.endpoints)
	)

	if err != nil {
		logger.Fatalf("unable to create coordinator: %v", err)
	}
	defer coordinator.Teardown()

	err = coordinator.Stop()
	if err != nil {
		logger.Fatalf("unable to stop workload: %v", err)
	}

	logger.Infof("workload stopped!")
}
