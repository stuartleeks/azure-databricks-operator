/*
Copyright 2019 microsoft.

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

package controllers

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	successMetric = "success"
	failureMetric = "failure"
)

var databricksRequestHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "databricks_request_duration_seconds",
	Help: "Duration of upstream calls to Databricks REST service endpoints",
}, []string{"object_type", "action"})

var databricksRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "databricks_request_total",
	Help: "Counter of upstream calls to Databricks REST service endpoints",
}, []string{"object_type", "action", "outcome"})

func trackExecutionMetrics(f func() error, objectType string, action string) error {
	labels := prometheus.Labels{"object_type": objectType, "action": action}
	observer := databricksRequestHistogram.With(labels)
	timer := prometheus.NewTimer(observer)
	defer timer.ObserveDuration()
	
	result := f()

	if result != nil {
		labels["outcome"] = successMetric
	} else {
		labels["outcome"] = failureMetric
	}

	databricksRequestCounter.With(labels).Inc()

	return result
}

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(databricksRequestHistogram, databricksRequestCounter)
}
