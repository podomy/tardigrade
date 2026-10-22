// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package storage

import (
	"context"
	"math/rand"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/podomy/tardigrade/pkg/base"
	"github.com/podomy/tardigrade/pkg/roachpb"
	"github.com/podomy/tardigrade/pkg/rpc"
	"github.com/podomy/tardigrade/pkg/rpc/nodedialer"
	"github.com/podomy/tardigrade/pkg/settings/cluster"
	"github.com/podomy/tardigrade/pkg/util"
	"github.com/podomy/tardigrade/pkg/util/hlc"
	"github.com/podomy/tardigrade/pkg/util/leaktest"
	"github.com/podomy/tardigrade/pkg/util/log"
	"github.com/podomy/tardigrade/pkg/util/netutil"
	"github.com/podomy/tardigrade/pkg/util/stop"
	"github.com/podomy/tardigrade/pkg/util/tracing"
	"github.com/podomy/tardigrade/pkg/util/uuid"
	"github.com/pkg/errors"
)

func TestRaftTransportStartNewQueue(t *testing.T) {
	defer leaktest.AfterTest(t)()
	ctx := context.Background()

	stopper := stop.NewStopper()
	defer stopper.Stop(ctx)

	st := cluster.MakeTestingClusterSettings()
	rpcC := rpc.NewContext(log.AmbientContext{}, &base.Config{Insecure: true},
		hlc.NewClock(hlc.UnixNano, 500*time.Millisecond), stopper, &st.Version)
	rpcC.ClusterID.Set(context.TODO(), uuid.MakeV4())

	// mrs := &dummyMultiRaftServer{}

	grpcServer := rpc.NewServer(rpcC)
	// RegisterMultiRaftServer(grpcServer, mrs)

	var addr net.Addr

	resolver := func(roachpb.NodeID) (net.Addr, error) {
		if addr == nil {
			return nil, errors.New("no addr yet") // should not happen in this test
		}
		return addr, nil
	}

	tp := NewRaftTransport(
		log.AmbientContext{Tracer: tracing.NewTracer()},
		cluster.MakeTestingClusterSettings(),
		nodedialer.New(rpcC, resolver),
		grpcServer,
		stopper,
	)

	ln, err := netutil.ListenAndServeGRPC(stopper, grpcServer, &util.UnresolvedAddr{NetworkField: "tcp", AddressField: "localhost:0"})
	if err != nil {
		t.Fatal(err)
	}

	addr = ln.Addr()

	defer func() {
		if ln != nil {
			_ = ln.Close()
		}
	}()

	_, existingQueue := tp.getQueue(1, rpc.SystemClass)
	if existingQueue {
		t.Fatal("queue already exists")
	}
	timeout := time.Duration(rand.Int63n(int64(5 * time.Millisecond)))
	log.Infof(ctx, "running test with a ctx cancellation of %s", timeout)
	ctxBoom, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		<-time.After(timeout)
		_ = ln.Close()
		ln = nil
		wg.Done()
	}()
	var stats raftTransportStats
	tp.startProcessNewQueue(ctxBoom, 1, rpc.SystemClass, &stats)

	wg.Wait()
}
