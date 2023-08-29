package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CorporateFare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V3h10v4h10v14H2Zm2-2h6v-2H4v2Zm0-4h6v-2H4v2Zm0-4h6V9H4v2Zm0-4h6V5H4v2Zm8 12h8V9h-8v10Zm2-6v-2h4v2h-4Zm0 4v-2h4v2h-4Z"/>`),
		g.Group(children),
	)
}