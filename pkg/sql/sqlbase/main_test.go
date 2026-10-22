// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package sqlbase_test

import (
	"os"
	"testing"

	"github.com/podomy/tardigrade/pkg/security"
	"github.com/podomy/tardigrade/pkg/security/securitytest"
	"github.com/podomy/tardigrade/pkg/server"
	"github.com/podomy/tardigrade/pkg/testutils/serverutils"
	"github.com/podomy/tardigrade/pkg/util/randutil"
)

func TestMain(m *testing.M) {
	security.SetAssetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	os.Exit(m.Run())
}
