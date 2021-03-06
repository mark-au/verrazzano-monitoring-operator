// Copyright (C) 2020, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package metrics

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/verrazzano/verrazzano-monitoring-operator/pkg/constants"
	"go.uber.org/zap"
)

// DanglingPVC GaugeVec
var DanglingPVC = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: constants.MetricsNameSpace,
		Name:      "dangling_pvc",
		Help:      "value tells the dangling pvc exists",
	},
	[]string{"pvc_name", "availability_domain"},
)

// Lock GaugeVec
var Lock = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: constants.MetricsNameSpace,
		Name:      "do_not_sync",
		Help:      "value tells the Lock flag is on",
	},
	[]string{"namespace", "vmo_name"},
)

// RegisterMetrics register Prometheus metrics
func RegisterMetrics() {
	prometheus.MustRegister(DanglingPVC)
	prometheus.MustRegister(Lock)
}

// StartServer starts server for Prometheus metrics
func StartServer(port int) {
	flag.IntVar(&port, "port", port, "Port on which to expose Prometheus metrics")
	flag.Parse()
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/metrics", promhttp.Handler())
	zap.S().Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
