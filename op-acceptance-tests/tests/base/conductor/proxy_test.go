package conductor

import (
	"encoding/json"
	"testing"

	"github.com/ethereum-optimism/optimism/op-devstack/devtest"
	"github.com/ethereum-optimism/optimism/op-devstack/presets"
	"github.com/ethereum-optimism/optimism/op-devstack/sysgo"
)

// TestConductorProxyServesOnlyLeader verifies the conductor RPC proxy
// contract that batchers and proposers rely on to follow the active
// sequencer: the leader's conductor forwards execution, rollup, and admin
// requests to its sequencer, while a follower's conductor refuses them.
func TestConductorProxyServesOnlyLeader(gt *testing.T) {
	t := devtest.ParallelT(gt)
	sysgo.SkipOnKonaNode(t, "kona-node conductor support is tracked by #21906")

	sys := presets.NewMinimalWithConductors(t)

	leader := sys.Conductors.AwaitLeader()
	follower := sys.Conductors.Without(leader)[0]

	probes := []struct {
		method string
		args   []any
	}{
		{method: "eth_getBlockByNumber", args: []any{"latest", false}},
		{method: "optimism_syncStatus"},
		{method: "admin_sequencerActive"},
	}
	for _, probe := range probes {
		var result json.RawMessage
		t.Require().NoErrorf(leader.CallProxy(&result, probe.method, probe.args...),
			"expected conductor %s to proxy %s to its sequencer", leader, probe.method)

		err := follower.CallProxy(&result, probe.method, probe.args...)
		t.Require().ErrorContainsf(err, "refusing to proxy request to non-leader sequencer",
			"expected conductor %s to refuse proxying %s", follower, probe.method)
	}
}
