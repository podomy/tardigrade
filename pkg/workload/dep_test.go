// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package workload

import (
	"testing"

	"github.com/podomy/tardigrade/pkg/testutils/buildutil"
	"github.com/podomy/tardigrade/pkg/util/leaktest"
)

func TestDepWhitelist(t *testing.T) {
	defer leaktest.AfterTest(t)()

	// We want workload to be lightweight. If you need to add a package to this
	// set of deps, run it by danhhz first.
	buildutil.VerifyTransitiveWhitelist(t, "github.com/podomy/tardigrade/pkg/workload",
		[]string{
			`github.com/podomy/tardigrade/pkg/build`,
			`github.com/podomy/tardigrade/pkg/col/coldata`,
			`github.com/podomy/tardigrade/pkg/col/coltypes`,
			`github.com/podomy/tardigrade/pkg/util/bufalloc`,
			`github.com/podomy/tardigrade/pkg/util/encoding/csv`,
			`github.com/podomy/tardigrade/pkg/util/syncutil`,
			`github.com/podomy/tardigrade/pkg/util/timeutil`,
			`github.com/podomy/tardigrade/pkg/util/version`,
			`github.com/podomy/tardigrade/pkg/workload/histogram`,
		},
	)
}
