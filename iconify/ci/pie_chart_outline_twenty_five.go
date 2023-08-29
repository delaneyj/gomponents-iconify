package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartOutlineTwentyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10ZM11 4.062a8 8 0 1 0 8.915 9.1V13H11V4.062Zm2 0V11h6.938A8 8 0 0 0 13 4.062Z"/>`),
		g.Group(children),
	)
}