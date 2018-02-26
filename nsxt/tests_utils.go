/* Copyright © 2018 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

// Default names or prefixed of NSX backend existing objects used in the acceptance tests.
// Those defaults can be overridden using environment parameters
const tier0RouterDefaultName string = "PLR-1 LogicalRouterTier0"
const edgeClusterDefaultName string = "edgecluster1"
const switchingProfileDefaultName string = "nsx-default-mac-profile"
const vlanTransportZoneName string = "transportzone2"
const overlayTransportZoneNamePrefix string = "1-transportzone"

const singleTag string = "[{scope = \"scope1\", tag = \"tag1\"}]"
const doubleTags string = "[{scope = \"scope1\", tag = \"tag1\"}, {scope = \"scope2\", tag = \"tag2\"}]"

func getTier0RouterName() string {
	name := os.Getenv("NSX_TIER0_ROUTER")
	if name == "" {
		name = tier0RouterDefaultName
	}
	return name
}

func getEdgeClusterName() string {
	name := os.Getenv("NSX_EDGE_CLUSTER")
	if name == "" {
		name = edgeClusterDefaultName
	}
	return name
}

func getSwitchingProfileName() string {
	name := os.Getenv("NSX_SWITCHING_PROFILE")
	if name == "" {
		name = switchingProfileDefaultName
	}
	return name
}

func getVlanTransportZoneName() string {
	name := os.Getenv("NSX_VLAN_TRANSPORT_ZONE")
	if name == "" {
		name = vlanTransportZoneName
	}
	return name
}

func getOverlayTransportZoneName() string {
	name := os.Getenv("NSX_OVERLAY_TRANSPORT_ZONE")
	if name == "" {
		name = overlayTransportZoneNamePrefix
	}
	return name
}

func getTestVMID() string {
	return os.Getenv("NSX_TEST_VM_ID")

}

// copyStatePtr returns a TestCheckFunc that copies the reference to the test
// run's state to t. This allows access to the state data in later steps where
// it's not normally accessible (ie: in pre-config parts in another test step).
func copyStatePtr(t **terraform.State) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		*t = s
		return nil
	}
}
