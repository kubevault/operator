/*
Copyright AppsCode Inc. and Contributors
Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package statefulset

import (
	"context"
	"fmt"
	conapi "kubevault.dev/apimachinery/apis"

	api "kubevault.dev/apimachinery/apis/kubevault/v1alpha1"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kmapi "kmodules.xyz/client-go/api/v1"
	dmcond "kmodules.xyz/client-go/dynamic/conditions"
)

type vaultserverInfo struct {
	opts          dmcond.DynamicOptions
	replicasReady bool
	msg           string
}

func (c *Controller) extractVaultserverInfo(sts *apps.StatefulSet) (*vaultserverInfo, error) {
	// read the controlling owner
	owner := metav1.GetControllerOf(sts)
	if owner == nil {
		return nil, fmt.Errorf("StatefulSet %s/%s has no controlling owner", sts.Namespace, sts.Name)
	}
	gv, err := schema.ParseGroupVersion(owner.APIVersion)
	if err != nil {
		return nil, err
	}
	vsInfo := &vaultserverInfo{
		opts: dmcond.DynamicOptions{
			Client:    c.DynamicClient,
			Kind:      owner.Kind,
			Name:      owner.Name,
			Namespace: sts.Namespace,
		},
	}
	vsInfo.opts.GVR = schema.GroupVersionResource{
		Group:   gv.Group,
		Version: gv.Version,
	}
	switch owner.Kind {
	case api.ResourceKindVaultServer:
		vsInfo.opts.GVR.Resource = api.ResourceKindVaultServer
		vs, err := c.VSClient.KubevaultV1alpha1().VaultServers(vsInfo.opts.Namespace).Get(context.TODO(), vsInfo.opts.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		vsInfo.replicasReady, vsInfo.msg, err = vs.ReplicasAreReady(c.StsLister)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown kind: %s", owner.Kind)
	}
	return vsInfo, nil
}

func (c *Controller) ensureReadyReplicasCond(vsInfo *vaultserverInfo) error {
	vsCond := kmapi.Condition{
		Type:    conapi.VaultserverReplicaReady,
		Message: vsInfo.msg,
	}

	if vsInfo.replicasReady {
		vsCond.Status = core.ConditionTrue
		vsCond.Reason = conapi.AllReplicasAreReady
	} else {
		vsCond.Status = core.ConditionFalse
		vsCond.Reason = conapi.SomeReplicasAreNotReady
	}

	// Add "ReplicasReady" condition to the respective vaultserver CR
	return vsInfo.opts.SetCondition(vsCond)
}
