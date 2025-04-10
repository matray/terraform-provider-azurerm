package virtualnetworkgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteFailoverRedundantRoute struct {
	PeeringLocations *[]string `json:"peeringLocations,omitempty"`
	Routes           *[]string `json:"routes,omitempty"`
}
