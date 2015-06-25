// Copyright 2014 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "fleet"
)

type engineFailure string
type duration string
type operation string

const (
	MachineLeft engineFailure = "machine_left"
	UnitRun     engineFailure = "unable_run_unit"
	JobInactive engineFailure = "job_inactive"

	Reconcile duration = "reconcile"

	Get operation = "get"
	Set operation = "get"
)

var (
	engineScheduleFailureCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: "engine",
		Name:      "schedule_failure_count_total",
		Help:      "Counter of scheduling failures.",
	}, []string{"type"})

	engineScheduleDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: Namespace,
		Subsystem: "engine",
		Name:      "schedule_duration_second",
		Help:      "Historgram of time (in seconds) each schedule round takes.",
	}, []string{"type"})

	engineLeader = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Subsystem: "engine",
		Name:      "leader",
		Help:      "Current fleet leader machine.",
	}, []string{"machine"})

	agentScheduleDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: Namespace,
		Subsystem: "agent",
		Name:      "schedule_duration_second",
		Help:      "Historgram of time (in seconds) each schedule round takes.",
	}, []string{"type"})

	registryDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: Namespace,
		Subsystem: "registry",
		Name:      "duration_second",
		Help:      "Historgram of time (in seconds) each schedule round takes.",
	}, []string{"ops"})
)

func init() {
	prometheus.MustRegister(engineScheduleFailureCount)
	prometheus.MustRegister(engineScheduleDuration)
	prometheus.MustRegister(engineLeader)
	prometheus.MustRegister(registryDuration)
}

func EngineScheduleFailure(reason engineFailure) {
	engineScheduleFailureCount.WithLabelValues(string(reason)).Inc()
}

func EngineScheduleDuration(what duration, start time.Time) {
	engineScheduleDuration.WithLabelValues(string(what)).Observe(float64(time.Since(start)) / float64(time.Second))
}

func AgentScheduleDuration(what duration, start time.Time) {
	agentScheduleDuration.WithLabelValues(string(what)).Observe(float64(time.Since(start)) / float64(time.Second))
}

func EngineLeader(machine string) {
	engineLeader.WithLabelValues(machine).Set(float64(1))
}

func RegistryDuration(ops operation, start time.Time) {
	registryDuration.WithLabelValues(string(ops)).Observe(float64(time.Since(start)) / float64(time.Second))
}
