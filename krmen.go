// Package krmen compiles 6-layer .jos UI atoms (L0 7-device CSS @container shells,
// L1 material backgrounds, L2 design tokens, L3 precompiled components, and L5 FX).
package krmen

// Version is the semantic version of josfox.cloud/krmen.
const Version = "0.1.0"

// DeviceContainers maps the 7 canonical KRMEN L0 device targets to native CSS container/media queries.
var DeviceContainers = map[string]string{
	"desktop":      "@container krmen-shell (min-width: 1024px)",
	"tablet":       "@container krmen-shell (min-width: 680px) and (max-width: 1023px)",
	"ios":          "@supports (-webkit-touch-callout: none)",
	"android":      "@container krmen-shell (max-width: 679px)",
	"foldable_duo": "@media (horizontal-viewport-segments: 2)",
	"kiosk":        "@media (pointer: coarse) and (min-height: 1280px)",
	"vr_headset":   "@media (dynamic-range: high) and (pointer: coarse) and (hover: hover)",
}
