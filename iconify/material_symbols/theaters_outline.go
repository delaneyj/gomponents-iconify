package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheatersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V3h2v2h2V3h8v2h2V3h2v18h-2v-2h-2v2H8v-2H6v2H4Zm2-4h2v-2H6v2Zm0-4h2v-2H6v2Zm0-4h2V7H6v2Zm10 8h2v-2h-2v2Zm0-4h2v-2h-2v2Zm0-4h2V7h-2v2Zm-6 10h4V5h-4v14Zm0-14h4h-4Z"/>`),
		g.Group(children),
	)
}