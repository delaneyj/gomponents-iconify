package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 8h14a1 1 0 0 0 0-2H5a1 1 0 0 0 0 2Zm16 3H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2Zm-2 5H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}