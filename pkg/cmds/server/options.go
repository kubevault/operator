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

package server

import (
	"flag"
	"time"

	cs "kubevault.dev/apimachinery/client/clientset/versioned"
	kubevaultinformers "kubevault.dev/apimachinery/client/informers/externalversions"
	sts "kubevault.dev/apimachinery/pkg/controller/statefulset"
	"kubevault.dev/operator/pkg/controller"
	"kubevault.dev/operator/pkg/docker"
	"kubevault.dev/operator/pkg/eventer"

	prom "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	"github.com/spf13/pflag"
	core "k8s.io/api/core/v1"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"kmodules.xyz/client-go/tools/clusterid"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1"
)

type ExtraOptions struct {
	LicenseFile             string
	DockerRegistry          string
	MaxNumRequeues          int
	NumThreads              int
	QPS                     float64
	Burst                   int
	ResyncPeriod            time.Duration
	EnableValidatingWebhook bool
	EnableMutatingWebhook   bool
}

func NewExtraOptions() *ExtraOptions {
	return &ExtraOptions{
		DockerRegistry: docker.ACRegistry,
		MaxNumRequeues: 5,
		NumThreads:     2,
		QPS:            100,
		Burst:          100,
		ResyncPeriod:   10 * time.Minute,
	}
}

func (s *ExtraOptions) AddGoFlags(fs *flag.FlagSet) {
	clusterid.AddGoFlags(fs)

	fs.StringVar(&s.LicenseFile, "license-file", s.LicenseFile, "Path to license file")

	fs.StringVar(&s.DockerRegistry, "docker-registry", s.DockerRegistry, "Docker image registry for sidecar, init-container, check-job, recovery-job and kubectl-job")

	fs.Float64Var(&s.QPS, "qps", s.QPS, "The maximum QPS to the master from this client")
	fs.IntVar(&s.Burst, "burst", s.Burst, "The maximum burst for throttle")
	fs.DurationVar(&s.ResyncPeriod, "resync-period", s.ResyncPeriod, "If non-zero, will re-list this often. Otherwise, re-list will be delayed aslong as possible (until the upstream source closes the watch or times out.")

	fs.BoolVar(&s.EnableMutatingWebhook, "enable-mutating-webhook", s.EnableMutatingWebhook, "If true, enables mutating webhooks for KubeVault CRDs.")
	fs.BoolVar(&s.EnableValidatingWebhook, "enable-validating-webhook", s.EnableValidatingWebhook, "If true, enables validating webhooks for KubeVault CRDs.")
}

func (s *ExtraOptions) AddFlags(fs *pflag.FlagSet) {
	pfs := flag.NewFlagSet("vault-server", flag.ExitOnError)
	s.AddGoFlags(pfs)
	fs.AddGoFlagSet(pfs)
}

func (s *ExtraOptions) ApplyTo(cfg *controller.Config) error {
	var err error

	cfg.LicenseFile = s.LicenseFile
	cfg.MaxNumRequeues = s.MaxNumRequeues
	cfg.NumThreads = s.NumThreads
	cfg.ResyncPeriod = s.ResyncPeriod
	cfg.ClientConfig.QPS = float32(s.QPS)
	cfg.ClientConfig.Burst = s.Burst
	cfg.WatchNamespace = core.NamespaceAll
	cfg.EnableMutatingWebhook = s.EnableMutatingWebhook
	cfg.EnableValidatingWebhook = s.EnableValidatingWebhook

	if cfg.KubeClient, err = kubernetes.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.ExtClient, err = cs.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.CRDClient, err = crd_cs.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.PromClient, err = prom.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.AppCatalogClient, err = appcat_cs.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}
	if cfg.VSClient, err = cs.NewForConfig(cfg.ClientConfig); err != nil {
		return err
	}

	cfg.KubeInformerFactory = informers.NewSharedInformerFactory(cfg.KubeClient, cfg.ResyncPeriod)
	cfg.KubevaultInformerFactory = kubevaultinformers.NewSharedInformerFactory(cfg.VSClient, cfg.ResyncPeriod)
	cfg.SecretInformer = cfg.KubeInformerFactory.InformerFor(&core.Secret{}, func(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
		return coreinformers.NewSecretInformer(
			client,
			cfg.WatchNamespace,
			resyncPeriod,
			cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		)
	})
	cfg.SecretLister = corelisters.NewSecretLister(cfg.SecretInformer.GetIndexer())
	// Create event recorder
	cfg.Recorder = eventer.NewEventRecorder(cfg.KubeClient, "VaultServer operator")

	// Initialize StatefulSet watcher
	sts.NewController(&cfg.Config, cfg.KubeClient, cfg.VSClient, cfg.DynamicClient).InitStsWatcher()

	return nil
}
