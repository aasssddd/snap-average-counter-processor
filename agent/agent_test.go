package agent

import (
	"testing"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProcessor(t *testing.T) {
	Convey("Test Processor", t, func() {
		// Convey("Process int metric", func() {
		// 	metrics := []plugin.Metric{
		// 		plugin.Metric{
		// 			Namespace: plugin.NewNamespace("x", "y", "z"),
		// 			Config:    map[string]interface{}{"pw": "123aB"},
		// 			Data:      345678,
		// 			Tags:      map[string]string{"hello": "world"},
		// 			Unit:      "int",
		// 			Timestamp: time.Now(),
		// 		},
		// 	}
		// 	mts, err := p.Process(metrics, plugin.Config{})
		// 	So(mts, ShouldNotBeNil)
		// 	So(err, ShouldBeNil)
		// 	So(mts[0].Data, ShouldEqual, 876543)
		// })

		Convey("Test Process", func() {
			p := NewProcessor()
			cfg := plugin.Config{
				"namespaces":           "default, hyperpilot",
				"filterMetricKeywords": "perc*, intel/docker/spec/size_rw",
			}
			mts := []plugin.Metric{
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "docker", "stats", "cgroups", "cpu_stats", "cpu_shares"),
					Config:    map[string]interface{}{"pw": "123aB"},
					Data:      123,
					Tags:      map[string]string{"io.kubernetes.pod.namespace": "default"},
					Unit:      "int",
					Timestamp: time.Now(),
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "docker", "spec", "size_root"),
					Config:    map[string]interface{}{"pw": "123aB"},
					Data:      456,
					Tags:      map[string]string{"io.kubernetes.pod.namespace": "default"},
					Unit:      "int",
					Timestamp: time.Now(),
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "docker", "spec", "size_rw"),
					Config:    map[string]interface{}{"pw": "123aB"},
					Data:      789,
					Tags:      map[string]string{"io.kubernetes.pod.namespace": "hyperpilot"},
					Unit:      "int",
					Timestamp: time.Now(),
				},
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "docker", "spec", "size_rw"),
					Config:    map[string]interface{}{"pw": "123aB"},
					Data:      789,
					Tags:      map[string]string{"io.kubernetes.pod.namespace": "haha"},
					Unit:      "int",
					Timestamp: time.Now(),
				},
			}
			result, err := p.Process(mts, cfg)
			Convey("Should only process 2 data", func() {
				So(len(result), ShouldEqual, 3)
			})
			Convey("No error returned", func() {
				So(err, ShouldBeNil)
			})

		})
	})
}
