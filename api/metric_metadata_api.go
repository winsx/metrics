// Copyright 2015 Square Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package api holds common data types and public interface exposed by the indexer library.
// Refer to the doc
// https://docs.google.com/a/squareup.com/document/d/1k0Wgi2wnJPQoyDyReb9dyIqRrD8-v0u8hz37S282ii4/edit
// for the terminology.

package api

type MetricMetadataConfig struct {
	// Location of conversion rules. All *.yaml files in here will be loaded.
	ConversionRulesPath string `yaml:"conversion_rules_path"`

	// Database configurations
	// mostly cassandra configurations from
	// https://github.com/gocql/gocql/blob/master/cluster.go
	Hosts    []string `yaml:"hosts"`
	Keyspace string   `yaml:"keyspace"`
}

type MetricMetadataAPI interface {
	// NewMetricMetadataAPI(config MetricMetadataConfig) error
	// AddMetric adds the metric to the system.
	AddMetric(metric TaggedMetric) error
	// Bulk metrics addition
	AddMetrics(metric []TaggedMetric) error
	// RemoveMetric removes the metric from the system.
	RemoveMetric(metric TaggedMetric) error
	// For a given MetricKey, retrieve all the tagsets associated with it.
	GetAllTags(metricKey MetricKey) ([]TagSet, error)
	// GetAllMetrics returns all metrics managed by the system.
	GetAllMetrics() ([]MetricKey, error)
	// For a given tag key-value pair, obtain the list of all the MetricKeys
	// associated with them.
	GetMetricsForTag(tagKey, tagValue string) ([]MetricKey, error)
}
