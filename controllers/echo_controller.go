/*


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

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	examplev1 "github.com/sethp-nr/echo-controller/api/v1"
)

// EchoReconciler reconciles a Echo object
type EchoReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=example.x-k8s.io,resources=echoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=example.x-k8s.io,resources=echoes/status,verbs=get;update;patch

func (r *EchoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	var obj examplev1.Echo
	err := r.Get(ctx, req.NamespacedName, &obj)
	if err != nil {
		return ctrl.Result{}, err
	}
	obj.Status.Pong = obj.Spec.Ping
	err = r.Status().Update(ctx, &obj)
	if err != nil {
		return ctrl.Result{}, err
	}
	r.Recorder.Eventf(&obj, corev1.EventTypeNormal, "EchoSuccess", "Successfully echoed request %q", obj.Spec.Ping)

	return ctrl.Result{}, nil
}

func (r *EchoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplev1.Echo{}).
		Complete(r)
}
