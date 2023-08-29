package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 6H2V3a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v3z"/><path fill="currentColor" d="M2 8h9v6H2zm0 8h9v6H3a1 1 0 0 1-1-1v-5zm11-8h9v6h-9zm8 14h-8v-6h9v5a1 1 0 0 1-1 1z" opacity=".5"/><path fill="currentColor" d="M22 8V6H2v2h9v6H2v2h9v6h2v-6h9v-2h-9V8z" opacity=".25"/>`),
		g.Group(children),
	)
}