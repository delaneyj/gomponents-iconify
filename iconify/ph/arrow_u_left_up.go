package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowULeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 80v88a64 64 0 0 1-128 0V51.31L45.66 85.66a8 8 0 0 1-11.32-11.32l48-48a8 8 0 0 1 11.32 0l48 48a8 8 0 0 1-11.32 11.32L96 51.31V168a48 48 0 0 0 96 0V80a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}