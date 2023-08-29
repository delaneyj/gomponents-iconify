package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M197.66 197.66a8 8 0 0 1-11.32 0L72 83.31V168a8 8 0 0 1-16 0V64a8 8 0 0 1 8-8h104a8 8 0 0 1 0 16H83.31l114.35 114.34a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}