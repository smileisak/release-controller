/*
Copyright 2022.

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

package controllers

import (
    "context"

    "k8s.io/apimachinery/pkg/runtime"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/log"

    releasecontrolleriov1alpha1 "github.com/smileisak/release-controller/api/v1alpha1"
    apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// ReleaseStoreReconciler reconciles a ReleaseStore object
type ReleaseStoreReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=release-controller.io,resources=releasestores,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=release-controller.io,resources=releasestores/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=release-controller.io,resources=releasestores/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// the ReleaseStore object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *ReleaseStoreReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    l := log.FromContext(ctx)
    var rs releasecontrolleriov1alpha1.ReleaseStore

    err := r.Get(ctx, req.NamespacedName, &rs)
    if apierrors.IsNotFound(err) {
        return ctrl.Result{}, nil
    } else if err != nil {
        l.Error(err, "unable to get SecretStore")
        return ctrl.Result{}, err
    }

    return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ReleaseStoreReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&releasecontrolleriov1alpha1.ReleaseStore{}).
        Complete(r)
}
