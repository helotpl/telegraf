package nullvalue

import "github.com/influxdata/telegraf"

var sampleConfig = `TODO`

type Nullvalue struct {
}

func (n *Nullvalue) SampleConfig() string {
	return sampleConfig
}

func (n *Nullvalue) Description() string {
	return "Remove values from metrics that are considered as null"
}

func (n *Nullvalue) Apply(in ...telegraf.Metric) []telegraf.Metric {

	return in
}
