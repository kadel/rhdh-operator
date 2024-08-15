//
// Copyright (c) 2023 Red Hat, Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"

	bsv1 "redhat-developer/red-hat-developer-hub-operator/api/v1alpha2"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DynamicPluginsRootPVCFactory struct{}

func (f DynamicPluginsRootPVCFactory) newBackstageObject() RuntimeObject {
	return &DynamicPluginsPVC{}
}

type DynamicPluginsPVC struct {
	pvc *corev1.PersistentVolumeClaim
}

func init() {
	registerConfig("pvc-dynamic-plugins-root.yaml", DynamicPluginsRootPVCFactory{})
}

func PVCName(backstageName string) string {
	//return utils.GenerateRuntimeObjectName(backstageName, "dynamic-plugins-root")
	// TODO(kadel): use name with backstageName, Deployment needs to be updated
	return "dynamic-plugins-root"
}

// implementation of RuntimeObject interface
func (pvc *DynamicPluginsPVC) Object() client.Object {
	return pvc.pvc
}

func (pvc *DynamicPluginsPVC) setObject(obj client.Object) {
	pvc.pvc = nil
	if obj != nil {
		pvc.pvc = obj.(*corev1.PersistentVolumeClaim)
	}
}

// implementation of RuntimeObject interface
func (pvc *DynamicPluginsPVC) addToModel(model *BackstageModel, _ bsv1.Backstage) (bool, error) {
	if pvc.pvc == nil {
		return false, fmt.Errorf("Backstage Service is not initialized, make sure there is pvc-dynamic-plugins-root.yaml in default or raw configuration")
	}
	model.setRuntimeObject(pvc)

	return true, nil

}

// implementation of RuntimeObject interface
func (pvc *DynamicPluginsPVC) EmptyObject() client.Object {
	return &corev1.PersistentVolumeClaim{}
}

// implementation of RuntimeObject interface
func (pvc *DynamicPluginsPVC) validate(_ *BackstageModel, _ bsv1.Backstage) error {
	return nil
}

func (pvc *DynamicPluginsPVC) setMetaInfo(backstageName string) {
	pvc.pvc.SetName(PVCName(backstageName))
}
