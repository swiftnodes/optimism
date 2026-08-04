# SDM test ownership after public policy extraction

This table records the destination of assertions removed with the public block-warming policy and
vendored `op-rbuilder` / `rollup-boost`. Public tests own consensus mechanism and structural
validity; `optimism-premium` owns warming economics and live producer behavior.

| Removed public source | Assertion summary | Final owner and replacement | Gate |
|---|---|---|---|
| `rust/alloy-op-evm/src/post_exec/inspector/tests.rs` | Exact account/SLOAD/SSTORE warming amounts, exclusions, first warmer, deposit behavior, settlement touches, and cross-subblock carry | Premium `sdm-policy/src/block_warming.rs` unit tests; actual-EVM dispatch regression in premium payload tests | Premium `rust-unit-and-inprocess-tests` |
| Warming sections removed from `rust/alloy-op-evm/src/block/tests.rs` | Producer/verify accounting, raw versus canonical gas, malformed claims, failed execution, and candidate rollback | Public fixed/scripted-policy executor tests retain structural bounds and rollback; premium policy and builder tests retain real warming semantics | Public Rust unit tests + premium Rust unit/in-process tests |
| `TestSDMEnabledPayloadAndReplayMatch` | Producer payload, canonical tx hash/indexing, receipts, gas identities, and policy-aware replay | Public `TestSDMFixturePayloadReceiptAndAccounting` uses an independent one-gas oracle and structural replay; premium MODE A/B tests own policy correctness | Public SDM acceptance + premium `sdm-acceptance-tests` |
| `TestSDMStorageRefundBreakdown` | Same-slot and many-slot exact refund totals and no EIP-3529 cap | Premium block-warming policy/workload tests | Premium Rust unit/in-process and SDM acceptance |
| `TestSDMMixedWorkloadSmoke` | Successful, reverted, and out-of-gas workload behavior | Premium mixed-workload MODE A/B acceptance with explicit non-zero refund and receipt/payload assertions | Premium `sdm-acceptance-tests` |
| `flashblocks_sdm_phantom_test.go` | Declined candidate cannot leak real warming state | Public `declined_candidate_restores_refund_policy_snapshot` covers the generic snapshot mechanism; premium real-policy declined-candidate regression covers semantics | Public Rust unit + premium Rust unit/in-process |
| `flashblocks_sdm_test.go` | Flashblock SDM payload and canonical materialization | Premium MODE A stream-to-canonical tests | Premium integration and SDM acceptance |
| `flashblocks_stream_test.go` | Out-of-band `post_exec_tx`, replacement semantics, and streamed bytes equal canonical trailing tx | Premium subblocks stream tests | Premium integration tests |
| `flashblocks_transfer_test.go` | Live transfer progression through vendored producer topology | Premium non-SDM Kupcake/subblocks integration | Premium integration tests |
| `rust/op-rbuilder/crates/op-rbuilder/src/tests/sdm.rs` | Standard and flashblocks producer policy integration | Premium MODE A/B payload tests | Premium Rust unit/in-process + SDM acceptance |
| `rust/op-rbuilder` candidate-loop tests | Candidate rejection, gas/DA limits, and rollback in the real builder | Premium subblocks builder tests; generic public rollback remains in `alloy-op-evm` | Premium Rust unit/in-process |
| `rust/rollup-boost` flashblocks tests | Stream ingestion, replacement, sealing, and canonical payload forwarding | Premium subblocks/Kupcake integration | Premium integration tests |

Public coverage that remains required is listed in
[`sdm-public-architecture.md`](sdm-public-architecture.md): stock inertness, fixed-fixture payload and
receipt accounting, operator gate, Lagoon boundary, span/singular derivation, isolated force-build,
malformed payload rejection, cross-language encoding, and Kona/Cannon fault proofs.

No public deletion is justified solely by this document: the premium replacement must exist and pass
against the exact public revision before the cleanup is merge-ready.
