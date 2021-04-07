package tests

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/openshift-kni/node-label-operator/api/v1beta1"
	. "github.com/openshift-kni/node-label-operator/pkg/test"
)

// Note: this file hasn't the _test.go postfix because it is reused by e2e tests,
// and _test.go files are only compiled if their own package is under test.

var _ = FDescribe("Node test", func() {

	var k8sClient client.Client

	BeforeEach(func() {

		k8sClient = *K8sClient // from test package

	})

	It("test", func() {

		node := v1.Node{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Node",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "foo",
			},
		}

		labels := &v1beta1.Labels{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "test",
			},
			Spec: v1beta1.LabelsSpec{
				NodeNamePatterns: []string{},
				Labels:           make(map[string]string),
				Node:             &node,
			},
		}

		Expect(k8sClient.Create(context.Background(), labels)).To(Succeed())

		updated := &v1beta1.Labels{}
		Expect(k8sClient.Get(context.Background(), client.ObjectKeyFromObject(labels), updated))
		Expect(updated.Spec.Node.Name).To(Equal("foo"))

		updated.Spec.Node.Name = "bar"
		Expect(k8sClient.Update(context.Background(), updated))

		updated2 := &v1beta1.Labels{}
		Expect(k8sClient.Get(context.Background(), client.ObjectKeyFromObject(labels), updated2))
		Expect(updated2.Spec.Node.Name).To(Equal("bar"))

	})

})
