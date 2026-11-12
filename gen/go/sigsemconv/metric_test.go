// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/SigNoz/signoz-semantic-conventions/gen/go/sigsemconv/meterconv"
	"github.com/SigNoz/signoz-semantic-conventions/gen/go/sigsemconv/spanmetricsconv"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

func TestMetricConstructorOptions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		run  func(*testing.T)
	}{
		{"meterconv/CPUNodeAllocatableUsage", constructorTest(meterconv.NewCPUNodeAllocatableUsage, "signoz.meter.cpu.node.allocatable.usage", "The amount of CPU allocatable on a Kubernetes node that runs SigNoz.", "{cpu}", false, nil)},
		{"meterconv/CPUNodeAllocatableUsageObservable", constructorTest(meterconv.NewCPUNodeAllocatableUsageObservable, "signoz.meter.cpu.node.allocatable.usage", "The amount of CPU allocatable on a Kubernetes node that runs SigNoz.", "{cpu}", false, nil)},
		{"meterconv/LogCount", constructorTest(meterconv.NewLogCount, "signoz.meter.log.count", "The number of log records observed.", "1", false, nil)},
		{"meterconv/LogCountObservable", constructorTest(meterconv.NewLogCountObservable, "signoz.meter.log.count", "The number of log records observed.", "1", false, nil)},
		{"meterconv/LogSize", constructorTest(meterconv.NewLogSize, "signoz.meter.log.size", "The size of log records observed.", "By", false, nil)},
		{"meterconv/LogSizeObservable", constructorTest(meterconv.NewLogSizeObservable, "signoz.meter.log.size", "The size of log records observed.", "By", false, nil)},
		{"meterconv/MetricDatapointCount", constructorTest(meterconv.NewMetricDatapointCount, "signoz.meter.metric.datapoint.count", "The number of metric data points indexed.", "1", false, nil)},
		{"meterconv/MetricDatapointCountObservable", constructorTest(meterconv.NewMetricDatapointCountObservable, "signoz.meter.metric.datapoint.count", "The number of metric data points indexed.", "1", false, nil)},
		{"meterconv/MetricDatapointIngestedCount", constructorTest(meterconv.NewMetricDatapointIngestedCount, "signoz.meter.metric.datapoint.ingested.count", "The number of metric data points ingested.", "1", false, nil)},
		{"meterconv/MetricDatapointIngestedCountObservable", constructorTest(meterconv.NewMetricDatapointIngestedCountObservable, "signoz.meter.metric.datapoint.ingested.count", "The number of metric data points ingested.", "1", false, nil)},
		{"meterconv/MetricDatapointSize", constructorTest(meterconv.NewMetricDatapointSize, "signoz.meter.metric.datapoint.size", "The size of metric data points observed.", "By", false, nil)},
		{"meterconv/MetricDatapointSizeObservable", constructorTest(meterconv.NewMetricDatapointSizeObservable, "signoz.meter.metric.datapoint.size", "The size of metric data points observed.", "By", false, nil)},
		{"meterconv/SpanCount", constructorTest(meterconv.NewSpanCount, "signoz.meter.span.count", "The number of spans observed.", "1", false, nil)},
		{"meterconv/SpanCountObservable", constructorTest(meterconv.NewSpanCountObservable, "signoz.meter.span.count", "The number of spans observed.", "1", false, nil)},
		{"meterconv/SpanSize", constructorTest(meterconv.NewSpanSize, "signoz.meter.span.size", "The size of spans observed.", "By", false, nil)},
		{"meterconv/SpanSizeObservable", constructorTest(meterconv.NewSpanSizeObservable, "signoz.meter.span.size", "The size of spans observed.", "By", false, nil)},
		{"spanmetricsconv/CallsTotal", constructorTest(spanmetricsconv.NewCallsTotal, "signoz_calls_total", "The number of spans observed.", "", false, nil)},
		{"spanmetricsconv/CallsTotalObservable", constructorTest(spanmetricsconv.NewCallsTotalObservable, "signoz_calls_total", "The number of spans observed.", "", false, nil)},
		{"spanmetricsconv/DBLatencyCount", constructorTest(spanmetricsconv.NewDBLatencyCount, "signoz_db_latency_count", "The number of database calls observed.", "1", false, nil)},
		{"spanmetricsconv/DBLatencyCountObservable", constructorTest(spanmetricsconv.NewDBLatencyCountObservable, "signoz_db_latency_count", "The number of database calls observed.", "1", false, nil)},
		{"spanmetricsconv/DBLatencySum", constructorTest(spanmetricsconv.NewDBLatencySum, "signoz_db_latency_sum", "The total duration of database calls observed.", "1", false, nil)},
		{"spanmetricsconv/DBLatencySumObservable", constructorTest(spanmetricsconv.NewDBLatencySumObservable, "signoz_db_latency_sum", "The total duration of database calls observed.", "1", false, nil)},
		{"spanmetricsconv/ExternalCallLatencyCount", constructorTest(spanmetricsconv.NewExternalCallLatencyCount, "signoz_external_call_latency_count", "The number of external calls observed.", "1", false, nil)},
		{"spanmetricsconv/ExternalCallLatencyCountObservable", constructorTest(spanmetricsconv.NewExternalCallLatencyCountObservable, "signoz_external_call_latency_count", "The number of external calls observed.", "1", false, nil)},
		{"spanmetricsconv/ExternalCallLatencySum", constructorTest(spanmetricsconv.NewExternalCallLatencySum, "signoz_external_call_latency_sum", "The total duration of external calls observed.", "1", false, nil)},
		{"spanmetricsconv/ExternalCallLatencySumObservable", constructorTest(spanmetricsconv.NewExternalCallLatencySumObservable, "signoz_external_call_latency_sum", "The total duration of external calls observed.", "1", false, nil)},
		{"spanmetricsconv/Latency", constructorTest(spanmetricsconv.NewLatency, "signoz_latency", "The duration of spans observed.", "ms", true, nil)},
	}
	for _, test := range tests {
		t.Run(test.name, test.run)
	}
}

func constructorTest[O, I any](
	newInstrument func(metric.Meter, ...O) (I, error),
	name, description, unit string,
	histogram bool,
	advice []float64,
) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		for _, callerOptions := range []bool{false, true} {
			label := "defaults"
			if callerOptions {
				label = "caller_options"
			}
			t.Run(label, func(t *testing.T) {
				m := &constructorMeter{Meter: noop.NewMeterProvider().Meter("test")}
				var options, storage, before []O
				wantBoundaries := advice
				if callerOptions {
					// Keep sentinels beyond len(options) to detect writes into spare capacity.
					storage = make([]O, 16)
					for j := range storage {
						storage[j] = any(metric.WithUnit("sentinel")).(O)
					}
					options = storage[:2]
					options[0] = any(metric.WithDescription("caller description")).(O)
					options[1] = any(metric.WithUnit("caller unit")).(O)
					if histogram {
						wantBoundaries = []float64{1, 2, 3}
						options = append(options, any(metric.WithExplicitBucketBoundaries(wantBoundaries...)).(O))
					}
					before = slices.Clone(storage)
				}
				if _, err := newInstrument(m, options...); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(storage, before) {
					t.Error("constructor changed caller option storage")
				}
				if m.config == nil {
					t.Fatal("constructor did not create an instrument")
				}
				if m.name != name {
					t.Errorf("name = %q, want %q", m.name, name)
				}
				if got := m.config.Description(); got != description {
					t.Errorf("description = %q, want %q", got, description)
				}
				if got := m.config.Unit(); got != unit {
					t.Errorf("unit = %q, want %q", got, unit)
				}
				if histogram {
					config, ok := m.config.(interface{ ExplicitBucketBoundaries() []float64 })
					if !ok {
						t.Fatal("constructor did not create a histogram")
					}
					if got := config.ExplicitBucketBoundaries(); !slices.Equal(got, wantBoundaries) {
						t.Errorf("bucket boundaries = %v, want %v", got, wantBoundaries)
					}
				}
			})
		}
	}
}

type constructorMeter struct {
	metric.Meter
	name   string
	config interface {
		Description() string
		Unit() string
	}
}

func (m *constructorMeter) Int64Counter(name string, options ...metric.Int64CounterOption) (metric.Int64Counter, error) {
	m.name = name
	m.config = metric.NewInt64CounterConfig(options...)
	return m.Meter.Int64Counter(name, options...)
}

func (m *constructorMeter) Int64UpDownCounter(name string, options ...metric.Int64UpDownCounterOption) (metric.Int64UpDownCounter, error) {
	m.name = name
	m.config = metric.NewInt64UpDownCounterConfig(options...)
	return m.Meter.Int64UpDownCounter(name, options...)
}

func (m *constructorMeter) Int64Histogram(name string, options ...metric.Int64HistogramOption) (metric.Int64Histogram, error) {
	m.name = name
	m.config = metric.NewInt64HistogramConfig(options...)
	return m.Meter.Int64Histogram(name, options...)
}

func (m *constructorMeter) Int64Gauge(name string, options ...metric.Int64GaugeOption) (metric.Int64Gauge, error) {
	m.name = name
	m.config = metric.NewInt64GaugeConfig(options...)
	return m.Meter.Int64Gauge(name, options...)
}

func (m *constructorMeter) Int64ObservableCounter(name string, options ...metric.Int64ObservableCounterOption) (metric.Int64ObservableCounter, error) {
	m.name = name
	m.config = metric.NewInt64ObservableCounterConfig(options...)
	return m.Meter.Int64ObservableCounter(name, options...)
}

func (m *constructorMeter) Int64ObservableUpDownCounter(name string, options ...metric.Int64ObservableUpDownCounterOption) (metric.Int64ObservableUpDownCounter, error) {
	m.name = name
	m.config = metric.NewInt64ObservableUpDownCounterConfig(options...)
	return m.Meter.Int64ObservableUpDownCounter(name, options...)
}

func (m *constructorMeter) Int64ObservableGauge(name string, options ...metric.Int64ObservableGaugeOption) (metric.Int64ObservableGauge, error) {
	m.name = name
	m.config = metric.NewInt64ObservableGaugeConfig(options...)
	return m.Meter.Int64ObservableGauge(name, options...)
}

func (m *constructorMeter) Float64Counter(name string, options ...metric.Float64CounterOption) (metric.Float64Counter, error) {
	m.name = name
	m.config = metric.NewFloat64CounterConfig(options...)
	return m.Meter.Float64Counter(name, options...)
}

func (m *constructorMeter) Float64UpDownCounter(name string, options ...metric.Float64UpDownCounterOption) (metric.Float64UpDownCounter, error) {
	m.name = name
	m.config = metric.NewFloat64UpDownCounterConfig(options...)
	return m.Meter.Float64UpDownCounter(name, options...)
}

func (m *constructorMeter) Float64Histogram(name string, options ...metric.Float64HistogramOption) (metric.Float64Histogram, error) {
	m.name = name
	m.config = metric.NewFloat64HistogramConfig(options...)
	return m.Meter.Float64Histogram(name, options...)
}

func (m *constructorMeter) Float64Gauge(name string, options ...metric.Float64GaugeOption) (metric.Float64Gauge, error) {
	m.name = name
	m.config = metric.NewFloat64GaugeConfig(options...)
	return m.Meter.Float64Gauge(name, options...)
}

func (m *constructorMeter) Float64ObservableCounter(name string, options ...metric.Float64ObservableCounterOption) (metric.Float64ObservableCounter, error) {
	m.name = name
	m.config = metric.NewFloat64ObservableCounterConfig(options...)
	return m.Meter.Float64ObservableCounter(name, options...)
}

func (m *constructorMeter) Float64ObservableUpDownCounter(name string, options ...metric.Float64ObservableUpDownCounterOption) (metric.Float64ObservableUpDownCounter, error) {
	m.name = name
	m.config = metric.NewFloat64ObservableUpDownCounterConfig(options...)
	return m.Meter.Float64ObservableUpDownCounter(name, options...)
}

func (m *constructorMeter) Float64ObservableGauge(name string, options ...metric.Float64ObservableGaugeOption) (metric.Float64ObservableGauge, error) {
	m.name = name
	m.config = metric.NewFloat64ObservableGaugeConfig(options...)
	return m.Meter.Float64ObservableGauge(name, options...)
}
