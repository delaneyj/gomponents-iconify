package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFreezeColumnAndRowSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 2A2.5 2.5 0 0 1 14 4.5v7a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5v-7A2.5 2.5 0 0 1 4.5 2h7ZM3 6v4h2V6H3Zm3 0v4h4V6H6Zm4 7v-2H6v2h4Zm-5 0v-2H3v.5A1.5 1.5 0 0 0 4.5 13H5ZM3 4.5V5h8v8h.5a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 11.5 3h-7A1.5 1.5 0 0 0 3 4.5Z"/>`),
		g.Group(children),
	)
}