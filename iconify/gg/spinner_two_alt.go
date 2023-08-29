package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerTwoAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22c5.523 0 10-4.477 10-10h-3a7 7 0 0 1-7 7v3ZM2 12C2 6.477 6.477 2 12 2v3a7 7 0 0 0-7 7H2Z"/>`),
		g.Group(children),
	)
}