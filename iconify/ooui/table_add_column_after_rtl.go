package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableAddColumnAfterRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 3h-6V2H0v16h12v-1h8V3zm-8 8H7v3H5v-3H2V9h3V6h2v3h3zm4 4h-2V5h2zm4 0h-2V5h2z"/>`),
		g.Group(children),
	)
}