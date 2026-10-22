// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package rowflow

import (
	"testing"

	"github.com/podomy/tardigrade/pkg/testutils/buildutil"
	"github.com/podomy/tardigrade/pkg/util/leaktest"
)

func TestNoLinkForbidden(t *testing.T) {
	defer leaktest.AfterTest(t)()

	buildutil.VerifyNoImports(t,
		"github.com/podomy/tardigrade/pkg/sql/rowflow", true,
		[]string{
			"github.com/podomy/tardigrade/pkg/sql/colflow",
			"github.com/podomy/tardigrade/pkg/sql/distsql",
		}, nil,
	)
}
