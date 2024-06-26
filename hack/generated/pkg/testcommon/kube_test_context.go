/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
)

// TODO: State Annotation parameter should be removed once the interface for Status determined and promoted
// TODO: to genruntime. Same for errorAnnotation
type KubeGlobalContext struct {
	TestContext

	createBaseTestContext func(PerTestContext) (*KubeBaseTestContext, error)

	namespace string
}

func (ctx KubeGlobalContext) Namespace() string {
	return ctx.namespace
}

func NewKubeContext(
	useEnvTest bool,
	recordReplay bool,
	namespace string,
	region string) KubeGlobalContext {

	result := KubeGlobalContext{
		TestContext: NewTestContext(region, recordReplay),
		namespace:   namespace,
	}

	if useEnvTest {
		result.createBaseTestContext = createEnvtestContext
	} else {
		result.createBaseTestContext = createRealKubeContext
	}

	return result
}

type KubeBaseTestContext struct {
	PerTestContext

	KubeConfig *rest.Config
}

func AsOwner(obj metav1.ObjectMeta) genruntime.KnownResourceReference {
	return genruntime.KnownResourceReference{
		Name: obj.Name,
	}
}
