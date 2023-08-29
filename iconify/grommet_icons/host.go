package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Host(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M17 4h1v1h-1V4ZM3 1h18v22H3V1Zm0 12h18H3Zm0 5h18H3ZM3 8h18H3Z"/>`),
		g.Group(children),
	)
}