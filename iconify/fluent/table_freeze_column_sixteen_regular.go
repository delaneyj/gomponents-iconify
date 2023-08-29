package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFreezeColumnSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 14a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 11.5 2h-7A2.5 2.5 0 0 0 2 4.5v7A2.5 2.5 0 0 0 4.5 14h7ZM6 3h4v2H6V3Zm0 3h4v4H6V6Zm0 5h4v2H6v-2Zm-1-1H3V6h2v4Zm0-5H3v-.5A1.5 1.5 0 0 1 4.5 3H5v2Zm-2 6.5V11h2v2h-.5A1.5 1.5 0 0 1 3 11.5ZM11.5 3A1.5 1.5 0 0 1 13 4.5v7a1.5 1.5 0 0 1-1.5 1.5H11V3h.5Z"/>`),
		g.Group(children),
	)
}