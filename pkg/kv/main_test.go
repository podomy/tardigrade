// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package kv_test

import (
	"os"
	"testing"

	"github.com/podomy/tardigrade/pkg/security"
	"github.com/podomy/tardigrade/pkg/security/securitytest"
	"github.com/podomy/tardigrade/pkg/server"
	"github.com/podomy/tardigrade/pkg/testutils/buildutil"
	"github.com/podomy/tardigrade/pkg/testutils/serverutils"
	"github.com/podomy/tardigrade/pkg/util/leaktest"
)

//go:generate ../util/leaktest/add-leaktest.sh *_test.go

func init() {
	security.SetAssetLoader(securitytest.EmbeddedAssets)
}

func TestForbiddenDeps(t *testing.T) {
	defer leaktest.AfterTest(t)()
	// Verify kv does not depend on storage (or any of its subpackages).
	buildutil.VerifyNoImports(t,
		"github.com/podomy/tardigrade/pkg/kv", true,
		// TODO(tschottdorf): should really disallow ./storage/... but at the
		// time of writing there's a (legit) dependency on `enginepb`.
		[]string{
			"github.com/podomy/tardigrade/pkg/storage",
			"github.com/podomy/tardigrade/pkg/storage/engine",
		},
		[]string{})
}

func TestMain(m *testing.M) {
	serverutils.InitTestServerFactory(server.TestServerFactory)
	os.Exit(m.Run())
}
