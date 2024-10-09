//go:build enterprise

/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

package measurements

// Regenerate the measurements by running go generate.
// The directive can be found in the file measurements.go since it does not have
// a build tag.
// The enterprise build tag is required to validate the measurements using production
// sigstore certificates.
//
// To add measurements for a new variant, add a new entry as `<csp>_<variant> = M{}` and run the generate tool.
// Entries defined as `<csp>_<variant> M` are ignored.

// revive:disable:var-naming
var (
	aws_AWSNitroTPM          = M{0: {Expected: []byte{0x73, 0x7f, 0x76, 0x7a, 0x12, 0xf5, 0x4e, 0x70, 0xee, 0xcb, 0xc8, 0x68, 0x40, 0x11, 0x32, 0x3a, 0xe2, 0xfe, 0x2d, 0xd9, 0xf9, 0x07, 0x85, 0x57, 0x79, 0x69, 0xd7, 0xa2, 0x01, 0x3e, 0x8c, 0x12}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0x0b, 0x02, 0x5a, 0xaf, 0x3a, 0x99, 0xd5, 0x27, 0x63, 0x18, 0x7c, 0x35, 0x52, 0x70, 0xf2, 0xd4, 0x06, 0x1c, 0x4e, 0xf0, 0x80, 0x70, 0xbb, 0x8e, 0x9e, 0xa9, 0x37, 0x44, 0xb7, 0xec, 0x66, 0x62}, ValidationOpt: Enforce}, 6: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x6d, 0x18, 0x5b, 0x6b, 0x1c, 0x4c, 0x96, 0x0b, 0x94, 0x3c, 0x37, 0x18, 0x41, 0x3b, 0x48, 0x9c, 0x7f, 0x9e, 0x5d, 0x10, 0x90, 0xa3, 0xdd, 0xf3, 0xbf, 0xe3, 0x2f, 0xeb, 0xc4, 0x89, 0x80, 0x68}, ValidationOpt: Enforce}, 11: {Expected: []byte{0xa0, 0x1e, 0x46, 0xad, 0x67, 0x3f, 0x3d, 0x46, 0xb6, 0x4a, 0xaa, 0xa8, 0xd4, 0xaf, 0x29, 0x38, 0x69, 0x91, 0xf5, 0xcd, 0xb2, 0x3c, 0xa2, 0x21, 0xfd, 0xe2, 0xe4, 0x9b, 0xbe, 0x0e, 0xe0, 0xda}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	aws_AWSSEVSNP            = M{0: {Expected: []byte{0x7b, 0x06, 0x8c, 0x0c, 0x3a, 0xc2, 0x9a, 0xfe, 0x26, 0x41, 0x34, 0x53, 0x6b, 0x9b, 0xe2, 0x6f, 0x1d, 0x4c, 0xcd, 0x57, 0x5b, 0x88, 0xd3, 0xc3, 0xce, 0xab, 0xf3, 0x6a, 0xc9, 0x9c, 0x02, 0x78}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0x47, 0x90, 0x1f, 0xbd, 0x0a, 0xa6, 0x83, 0xc2, 0x3c, 0x8c, 0xab, 0x2b, 0xa0, 0xf5, 0x96, 0xf4, 0x92, 0x00, 0x3e, 0xa3, 0x53, 0xa4, 0x47, 0x2a, 0x39, 0xd6, 0x62, 0x59, 0x94, 0xe1, 0xbb, 0x6d}, ValidationOpt: Enforce}, 6: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0xdc, 0x31, 0x02, 0xb2, 0x73, 0x50, 0x65, 0x71, 0x3e, 0x86, 0x08, 0x47, 0xe9, 0x34, 0x64, 0xea, 0x6b, 0xf6, 0xc5, 0x4c, 0x81, 0xe0, 0x25, 0x80, 0x2d, 0xfe, 0xf4, 0x8e, 0x35, 0xc1, 0xee, 0x6b}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x9e, 0x11, 0xf5, 0x8a, 0x7c, 0xb6, 0x97, 0x6f, 0xdf, 0x68, 0x2d, 0x91, 0x00, 0x40, 0x43, 0x58, 0x75, 0x64, 0x37, 0x6e, 0x84, 0x25, 0x0c, 0xea, 0x98, 0xe6, 0xca, 0xee, 0xbe, 0xb4, 0xad, 0xd1}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	azure_AzureSEVSNP        = M{1: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0x81, 0x0a, 0x42, 0x7c, 0x0a, 0x5e, 0x7c, 0x4d, 0xef, 0x7d, 0xb7, 0x5f, 0x91, 0x8b, 0xea, 0x07, 0x2f, 0x0a, 0xa4, 0x61, 0x2a, 0x43, 0x69, 0x2c, 0xcf, 0xa1, 0x38, 0xa3, 0xa2, 0x18, 0xd6, 0x0f}, ValidationOpt: Enforce}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x36, 0x4e, 0xf0, 0x1b, 0x1a, 0x9b, 0x04, 0x94, 0x91, 0xf5, 0xec, 0xf1, 0xb6, 0xf0, 0x59, 0xa1, 0x9d, 0x4a, 0xe4, 0x71, 0xfd, 0x71, 0x94, 0xfd, 0x4a, 0x3b, 0x72, 0xd6, 0x3e, 0x54, 0xe5, 0xef}, ValidationOpt: Enforce}, 11: {Expected: []byte{0xe5, 0x84, 0x96, 0xb3, 0x33, 0x1f, 0x9d, 0x47, 0x8d, 0x26, 0x76, 0x29, 0x26, 0xf7, 0x3e, 0xaa, 0x5c, 0x9c, 0x77, 0x3c, 0x41, 0xc6, 0x41, 0x39, 0x60, 0xb8, 0xed, 0x57, 0x0d, 0x6e, 0x00, 0xd8}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	azure_AzureTDX           = M{1: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0x1a, 0x58, 0xa6, 0xf5, 0x00, 0x1d, 0xa3, 0xe5, 0xec, 0xc3, 0xa8, 0x26, 0xd9, 0x4d, 0xc5, 0x09, 0xe9, 0x88, 0x65, 0xd8, 0x9c, 0x75, 0x9e, 0xc2, 0xd3, 0x6d, 0x95, 0x34, 0x05, 0xe1, 0x00, 0xa1}, ValidationOpt: Enforce}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x53, 0x75, 0x48, 0x58, 0x81, 0x19, 0xaa, 0xfd, 0x3d, 0x16, 0xba, 0x4b, 0xd0, 0x70, 0xde, 0x60, 0x72, 0xe9, 0x30, 0xf1, 0xd0, 0x23, 0xe0, 0x79, 0xe1, 0x2c, 0xea, 0x51, 0x4a, 0x32, 0x77, 0x65}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x8c, 0x58, 0xef, 0xa1, 0xae, 0xf2, 0x96, 0x1f, 0x30, 0x6c, 0x4c, 0x0d, 0x9b, 0x1b, 0xbb, 0xb3, 0x97, 0xd7, 0xfa, 0xff, 0x4a, 0xb8, 0x69, 0x14, 0xe3, 0x99, 0x46, 0x5c, 0xaa, 0x0f, 0xa1, 0x04}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	azure_AzureTrustedLaunch M
	gcp_GCPSEVES             = M{1: {Expected: []byte{0x36, 0x95, 0xdc, 0xc5, 0x5e, 0x3a, 0xa3, 0x40, 0x27, 0xc2, 0x77, 0x93, 0xc8, 0x5c, 0x72, 0x3c, 0x69, 0x7d, 0x70, 0x8c, 0x42, 0xd1, 0xf7, 0x3b, 0xd6, 0xfa, 0x4f, 0x26, 0x60, 0x8a, 0x5b, 0x24}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0xe1, 0x54, 0x1d, 0xa3, 0x45, 0x64, 0x92, 0x1c, 0xcd, 0x42, 0x45, 0x81, 0x64, 0xa4, 0x03, 0x43, 0x34, 0x04, 0x6f, 0x7d, 0x7f, 0x31, 0x39, 0xc8, 0xca, 0xdd, 0x16, 0xc6, 0x17, 0x31, 0xaa, 0x57}, ValidationOpt: Enforce}, 6: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x7c, 0x8f, 0x0c, 0x9d, 0x69, 0x3e, 0x42, 0xaa, 0x08, 0x08, 0xbf, 0x0b, 0x2f, 0xfc, 0x1f, 0x0b, 0xd7, 0x83, 0x4e, 0x96, 0xb4, 0xcd, 0xce, 0x65, 0x57, 0x9d, 0x7e, 0x44, 0x01, 0x39, 0x50, 0xde}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x39, 0xa8, 0x14, 0x4a, 0xc1, 0xfc, 0x83, 0x8e, 0xd6, 0xa9, 0xa5, 0x64, 0xb9, 0xf5, 0xb9, 0x78, 0x8c, 0x1f, 0x98, 0x64, 0x40, 0x20, 0x8e, 0x99, 0x5a, 0x4a, 0x28, 0x6c, 0xed, 0xac, 0x6e, 0x02}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	gcp_GCPSEVSNP            = M{1: {Expected: []byte{0x36, 0x95, 0xdc, 0xc5, 0x5e, 0x3a, 0xa3, 0x40, 0x27, 0xc2, 0x77, 0x93, 0xc8, 0x5c, 0x72, 0x3c, 0x69, 0x7d, 0x70, 0x8c, 0x42, 0xd1, 0xf7, 0x3b, 0xd6, 0xfa, 0x4f, 0x26, 0x60, 0x8a, 0x5b, 0x24}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0xc7, 0xb6, 0x0d, 0x48, 0x44, 0xd7, 0x2f, 0x3e, 0x9f, 0x00, 0xd1, 0x01, 0x61, 0x26, 0x95, 0x60, 0xe2, 0x73, 0xbb, 0x8d, 0x2e, 0x76, 0xa2, 0xb6, 0x9c, 0x31, 0xb6, 0xc5, 0x11, 0xd7, 0x57, 0x38}, ValidationOpt: Enforce}, 6: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0xd8, 0x24, 0x14, 0x85, 0x39, 0x16, 0x20, 0x14, 0x22, 0x50, 0x4a, 0x5f, 0x96, 0x29, 0xfd, 0x36, 0xa9, 0xd7, 0xb1, 0x81, 0xa3, 0x0c, 0x23, 0xd2, 0x7d, 0x68, 0xeb, 0xa4, 0xc9, 0xc2, 0x9c, 0x16}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x83, 0x58, 0x80, 0x54, 0xda, 0x69, 0xc2, 0xa2, 0xf9, 0x69, 0x57, 0xcb, 0xd3, 0x61, 0x1a, 0x9c, 0x93, 0x9c, 0x8d, 0xc1, 0x38, 0x6b, 0x26, 0x03, 0xc5, 0x9b, 0x33, 0xd8, 0x5d, 0xed, 0x19, 0x95}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	openstack_QEMUVTPM       = M{4: {Expected: []byte{0x0c, 0xc7, 0x30, 0xbc, 0x93, 0xf2, 0x0c, 0xd2, 0x6b, 0x71, 0xb8, 0x88, 0x5c, 0x79, 0xa7, 0xdb, 0x6e, 0x12, 0xcf, 0x97, 0xdf, 0x0c, 0x0e, 0xe3, 0x19, 0x2b, 0xb7, 0x75, 0xce, 0xea, 0xb9, 0x36}, ValidationOpt: Enforce}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x6d, 0xa0, 0x09, 0x7a, 0x5d, 0xdc, 0x1a, 0x94, 0xb8, 0x06, 0xa9, 0x68, 0x7b, 0xf6, 0xeb, 0xbd, 0x5e, 0x9a, 0x9d, 0xfd, 0xc6, 0x3c, 0x08, 0xe4, 0x7b, 0x0b, 0x92, 0x25, 0x46, 0xe5, 0x9e, 0x25}, ValidationOpt: Enforce}, 11: {Expected: []byte{0xba, 0xe0, 0xb9, 0x2f, 0xb1, 0xdd, 0x7e, 0xef, 0x45, 0xe2, 0xc1, 0xca, 0x4c, 0xc3, 0x54, 0x1f, 0xd4, 0x03, 0x3d, 0xc5, 0x80, 0xc2, 0x49, 0x09, 0xa0, 0xea, 0x00, 0xc4, 0xb2, 0x01, 0x25, 0xf8}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	qemu_QEMUTDX             M
	qemu_QEMUVTPM            = M{4: {Expected: []byte{0xc4, 0x42, 0x73, 0xe6, 0x07, 0x64, 0x77, 0x8f, 0x5c, 0x28, 0x40, 0x10, 0x05, 0x7c, 0xd8, 0xe3, 0xb3, 0x7c, 0x9e, 0xa9, 0x59, 0x96, 0x7e, 0xe7, 0xa5, 0x0c, 0xd6, 0x3d, 0xc0, 0xd9, 0xe9, 0x35}, ValidationOpt: Enforce}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x5e, 0xe4, 0xd7, 0x0d, 0x17, 0x5d, 0xdb, 0x58, 0x15, 0xf3, 0x4a, 0xad, 0x08, 0xb8, 0x13, 0x25, 0x1c, 0xa6, 0x77, 0x93, 0xd6, 0x93, 0xa9, 0xce, 0x5d, 0x25, 0x71, 0x33, 0xdf, 0x71, 0x25, 0xfc}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x90, 0x7f, 0x15, 0x26, 0x05, 0x18, 0xe3, 0xba, 0xd6, 0xce, 0xcd, 0x9b, 0xd3, 0xc4, 0xd0, 0x60, 0xef, 0x83, 0xe5, 0x1d, 0xff, 0x34, 0xf6, 0x53, 0xb3, 0xa8, 0x39, 0x26, 0x25, 0x3c, 0x4e, 0xf0}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
)
