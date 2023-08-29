package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFreezeColumnAndRowSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 4.5A2.5 2.5 0 0 0 11.5 2h-7A2.5 2.5 0 0 0 2 4.5V10h3V6H3V4.5A1.5 1.5 0 0 1 4.5 3h7A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5H10v-2H6v3h5.5a2.5 2.5 0 0 0 2.5-2.5v-7ZM10 6H6v4h4V6Zm-5 5H2v.5A2.5 2.5 0 0 0 4.5 14H5v-3Z"/>`),
		g.Group(children),
	)
}