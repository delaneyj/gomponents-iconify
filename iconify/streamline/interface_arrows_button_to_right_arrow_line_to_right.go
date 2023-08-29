package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsButtonToRightArrowLineToRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.71.5l6.14 6.15a.48.48 0 0 1 0 .7L5.71 13.5M2 .5v13"/>`),
		g.Group(children),
	)
}