package attestation

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
)

var (
	validatorLastVoteGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "validators_last_vote",
		Help: "Votes of validators, updated when there's a new attestation",
	}, []string{
		"validatorIndex",
	})
	totalAttestationSeen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_seen_attestations",
		Help: "Total number of attestations seen by the validators",
	})

	attestationPoolLimit = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "attestation_pool_limit",
		Help: "The limit of the attestation pool",
	})
	attestationPoolSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "attestation_pool_size",
		Help: "The current size of the attestation pool",
	})
)

func reportVoteMetrics(index uint64, block *ethpb.BeaconBlock) {
	// Don't update vote metrics if the incoming block is nil.
	if block == nil {
		return
	}

	validatorLastVoteGauge.WithLabelValues(
		"v" + strconv.Itoa(int(index))).Set(float64(block.Slot))
}
