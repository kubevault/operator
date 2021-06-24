/*
Copyright AppsCode Inc. and Contributors

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

package apis

import (
	core "k8s.io/api/core/v1"
)

const (
	Finalizer = "kubevault.com"

	// required fields:
	// - Secret.Data["token"] - a vault token
	SecretTypeTokenAuth core.SecretType = "kubevault.com/token"

	// required for SecretTypeTokenAut
	TokenAuthTokenKey = "token"

	// required fields:
	// - Secret.Data["access_key_id"] - aws access key id
	// - Secret.Data["secret_access_key"] - aws access secret key
	SecretTypeAWSAuth core.SecretType = "kubevault.com/aws"

	// required for SecretTypeAWSAuth
	AWSAuthAccessKeyIDKey = "access_key_id"
	// required for SecretTypeAWSAuth
	AWSAuthAccessSecretKey = "secret_access_key"
	// optional for SecretTypeAWSAuth
	AWSAuthSecurityTokenKey = "security_token"

	// required fields:
	// - Secret.Data["sa.json"] - gcp access secret key
	SecretTypeGCPAuth core.SecretType = "kubevault.com/gcp"
	// required for SecretTypeGCPAuth
	GCPAuthSACredentialJson = "sa.json"

	// - Secret.Data["msiToken"] - azure managed service identity (MSI)  jwt token
	SecretTypeAzureAuth = "kubevault.com/azure"

	// required for SecretTypeAzureAuth
	AzureMSIToken = "msiToken"
)

const (
	// moved from operator/pkg/controller/vault.go
	TLSCACertKey = "ca.crt"
)

const (
	VaultAuthK8sRole    = "role"
	VaultAuthApprole    = "role"
	VaultAuthLDAPGroups = "groups"
	VaultAuthLDAPUsers  = "users"
	VaultAuthJWTRole    = "role"
)

// Resource kind related constants
const (
	ResourceKindStatefulSet = "StatefulSet"
)

// List of possible condition types for a KubeVault object
const (
	// used for Vaultserver that have started provisioning
	VaultserverProvisioningStarted = "ProvisioningStarted"
	// used for Vaultserver which completed provisioning
	VaultserverProvisioned = "Provisioned"
	// used for Vaultserver that are currently being initialized using stash
	VaultserverDataRestoreStarted = "DataRestoreStarted"
	// used for Vaultserver that have been initialized using stash
	VaultserverDataRestored = "DataRestored"
	// used for Vaultserver whose pods are ready
	VaultserverReplicaReady = "ReplicaReady"
	// used for Vaultserver that are currently accepting connection
	VaultserverAcceptingConnection = "AcceptingConnection"
	// used for Vaultserver that report status OK (also implies that we can connect to it)
	VaultserverReady = "Ready"
	// used for Vaultserver that are paused
	VaultserverPaused = "Paused"
	// used for Vaultserver that are halted
	VaultserverHalted = "Halted"

	// Condition reasons
	DataRestoreStartedByExternalInitializer    = "DataRestoreStartedByExternalInitializer"
	VaultserverSuccessfullyRestored            = "SuccessfullyDataRestored"
	FailedToRestoreData                        = "FailedToRestoreData"
	AllReplicasAreReady                        = "AllReplicasReady"
	SomeReplicasAreNotReady                    = "SomeReplicasNotReady"
	VaultserverAcceptingConnectionRequest      = "VaultserverAcceptingConnectionRequest"
	VaultserverNotAcceptingConnectionRequest   = "VaultserverNotAcceptingConnectionRequest"
	ReadinessCheckSucceeded                    = "ReadinessCheckSucceeded"
	ReadinessCheckFailed                       = "ReadinessCheckFailed"
	VaultserverProvisioningStartedSuccessfully = "VaultserverProvisioningStartedSuccessfully"
	VaultserverSuccessfullyProvisioned         = "VaultserverSuccessfullyProvisioned"
	VaultserverHaltedSuccessfully              = "VaultserverHaltedSuccessfully"
)
