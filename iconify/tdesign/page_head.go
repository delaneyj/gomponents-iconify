package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageHead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10v12h-2V12H5v10H3V10h18Zm0-8v6H3V2h18Zm-2 2H5v2h14V4Z"/>`),
		g.Group(children),
	)
}