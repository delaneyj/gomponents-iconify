package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 8a8 8 0 1 0 0 16h8a8 8 0 1 0 0-16h-8ZM2 16C2 10.477 6.477 6 12 6h8c5.523 0 10 4.477 10 10s-4.477 10-10 10h-8C6.477 26 2 21.523 2 16Z"/>`),
		g.Group(children),
	)
}