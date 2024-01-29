/*
Copyright 2023 The Kubernetes Authors.

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

package controller

import (
	"sync"

	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
)

const (
	namespace = "apiserver"
	subsystem = "clusterip_repair"
)

var (
	// clusterIPRepairIPErrors indicates the number of errors found by the repair loop
	// divided by the type of error:
	// leak, repair, full, outOfRange, duplicate, invalid, unknown
	clusterIPRepairIPErrors = metrics.NewCounterVec(
		&metrics.CounterOpts{
			Namespace:      namespace,
			Subsystem:      subsystem,
			Name:           "ip_errors_total",
			Help:           "Number of errors detected on clusterips by the repair loop broken down by type of error: leak, repair, full, outOfRange, duplicate, unknown, invalid",
			StabilityLevel: metrics.ALPHA,
		},
		[]string{"type"},
	)
	// clusterIPRepairReconcileErrors indicates the number of times the repair loop has failed to repair
	// the errors it detected.
	clusterIPRepairReconcileErrors = metrics.NewCounter(
		&metrics.CounterOpts{
			Namespace:      namespace,
			Subsystem:      subsystem,
			Name:           "reconcile_errors_total",
			Help:           "Number of reconciliation failures on the clusterip repair reconcile loop",
			StabilityLevel: metrics.ALPHA,
		},
	)
)

var registerMetricsOnce sync.Once

func registerMetrics() {
	registerMetricsOnce.Do(func() {
		legacyregistry.MustRegister(clusterIPRepairIPErrors)
		legacyregistry.MustRegister(clusterIPRepairReconcileErrors)
	})
}
