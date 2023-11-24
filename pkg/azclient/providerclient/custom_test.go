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
package providerclient

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func init() {
	additionalTestCases = func() {
		When("ListProviders", func() {
			It("should return list of providers", func() {
				providers, err := realClient.ListProviders(context.Background())
				Expect(err).NotTo(HaveOccurred())
				Expect(providers).NotTo(BeEmpty())
			})
		})

		When("GetVirtualMachineSupportedZones", func() {
			It("should return list of providers", func() {
				zones, err := realClient.GetVirtualMachineSupportedZones(context.Background())
				Expect(err).NotTo(HaveOccurred())
				Expect(zones).NotTo(BeEmpty())
			})
		})
	}

	beforeAllFunc = func(ctx context.Context) {
	}
	afterAllFunc = func(ctx context.Context) {
	}
}