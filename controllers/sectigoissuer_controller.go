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

	api "github.com/CUB-OIT-PE/sectigo-issuer"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	certmanagersectigov1beta1 "sectigo-issuer/api/v1beta1"
)

// SectigoIssuerReconciler reconciles a SectigoIssuer object
type SectigoIssuerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=certmanager.sectigo.oit.colorado.edu,resources=sectigoissuers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certmanager.sectigo.oit.colorado.edu,resources=sectigoissuers/status,verbs=get;update;patch

func (r *SectigoIssuerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sectigoissuer", req.NamespacedName)

	// your logic here
	iss := new(api.StepIssuer)

	return ctrl.Result{}, nil
}

func (r *SectigoIssuerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&certmanagersectigov1beta1.SectigoIssuer{}).
		Complete(r)
}
