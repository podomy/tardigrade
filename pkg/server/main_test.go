// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package server_test

import (
	"os"
	"testing"

	"github.com/podomy/tardigrade/pkg/security"
	"github.com/podomy/tardigrade/pkg/security/securitytest"
	"github.com/podomy/tardigrade/pkg/server"
	"github.com/podomy/tardigrade/pkg/testutils/serverutils"
	"github.com/podomy/tardigrade/pkg/testutils/testcluster"
)

func TestMain(m *testing.M) {
	security.SetAssetLoader(securitytest.EmbeddedAssets)
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)
	os.Exit(m.Run())
}

//go:generate ../util/leaktest/add-leaktest.sh *_test.go
