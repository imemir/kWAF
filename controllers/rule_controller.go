package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kwafv1alpha1 "kwaf/api/v1alpha1"
	"kwaf/internal/waf"
)

// RuleReconciler reconciles Rule CRDs and keeps the WAF engine in sync.
type RuleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Engine *waf.Engine
}

func (r *RuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var rule kwafv1alpha1.Rule
	if err := r.Get(ctx, req.NamespacedName, &rule); err != nil {
		if errors.IsNotFound(err) {
			r.Engine.DeleteRule(req.Name)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	compiled, err := waf.NewRule(rule.Spec.ID, rule.Spec.Pattern, rule.Spec.Action)
	if err != nil {
		logger.Error(err, "failed to compile rule")
		return ctrl.Result{}, nil
	}
	r.Engine.UpsertRule(compiled)
	return ctrl.Result{}, nil
}

func (r *RuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kwafv1alpha1.Rule{}).
		Complete(r)
}

// SetupRuleController adds the rule controller to the manager with a new engine.
func SetupRuleController(mgr ctrl.Manager) error {
	rec := &RuleReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Engine: waf.NewEngine(),
	}
	return rec.SetupWithManager(mgr)
}
