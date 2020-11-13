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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var appoperatorlog = logf.Log.WithName("appoperator-resource")

func (r *AppOperator) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-webapp-my-domain-v1-appoperator,mutating=true,failurePolicy=fail,groups=webapp.my.domain,resources=appoperators,verbs=create;update,versions=v1,name=webhook.pingcap.io

var _ webhook.Defaulter = &AppOperator{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *AppOperator) Default() {
	appoperatorlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.

}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-webapp-my-domain-v1-appoperator,mutating=false,failurePolicy=fail,groups=webapp.my.domain,resources=appoperators,versions=v1,name=vappoperator.kb.io

var _ webhook.Validator = &AppOperator{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AppOperator) ValidateCreate() error {
	appoperatorlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.

	appoperatorlog.Info("Create", "address", r.Spec.Address)
	appoperatorlog.Info("Create", "maxbackups", r.Spec.MaxBackups)
	appoperatorlog.Info("Create", "StorageSize", r.Spec.StorageSize)
	appoperatorlog.Info("Create", "Type", r.Spec.Type)
	appoperatorlog.Info("Create", "User", r.Spec.User)
	appoperatorlog.Info("Create", "Limits", r.Spec.Limits)

	err := Validation(r, appoperatorlog)
	if err != nil {
		appoperatorlog.Info("validate failed")
		return err
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AppOperator) ValidateUpdate(old runtime.Object) error {
	appoperatorlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.

	err := Validation(r, appoperatorlog)
	if err != nil {
		appoperatorlog.Info("validate failed")
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AppOperator) ValidateDelete() error {
	appoperatorlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
