/*
Copyright The KubeDB Authors.

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

package v1alpha1

import (
	"kubedb.dev/apimachinery/apis"
	"kubedb.dev/apimachinery/crds"

	"kmodules.xyz/client-go/apiextensions"
)

func (_ MongoDBOpsRequest) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralMongoDBOpsRequest))
}

var _ apis.ResourceInfo = &MongoDBOpsRequest{}

func (m MongoDBOpsRequest) ResourceShortCode() string {
	return ResourceCodeMongoDBOpsRequest
}

func (m MongoDBOpsRequest) ResourceKind() string {
	return ResourceKindMongoDBOpsRequest
}

func (m MongoDBOpsRequest) ResourceSingular() string {
	return ResourceSingularMongoDBOpsRequest
}

func (m MongoDBOpsRequest) ResourcePlural() string {
	return ResourcePluralMongoDBOpsRequest
}

func (m MongoDBOpsRequest) ValidateSpecs() error {
	return nil
}
