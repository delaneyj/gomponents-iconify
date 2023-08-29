package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func State(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 6v12h5V6Zm4 7h-3V7h3Zm5-12H6a2 2 0 0 0-2 2v18a2.005 2.005 0 0 0 2 2h13a2 2 0 0 0 2-2V3a2.006 2.006 0 0 0-2-2Zm0 20H6V3h13Z"/>`),
		g.Group(children),
	)
}