package main

import (
	"flag"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	kwafv1alpha1 "kwaf/api/v1alpha1"
	"kwaf/controllers"
)

func main() {
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	var metricsAddr string
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.Parse()

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:  runtime.NewScheme(),
		Metrics: metricsserver.Options{BindAddress: metricsAddr},
	})
	if err != nil {
		panic(err)
	}

	if err := kwafv1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		panic(err)
	}

	if err := controllers.SetupRuleController(mgr); err != nil {
		panic(err)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		panic(err)
	}
}
