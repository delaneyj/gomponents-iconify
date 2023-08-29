package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanMountainBikingDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#6a462f"><circle cx="45.532" cy="8.013" r="2.999"/><path d="m45.441 14.87l3.382-.664l1.65 1.954l3.762 12.525l-3.176 2.654l-3.084 1.361l-6.614 8.887l-5.33 4.784l-1.444-2.605l3.246-6l-.966-7.532l1.226-2.898l8.693-2.49l-1.22-6.804"/></g><path fill="#d0cfce" stroke="#d0cfce" stroke-linejoin="round" stroke-width="2.32" d="M60.237 63.782c-2.75 0-6.815-.153-6.815-.153l-15.634-2.207l-14.996-3.445L5.715 49.2l-.024-5.02l1.723-.706l2.333 1.625l4.633 4.158l4.87 1.45l9.343 4.773l5.182-1.196l7.351 1.414l10.16 5.201l5.22-.126"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m51.525 60.693l4.98-.121L60.173 64M22.97 52.338l6.204 2.815l4.795-1.057l5.266 1.021M5.609 43.947l1.786-.721l2.95 2.14"/><circle cx="21.599" cy="36.334" r="9.723" stroke-miterlimit="10"/><circle cx="52.485" cy="48.247" r="9.723" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m21.687 36.368l14.305-16.691l2.944 2.11"/><circle cx="45.392" cy="8.152" r="3" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m38.458 20.558l2.512-.547a13.299 13.299 0 0 0 3.797-1.244s.429-2.64.623-3.502a.801.801 0 0 1 .423-.68a3.121 3.121 0 0 1 3.01-.38c1.302.552 1.922 2.39 2.16 2.99a63.005 63.005 0 0 1 3.127 9.703a4.136 4.136 0 0 1-1.33 3.166l-.158.143a5.875 5.875 0 0 1-3.35 1.257l-5.715.239a1.935 1.935 0 0 0-1.874 1.932l.106 6.584a2.285 2.285 0 0 1-1.586 2.248a1.73 1.73 0 0 1-1.947-1.573L37.17 32.09a8.651 8.651 0 0 1 .371-3.332a3.666 3.666 0 0 1 2.442-1.814l5.087-.993c1.007-.223 2.492-1.707 2.316-2.863"/><path stroke-miterlimit="10" d="m40.203 42.467l-1.957 2.41a2.66 2.66 0 0 1-2.788.862a1.622 1.622 0 0 1-.733-2.328l3.108-5.645m11.138-6.316l-7.162 8.981"/><path stroke-linecap="round" stroke-linejoin="round" d="M49.008 5.994s.598-.275.954.023c.419.35.698 1.653 1.192 2.219a3.362 3.362 0 0 0 1.864 1.115"/></g>`),
		g.Group(children),
	)
}