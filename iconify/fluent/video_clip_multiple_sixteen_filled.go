package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoClipMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.25 2A2.25 2.25 0 0 0 2 4.25v5.5A2.25 2.25 0 0 0 4.25 12h6.5A2.25 2.25 0 0 0 13 9.75v-5.5A2.25 2.25 0 0 0 10.75 2h-6.5ZM6 8.996V5.004a.5.5 0 0 1 .778-.416l2.997 1.996a.5.5 0 0 1 0 .833L6.777 9.413A.5.5 0 0 1 6 8.996ZM6 14a2.496 2.496 0 0 1-2-1h7.25A2.75 2.75 0 0 0 14 10.25V4c.607.456 1 1.182 1 2v4.25A3.75 3.75 0 0 1 11.25 14H6Z"/>`),
		g.Group(children),
	)
}