/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloud

import (
	"context"

	cloudv1 "github.com/KETI-ExaScale/exascale-resource-controller/apis/cloud/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClusterReconciler reconciles a Cluster object
type ClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cloud.exascale.keti,resources=clusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cloud.exascale.keti,resources=clusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cloud.exascale.keti,resources=clusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// MetricPort := ""
	// _ = log.FromContext(ctx)

	// currentCluster := &cloudv1.Cluster{}
	// nodeList := corev1.NodeList{}

	// err := r.Client.Get(ctx, req.NamespacedName, currentCluster)
	// if err != nil {
	// 	klog.Errorln(err)
	// }

	// err = r.Client.List(ctx, &nodeList)
	// if err != nil {
	// 	klog.Errorln(err)
	// }

	// for _, nodes := range nodeList.Items {

	// }

	// r.Client.Patch()

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cloudv1.Cluster{}).
		Complete(r)
}
