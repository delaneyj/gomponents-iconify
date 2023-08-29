package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 7H6c-.6 0-1 .4-1 1v2.5h7.5c.6 0 1 .4 1 1V19H16c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1zm-5 5H3c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1h8c.6 0 1-.4 1-1v-8c0-.6-.4-1-1-1zM21 2H9c-.6 0-1 .4-1 1v2.5h9.5c.6 0 1 .4 1 1V16H21c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}