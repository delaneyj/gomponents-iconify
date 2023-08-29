package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableAddColumnAfterLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 3v14h8v1h12V2H8v1zm10 6h3V6h2v3h3v2h-3v3h-2v-3h-3zM6 5h2v10H6zM2 5h2v10H2z"/>`),
		g.Group(children),
	)
}