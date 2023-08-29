package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224.49 88.49a12 12 0 0 1-17 0L180 61v107a68 68 0 0 1-136 0V80a12 12 0 0 1 24 0v88a44 44 0 0 0 88 0V61l-27.51 27.49a12 12 0 1 1-17-17l48-48a12 12 0 0 1 17 0l48 48a12 12 0 0 1 0 17Z"/>`),
		g.Group(children),
	)
}