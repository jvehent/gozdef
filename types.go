// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Contributor: Julien Vehent jvehent@mozilla.com [:ulfr]
package gozdef

// An HostAssetHint describes information about a host that can be used
// to correlate asset information in MozDef. This is primarily used by MIG
type AssetHint struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	IPv4      []string `json:"ipv4"`
	IPv6      []string `json:"ipv6"`
	OS        string   `json:"os"`
	Arch      string   `json:"arch"`
	Ident     string   `json:"ident"`
	Init      string   `json:"init"`
	IsProxied bool     `json:"isproxied"`
	Operator  string   `json:"operator"`
}
