package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AugmentedReality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="m6.808 55.158l15.261-21.746l8.591 8.59l15.982-15.981l19.077 29.854H6.281z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m6.808 55.158l15.261-21.746l8.591 8.59l15.982-15.981l19.077 29.854H6.281z"/><path d="M38 17.021h22.021A7.979 7.979 0 0 1 68 25v24.235H4V25a7.979 7.979 0 0 1 7.979-7.979h21.938L36 21.625l2-4.604z"/><path d="M8.208 24.125v-3H11m4.521 0h3v2.792m0 4.521v3h-2.792m-4.521 0h-3v-2.792"/><circle cx="13.229" cy="26.146" r="2.021"/><path d="M62.837 26.021H40.642m27.217 0h-5.022m5.022 3.217h-5.022m5.022 3.218h-5.022m5.022 3.217h-5.022m5.022 3.218h-5.022m5.022 3.217h-5.022m5.022 3.217h-5.022m-16.266 3.91V20.06"/></g>`),
		g.Group(children),
	)
}