package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetTopBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15.5a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5V15H2a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-1v.5a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5V15H5v.5M3 12v1h2v-1H3m3 0v1h2v-1H6m14.5-.5a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}