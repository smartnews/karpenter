/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	NodeSubsystem      = "nodes"
	nodeClaimSubsystem = "nodeclaims"
)

var (
	NodeClaimsCreatedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: nodeClaimSubsystem,
			Name:      "created",
			Help:      "Number of nodeclaims created in total by Karpenter. Labeled by reason the nodeclaim was created and the owning nodepool.",
		},
		[]string{
			ReasonLabel,
			NodePoolLabel,
			CapacityTypeLabel,
		},
	)
	NodeClaimsTerminatedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: nodeClaimSubsystem,
			Name:      "terminated",
			Help:      "Number of nodeclaims terminated in total by Karpenter. Labeled by reason the nodeclaim was terminated and the owning nodepool.",
		},
		[]string{
			ReasonLabel,
			NodePoolLabel,
			CapacityTypeLabel,
		},
	)
	NodeClaimsLaunchedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: nodeClaimSubsystem,
			Name:      "launched",
			Help:      "Number of nodeclaims launched in total by Karpenter. Labeled by the owning nodepool.",
		},
		[]string{
			NodePoolLabel,
		},
	)
	NodeClaimsRegisteredCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: nodeClaimSubsystem,
			Name:      "registered",
			Help:      "Number of nodeclaims registered in total by Karpenter. Labeled by the owning nodepool.",
		},
		[]string{
			NodePoolLabel,
		},
	)
	NodeClaimsInitializedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: nodeClaimSubsystem,
			Name:      "initialized",
			Help:      "Number of nodeclaims initialized in total by Karpenter. Labeled by the owning nodepool.",
		},
		[]string{
			NodePoolLabel,
		},
	)
	NodeClaimsDisruptedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: nodeClaimSubsystem,
			Name:      "disrupted",
			Help:      "Number of nodeclaims disrupted in total by Karpenter. Labeled by disruption type of the nodeclaim and the owning nodepool.",
		},
		[]string{
			TypeLabel,
			NodePoolLabel,
		},
	)
	NodeClaimsDriftedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: nodeClaimSubsystem,
			Name:      "drifted",
			Help:      "Number of nodeclaims drifted reasons in total by Karpenter. Labeled by drift type of the nodeclaim and the owning nodepool.",
		},
		[]string{
			TypeLabel,
			NodePoolLabel,
		},
	)
	NodesCreatedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: NodeSubsystem,
			Name:      "created",
			Help:      "Number of nodes created in total by Karpenter. Labeled by owning nodepool.",
		},
		[]string{
			NodePoolLabel,
		},
	)
	NodesTerminatedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: NodeSubsystem,
			Name:      "terminated",
			Help:      "Number of nodes terminated in total by Karpenter. Labeled by owning nodepool.",
		},
		[]string{
			NodePoolLabel,
		},
	)
)

func init() {
	crmetrics.Registry.MustRegister(NodeClaimsCreatedCounter, NodeClaimsTerminatedCounter, NodeClaimsLaunchedCounter,
		NodeClaimsRegisteredCounter, NodeClaimsInitializedCounter, NodeClaimsDisruptedCounter, NodeClaimsDriftedCounter,
		NodesCreatedCounter, NodesTerminatedCounter)
}
