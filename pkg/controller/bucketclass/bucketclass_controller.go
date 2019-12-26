package bucketclass

import (
	nbv1 "github.com/noobaa/noobaa-operator/v2/pkg/apis/noobaa/v1alpha1"
	"github.com/noobaa/noobaa-operator/v2/pkg/bucketclass"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Add creates a Controller and adds it to the Manager.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {

	// Create a controller that runs reconcile on noobaa bucket class

	c, err := controller.New("noobaa-controller", mgr, controller.Options{
		MaxConcurrentReconciles: 1,
		Reconciler: reconcile.Func(
			func(req reconcile.Request) (reconcile.Result, error) {
				return bucketclass.NewReconciler(
					req.NamespacedName,
					mgr.GetClient(),
					mgr.GetScheme(),
					mgr.GetEventRecorderFor("noobaa-operator"),
				).Reconcile()
			}),
	})
	if err != nil {
		return err
	}

	// Watch for changes on resources to trigger reconcile

	err = c.Watch(&source.Kind{Type: &nbv1.BucketClass{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// err = c.Watch(&source.Kind{Type: &nbv1.BackingStore{}}, &handler.EnqueueRequestsFromMapFunc{
	// 	ToRequests: handler.ToRequestsFunc(func(obj handler.MapObject) []reconcile.Request {
	// 		return nil
	// 	}),
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}
