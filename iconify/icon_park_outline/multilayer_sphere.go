package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultilayerSphere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path stroke-linecap="round" d="M4.4 20c1.853 9.129 9.924 16 19.6 16c9.676 0 17.747-6.871 19.6-16"/><path stroke-linecap="round" d="M5.664 16C8.75 23.064 15.8 28 24 28c8.201 0 15.25-4.936 18.336-12"/><path stroke-linecap="round" d="M7.999 12c3.648 4.858 9.458 8 16.001 8c6.543 0 12.353-3.142 16.015-8M11.998 8A19.91 19.91 0 0 0 24 12a19.91 19.91 0 0 0 12.002-4"/></g>`),
		g.Group(children),
	)
}