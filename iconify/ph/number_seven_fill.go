package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSevenFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm-32.47 50.69l-40 112a8 8 0 1 1-15.06-5.38L148.65 80H96a8 8 0 0 1 0-16h64a8 8 0 0 1 7.53 10.69Z"/>`),
		g.Group(children),
	)
}