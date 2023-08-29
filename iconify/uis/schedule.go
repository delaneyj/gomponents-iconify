package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schedule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4h-1V3c0-.6-.4-1-1-1s-1 .4-1 1v1H8V3c0-.6-.4-1-1-1s-1 .4-1 1v1H5C3.3 4 2 5.3 2 7v1h20V7c0-1.7-1.3-3-3-3zM2 19c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3v-9H2v9zm15-7c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm-5-4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm-5-4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1zm0 4c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1z"/>`),
		g.Group(children),
	)
}