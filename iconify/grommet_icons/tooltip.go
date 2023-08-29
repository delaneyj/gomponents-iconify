package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tooltip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M16.5 18L12 22.5L7.5 18H1V1h22v17h-6.5ZM6 10h1V9H6v1Zm5.5 0h1V9h-1v1Zm5.5 0h1V9h-1v1Z"/>`),
		g.Group(children),
	)
}