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

package controller

import (
	"time"

	cs "kubevault.dev/apimachinery/client/clientset/versioned"
	vaultinformers "kubevault.dev/apimachinery/client/informers/externalversions"
	amc "kubevault.dev/apimachinery/pkg/controller"
	"kubevault.dev/operator/pkg/eventer"

	pcm "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	auditlib "go.bytebuilders.dev/audit/lib"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	reg_util "kmodules.xyz/client-go/admissionregistration/v1beta1"
	"kmodules.xyz/client-go/discovery"
	"kmodules.xyz/client-go/tools/cli"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1"
)

const (
	mutatingWebhook   = "mutators.kubevault.com"
	validatingWebhook = "validators.kubevault.com"
)

type config struct {
	EnableRBAC              bool
	DockerRegistry          string
	MaxNumRequeues          int
	NumThreads              int
	ResyncPeriod            time.Duration
	EnableValidatingWebhook bool
	EnableMutatingWebhook   bool
}

type Config struct {
	amc.Config

	LicenseFile      string
	ClientConfig     *rest.Config
	KubeClient       kubernetes.Interface
	ExtClient        cs.Interface
	CRDClient        crd_cs.Interface
	AppCatalogClient appcat_cs.AppcatalogV1alpha1Interface
	PromClient       pcm.MonitoringV1Interface
	VSClient         cs.Interface
	Recorder         record.EventRecorder
	DynamicClient    dynamic.Interface
}

func NewConfig(clientConfig *rest.Config) *Config {
	return &Config{
		ClientConfig: clientConfig,
	}
}

func (c *Config) New() (*VaultController, error) {
	if err := discovery.IsDefaultSupportedVersion(c.KubeClient); err != nil {
		return nil, err
	}

	// audit event publisher
	// WARNING: https://stackoverflow.com/a/46275411/244009
	var auditor *auditlib.EventPublisher
	if c.LicenseFile != "" && cli.EnableAnalytics {
		mapper, err := discovery.NewDynamicResourceMapper(c.ClientConfig)
		if err != nil {
			return nil, err
		}
		fn := auditlib.BillingEventCreator{
			Mapper: mapper,
		}
		auditor = auditlib.NewResilientEventPublisher(func() (*auditlib.NatsConfig, error) {
			return auditlib.NewNatsConfig(c.KubeClient.CoreV1().Namespaces(), c.LicenseFile)
		}, mapper, fn.CreateEvent)
	}

	ctrl := &VaultController{
		clientConfig:        c.ClientConfig,
		ctxCancels:          make(map[string]CtxWithCancel),
		finalizerInfo:       NewMapFinalizer(),
		authMethodCtx:       make(map[string]CtxWithCancel),
		kubeClient:          c.KubeClient,
		extClient:           c.ExtClient,
		crdClient:           c.CRDClient,
		promClient:          c.PromClient,
		appCatalogClient:    c.AppCatalogClient,
		kubeInformerFactory: informers.NewSharedInformerFactory(c.KubeClient, c.ResyncPeriod),
		extInformerFactory:  vaultinformers.NewSharedInformerFactory(c.ExtClient, c.ResyncPeriod),
		recorder:            eventer.NewEventRecorder(c.KubeClient, "vault-operator"),
		auditor:             auditor,
	}

	if err := ctrl.ensureCustomResourceDefinitions(); err != nil {
		return nil, err
	}
	if c.EnableMutatingWebhook {
		if err := reg_util.UpdateMutatingWebhookCABundle(c.ClientConfig, mutatingWebhook); err != nil {
			return nil, err
		}
	}
	if c.EnableValidatingWebhook {
		if err := reg_util.UpdateValidatingWebhookCABundle(c.ClientConfig, validatingWebhook); err != nil {
			return nil, err
		}
	}

	// For VaultServer
	ctrl.initVaultServerWatcher()
	// For VaultPolicy
	ctrl.initVaultPolicyWatcher()
	// For VaultPolicyBinding
	ctrl.initVaultPolicyBindingWatcher()

	// For DB manager
	ctrl.initPostgresRoleWatcher()
	ctrl.initMySQLRoleWatcher()
	ctrl.initMongoDBRoleWatcher()
	ctrl.initElasticsearchRoleWatcher()
	ctrl.initDatabaseAccessWatcher()

	// For AWSRole
	ctrl.initAWSRoleWatcher()
	ctrl.initAWSAccessKeyWatcher()

	// For GCPRole
	ctrl.initGCPRoleWatcher()
	ctrl.initGCPAccessKeyWatcher()

	// For AzureRole
	ctrl.initAzureRoleWatcher()
	ctrl.initAzureAccessKeyWatcher()

	// For secretEngine
	ctrl.initSecretEngineWatcher()

	return ctrl, nil
}
