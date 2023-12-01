// /*
// Copyright The Kubernetes Authors.
//
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
// */

// Code generated by client-gen. DO NOT EDIT.
package azclient

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/accountclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobcontainerclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobservicepropertiesclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/fileshareclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/ipgroupclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/policy/ratelimit"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/providerclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/registryclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/resourcegroupclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/secretclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/sshpublickeyresourceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/vaultclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualnetworkclient"
)

type ClientFactoryImpl struct {
	*ClientFactoryConfig
	cred                                    azcore.TokenCredential
	accountclientInterface                  accountclient.Interface
	availabilitysetclientInterface          availabilitysetclient.Interface
	blobcontainerclientInterface            blobcontainerclient.Interface
	blobservicepropertiesclientInterface    blobservicepropertiesclient.Interface
	deploymentclientInterface               deploymentclient.Interface
	diskclientInterface                     diskclient.Interface
	fileshareclientInterface                fileshareclient.Interface
	interfaceclientInterface                interfaceclient.Interface
	ipgroupclientInterface                  ipgroupclient.Interface
	loadbalancerclientInterface             loadbalancerclient.Interface
	managedclusterclientInterface           managedclusterclient.Interface
	privateendpointclientInterface          privateendpointclient.Interface
	privatelinkserviceclientInterface       privatelinkserviceclient.Interface
	privatezoneclientInterface              privatezoneclient.Interface
	providerclientInterface                 providerclient.Interface
	publicipaddressclientInterface          publicipaddressclient.Interface
	publicipprefixclientInterface           publicipprefixclient.Interface
	registryclientInterface                 registryclient.Interface
	resourcegroupclientInterface            resourcegroupclient.Interface
	routetableclientInterface               routetableclient.Interface
	secretclientInterface                   secretclient.Interface
	securitygroupclientInterface            securitygroupclient.Interface
	snapshotclientInterface                 snapshotclient.Interface
	sshpublickeyresourceclientInterface     sshpublickeyresourceclient.Interface
	subnetclientInterface                   subnetclient.Interface
	vaultclientInterface                    vaultclient.Interface
	virtualmachineclientInterface           virtualmachineclient.Interface
	virtualmachinescalesetclientInterface   virtualmachinescalesetclient.Interface
	virtualmachinescalesetvmclientInterface virtualmachinescalesetvmclient.Interface
	virtualnetworkclientInterface           virtualnetworkclient.Interface
}

func NewClientFactory(config *ClientFactoryConfig, armConfig *ARMClientConfig, cred azcore.TokenCredential, clientOptionsMutFn ...func(option *arm.ClientOptions)) (ClientFactory, error) {
	if config == nil {
		config = &ClientFactoryConfig{}
	}
	if cred == nil {
		cred = &azidentity.DefaultAzureCredential{}
	}

	var options *arm.ClientOptions
	var err error

	//initialize {accountclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/accountclient Account  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	var ratelimitOption *ratelimit.Config
	var rateLimitPolicy policy.Policy

	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	accountclientInterface, err := accountclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {availabilitysetclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient AvailabilitySet  Interface availabilitySetRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("availabilitySetRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	availabilitysetclientInterface, err := availabilitysetclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {blobcontainerclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobcontainerclient Account BlobContainer Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	blobcontainerclientInterface, err := blobcontainerclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {blobservicepropertiesclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobservicepropertiesclient BlobServiceProperties  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	blobservicepropertiesclientInterface, err := blobservicepropertiesclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {deploymentclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient Deployment  Interface deploymentRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("deploymentRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	deploymentclientInterface, err := deploymentclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {diskclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient Disk  Interface diskRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("diskRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	diskclientInterface, err := diskclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {fileshareclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/fileshareclient Account FileShare Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	fileshareclientInterface, err := fileshareclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {interfaceclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient Interface  Interface interfaceRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("interfaceRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	interfaceclientInterface, err := interfaceclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {ipgroupclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/ipgroupclient IPGroup  Interface ipGroupRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("ipGroupRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	ipgroupclientInterface, err := ipgroupclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {loadbalancerclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient LoadBalancer  Interface loadBalancerRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("loadBalancerRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	loadbalancerclientInterface, err := loadbalancerclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {managedclusterclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient ManagedCluster  Interface containerServiceRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("containerServiceRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	managedclusterclientInterface, err := managedclusterclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {privateendpointclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient PrivateEndpoint  Interface privateEndpointRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("privateEndpointRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	privateendpointclientInterface, err := privateendpointclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {privatelinkserviceclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient PrivateLinkService  Interface privateLinkServiceRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("privateLinkServiceRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	privatelinkserviceclientInterface, err := privatelinkserviceclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {privatezoneclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient PrivateZone  Interface privateDNSRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("privateDNSRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	privatezoneclientInterface, err := privatezoneclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {providerclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/providerclient Provider  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	providerclientInterface, err := providerclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {publicipaddressclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient PublicIPAddress  Interface publicIPAddressRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("publicIPAddressRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	publicipaddressclientInterface, err := publicipaddressclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {publicipprefixclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient PublicIPPrefix  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	publicipprefixclientInterface, err := publicipprefixclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {registryclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/registryclient Registry  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	registryclientInterface, err := registryclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {resourcegroupclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/resourcegroupclient ResourceGroup  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	resourcegroupclientInterface, err := resourcegroupclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {routetableclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient RouteTable  Interface routeTableRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("routeTableRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	routetableclientInterface, err := routetableclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {secretclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/secretclient Vault Secret Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	secretclientInterface, err := secretclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {securitygroupclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient SecurityGroup  Interface securityGroupRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("securityGroupRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	securitygroupclientInterface, err := securitygroupclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {snapshotclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient Snapshot  Interface snapshotRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("snapshotRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	snapshotclientInterface, err := snapshotclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {sshpublickeyresourceclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/sshpublickeyresourceclient SSHPublicKeyResource  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	sshpublickeyresourceclientInterface, err := sshpublickeyresourceclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {subnetclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient VirtualNetwork Subnet Interface subnetsRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("subnetsRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	subnetclientInterface, err := subnetclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {vaultclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/vaultclient Vault  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	vaultclientInterface, err := vaultclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {virtualmachineclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient VirtualMachine  Interface virtualMachineRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("virtualMachineRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	virtualmachineclientInterface, err := virtualmachineclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {virtualmachinescalesetclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient VirtualMachineScaleSet  Interface virtualMachineSizesRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("virtualMachineSizesRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	if rateLimitPolicy != nil {
		options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	virtualmachinescalesetclientInterface, err := virtualmachinescalesetclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {virtualmachinescalesetvmclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient VirtualMachineScaleSet VirtualMachineScaleSetVM Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	virtualmachinescalesetvmclientInterface, err := virtualmachinescalesetvmclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {virtualnetworkclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualnetworkclient VirtualNetwork  Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	for _, optionMutFn := range clientOptionsMutFn {
		if optionMutFn != nil {
			optionMutFn(options)
		}
	}
	virtualnetworkclientInterface, err := virtualnetworkclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	return &ClientFactoryImpl{
		ClientFactoryConfig: config,
		cred:                cred, accountclientInterface: accountclientInterface,
		availabilitysetclientInterface:          availabilitysetclientInterface,
		blobcontainerclientInterface:            blobcontainerclientInterface,
		blobservicepropertiesclientInterface:    blobservicepropertiesclientInterface,
		deploymentclientInterface:               deploymentclientInterface,
		diskclientInterface:                     diskclientInterface,
		fileshareclientInterface:                fileshareclientInterface,
		interfaceclientInterface:                interfaceclientInterface,
		ipgroupclientInterface:                  ipgroupclientInterface,
		loadbalancerclientInterface:             loadbalancerclientInterface,
		managedclusterclientInterface:           managedclusterclientInterface,
		privateendpointclientInterface:          privateendpointclientInterface,
		privatelinkserviceclientInterface:       privatelinkserviceclientInterface,
		privatezoneclientInterface:              privatezoneclientInterface,
		providerclientInterface:                 providerclientInterface,
		publicipaddressclientInterface:          publicipaddressclientInterface,
		publicipprefixclientInterface:           publicipprefixclientInterface,
		registryclientInterface:                 registryclientInterface,
		resourcegroupclientInterface:            resourcegroupclientInterface,
		routetableclientInterface:               routetableclientInterface,
		secretclientInterface:                   secretclientInterface,
		securitygroupclientInterface:            securitygroupclientInterface,
		snapshotclientInterface:                 snapshotclientInterface,
		sshpublickeyresourceclientInterface:     sshpublickeyresourceclientInterface,
		subnetclientInterface:                   subnetclientInterface,
		vaultclientInterface:                    vaultclientInterface,
		virtualmachineclientInterface:           virtualmachineclientInterface,
		virtualmachinescalesetclientInterface:   virtualmachinescalesetclientInterface,
		virtualmachinescalesetvmclientInterface: virtualmachinescalesetvmclientInterface,
		virtualnetworkclientInterface:           virtualnetworkclientInterface,
	}, nil
}

func (factory *ClientFactoryImpl) GetAccountClient() accountclient.Interface {
	return factory.accountclientInterface
}

func (factory *ClientFactoryImpl) GetAvailabilitySetClient() availabilitysetclient.Interface {
	return factory.availabilitysetclientInterface
}

func (factory *ClientFactoryImpl) GetBlobContainerClient() blobcontainerclient.Interface {
	return factory.blobcontainerclientInterface
}

func (factory *ClientFactoryImpl) GetBlobServicePropertiesClient() blobservicepropertiesclient.Interface {
	return factory.blobservicepropertiesclientInterface
}

func (factory *ClientFactoryImpl) GetDeploymentClient() deploymentclient.Interface {
	return factory.deploymentclientInterface
}

func (factory *ClientFactoryImpl) GetDiskClient() diskclient.Interface {
	return factory.diskclientInterface
}

func (factory *ClientFactoryImpl) GetFileShareClient() fileshareclient.Interface {
	return factory.fileshareclientInterface
}

func (factory *ClientFactoryImpl) GetInterfaceClient() interfaceclient.Interface {
	return factory.interfaceclientInterface
}

func (factory *ClientFactoryImpl) GetIPGroupClient() ipgroupclient.Interface {
	return factory.ipgroupclientInterface
}

func (factory *ClientFactoryImpl) GetLoadBalancerClient() loadbalancerclient.Interface {
	return factory.loadbalancerclientInterface
}

func (factory *ClientFactoryImpl) GetManagedClusterClient() managedclusterclient.Interface {
	return factory.managedclusterclientInterface
}

func (factory *ClientFactoryImpl) GetPrivateEndpointClient() privateendpointclient.Interface {
	return factory.privateendpointclientInterface
}

func (factory *ClientFactoryImpl) GetPrivateLinkServiceClient() privatelinkserviceclient.Interface {
	return factory.privatelinkserviceclientInterface
}

func (factory *ClientFactoryImpl) GetPrivateZoneClient() privatezoneclient.Interface {
	return factory.privatezoneclientInterface
}

func (factory *ClientFactoryImpl) GetProviderClient() providerclient.Interface {
	return factory.providerclientInterface
}

func (factory *ClientFactoryImpl) GetPublicIPAddressClient() publicipaddressclient.Interface {
	return factory.publicipaddressclientInterface
}

func (factory *ClientFactoryImpl) GetPublicIPPrefixClient() publicipprefixclient.Interface {
	return factory.publicipprefixclientInterface
}

func (factory *ClientFactoryImpl) GetRegistryClient() registryclient.Interface {
	return factory.registryclientInterface
}

func (factory *ClientFactoryImpl) GetResourceGroupClient() resourcegroupclient.Interface {
	return factory.resourcegroupclientInterface
}

func (factory *ClientFactoryImpl) GetRouteTableClient() routetableclient.Interface {
	return factory.routetableclientInterface
}

func (factory *ClientFactoryImpl) GetSecretClient() secretclient.Interface {
	return factory.secretclientInterface
}

func (factory *ClientFactoryImpl) GetSecurityGroupClient() securitygroupclient.Interface {
	return factory.securitygroupclientInterface
}

func (factory *ClientFactoryImpl) GetSnapshotClient() snapshotclient.Interface {
	return factory.snapshotclientInterface
}

func (factory *ClientFactoryImpl) GetSSHPublicKeyResourceClient() sshpublickeyresourceclient.Interface {
	return factory.sshpublickeyresourceclientInterface
}

func (factory *ClientFactoryImpl) GetSubnetClient() subnetclient.Interface {
	return factory.subnetclientInterface
}

func (factory *ClientFactoryImpl) GetVaultClient() vaultclient.Interface {
	return factory.vaultclientInterface
}

func (factory *ClientFactoryImpl) GetVirtualMachineClient() virtualmachineclient.Interface {
	return factory.virtualmachineclientInterface
}

func (factory *ClientFactoryImpl) GetVirtualMachineScaleSetClient() virtualmachinescalesetclient.Interface {
	return factory.virtualmachinescalesetclientInterface
}

func (factory *ClientFactoryImpl) GetVirtualMachineScaleSetVMClient() virtualmachinescalesetvmclient.Interface {
	return factory.virtualmachinescalesetvmclientInterface
}

func (factory *ClientFactoryImpl) GetVirtualNetworkClient() virtualnetworkclient.Interface {
	return factory.virtualnetworkclientInterface
}
