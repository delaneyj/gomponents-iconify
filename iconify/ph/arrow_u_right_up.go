package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M221.66 85.66a8 8 0 0 1-11.32 0L176 51.31V168a64 64 0 0 1-128 0V80a8 8 0 0 1 16 0v88a48 48 0 0 0 96 0V51.31l-34.34 34.35a8 8 0 0 1-11.32-11.32l48-48a8 8 0 0 1 11.32 0l48 48a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}