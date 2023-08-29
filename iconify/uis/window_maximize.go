package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowMaximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v5h20V3c0-.6-.4-1-1-1zM2 21c0 .6.4 1 1 1h18c.6 0 1-.4 1-1V10H2v11z"/>`),
		g.Group(children),
	)
}