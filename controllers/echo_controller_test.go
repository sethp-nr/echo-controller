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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"

	examplev1 "github.com/sethp-nr/echo-controller/api/v1"
)

var _ = Describe("Echo Controller", func() {
	It("Echoes requests", func() {
		ctx := context.Background()
		echo := examplev1.Echo{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-echo",
				Namespace:    "default",
			},
			Spec: examplev1.EchoSpec{
				Ping: "pong",
			},
		}

		Expect(k8sClient.Create(ctx, &echo)).To(Succeed())

		name := types.NamespacedName{Namespace: echo.Namespace, Name: echo.Name}

		Expect((&EchoReconciler{
			Client:   k8sClient,
			Recorder: record.NewFakeRecorder(1),
			Scheme:   scheme.Scheme,
		}).Reconcile(ctrl.Request{NamespacedName: name})).To(Equal(ctrl.Result{}))

		response := examplev1.Echo{}
		Expect(k8sClient.Get(ctx, name, &response)).To(Succeed())
		Expect(response.Status.Pong).To(Equal(echo.Spec.Ping), "Expected reconcile to echo request in response")
	})
})
