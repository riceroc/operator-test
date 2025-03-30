package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demoappv1alpha1 "github.com/riceroc/k8s-operator-demo/api/v1alpha1"
)

// DemoAppReconciler reconciles a DemoApp object
type DemoAppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demoapp.example.com,resources=demoapps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demoapp.example.com,resources=demoapps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demoapp.example.com,resources=demoapps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *DemoAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	demoApp := &demoappv1alpha1.DemoApp{}

	if err := r.Get(ctx, req.NamespacedName, demoApp); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Log the message from the spec
	log.Info("Reconciling DemoApp", "message", demoApp.Spec.Message)

	// Update the status
	demoApp.Status.Phase = "Reconciled"
	if err := r.Status().Update(ctx, demoApp); err != nil {
		return ctrl.Result{}, err
	}

	// Requeue after 10 seconds
	return ctrl.Result{RequeueAfter: time.Second * 10}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DemoAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demoappv1alpha1.DemoApp{}).
		Complete(r)
} 