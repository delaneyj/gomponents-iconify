package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v4h20V3c0-.6-.4-1-1-1zM2 15h9V9H2v6zm0 6c0 .6.4 1 1 1h8v-5H2v4zm11-6h9V9h-9v6zm0 7h8c.6 0 1-.4 1-1v-4h-9v5z"/>`),
		g.Group(children),
	)
}