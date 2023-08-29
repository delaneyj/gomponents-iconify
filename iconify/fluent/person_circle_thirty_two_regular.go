package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCircleThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.373 4 4 9.373 4 16s5.373 12 12 12s12-5.373 12-12S22.627 4 16 4ZM2 16C2 8.268 8.268 2 16 2s14 6.268 14 14s-6.268 14-14 14S2 23.732 2 16Zm14 8.5c3.866 0 7-2.429 7-6.071A2.429 2.429 0 0 0 20.571 16H11.43A2.429 2.429 0 0 0 9 18.429c0 3.642 3.134 6.071 7 6.071Zm0-10A3.75 3.75 0 1 0 16 7a3.75 3.75 0 0 0 0 7.5Z"/>`),
		g.Group(children),
	)
}