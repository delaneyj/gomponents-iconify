package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polygon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 20.794h-9a1 1 0 0 1-.866-.5l-4.5-7.794a1.002 1.002 0 0 1 0-1l4.5-7.794a1 1 0 0 1 .866-.5h9a1 1 0 0 1 .866.5l4.5 7.794a1.002 1.002 0 0 1 0 1l-4.5 7.794a1 1 0 0 1-.866.5Z"/>`),
		g.Group(children),
	)
}