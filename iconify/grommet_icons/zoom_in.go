package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m16 16l7 7l-7-7Zm-6 2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0-3V5m-5 5h10"/>`),
		g.Group(children),
	)
}