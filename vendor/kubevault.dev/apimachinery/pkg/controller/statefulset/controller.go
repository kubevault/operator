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
	"fmt"

	conapi "kubevault.dev/apimachinery/apis"
	"kubevault.dev/apimachinery/apis/kubevault"
	vs_cs "kubevault.dev/apimachinery/client/clientset/versioned"
	amc "kubevault.dev/apimachinery/pkg/controller"

	apps "k8s.io/api/apps/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	core_util "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/client-go/tools/queue"
)

type Controller struct {
	*amc.Controller
	*amc.Config
}

func NewController(
	config *amc.Config,
	client kubernetes.Interface,
	vsClient vs_cs.Interface,
	dmClient dynamic.Interface,
) *Controller {
	return &Controller{
		Controller: &amc.Controller{
			Client:        client,
			VSClient:      vsClient,
			DynamicClient: dmClient,
		},
		Config: config,
	}
}

func (c *Controller) InitStsWatcher() {
	klog.Infoln("Initializing StatefulSet watcher.....")
	// Initialize RestoreSession Watcher
	c.StsInformer = c.KubeInformerFactory.Apps().V1().StatefulSets().Informer()
	// Todo: Add ResourceKindStatefulSet in apimachinery - added in conapi
	c.StsQueue = queue.New(conapi.ResourceKindStatefulSet, c.MaxNumRequeues, c.NumThreads, c.processStatefulSet)
	c.StsLister = c.KubeInformerFactory.Apps().V1().StatefulSets().Lister()
	c.StsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if sts, ok := obj.(*apps.StatefulSet); ok {
				c.enqueueOnlyKubeVaultSts(sts)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			if sts, ok := newObj.(*apps.StatefulSet); ok {
				c.enqueueOnlyKubeVaultSts(sts)
			}
		},
		DeleteFunc: func(obj interface{}) {
			if sts, ok := obj.(*apps.StatefulSet); ok {
				ok, _, err := core_util.IsOwnerOfGroup(metav1.GetControllerOf(sts), kubevault.GroupName)
				if err != nil || !ok {
					klog.Warningln(err)
					return
				}
				dbInfo, err := c.extractVaultserverInfo(sts)
				if err != nil {
					if !kerr.IsNotFound(err) {
						klog.Warningf("failed to extract vaultserver info from StatefulSet: %s/%s. Reason: %v", sts.Namespace, sts.Name, err)
					}
					return
				}
				err = c.ensureReadyReplicasCond(dbInfo)
				if err != nil {
					klog.Warningf("failed to update ReadyReplicas condition. Reason: %v", err)
					return
				}
			}
		},
	})
}

func (c *Controller) enqueueOnlyKubeVaultSts(sts *apps.StatefulSet) {
	// only enqueue if the controlling owner is a kubevault resource
	ok, _, err := core_util.IsOwnerOfGroup(metav1.GetControllerOf(sts), kubevault.GroupName)
	if err != nil {
		klog.Warningf("failed to enqueue StatefulSet: %s/%s. Reason: %v", sts.Namespace, sts.Name, err)
		return
	}
	if ok {
		queue.Enqueue(c.StsQueue.GetQueue(), cache.ExplicitKey(sts.Namespace+"/"+sts.Name))
	}
}

func (c *Controller) processStatefulSet(key string) error {
	klog.Infof("Started processing, key: %v", key)
	obj, exists, err := c.StsInformer.GetIndexer().GetByKey(key)
	if err != nil {
		klog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		klog.V(5).Infof("StatefulSet %s does not exist anymore", key)
	} else {
		sts := obj.(*apps.StatefulSet).DeepCopy()
		dbInfo, err := c.extractVaultserverInfo(sts)
		if err != nil {
			return fmt.Errorf("failed to extract vaultserver info from StatefulSet: %s/%s. Reason: %v", sts.Namespace, sts.Name, err)
		}
		return c.ensureReadyReplicasCond(dbInfo)
	}
	return nil
}
