package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9h2V7h-2v2Zm0 4h2v-2h-2v2Zm0 4h2v-2h-2v2ZM1 21V11l7-5l7 5v10h-5v-6H6v6H1Zm16 0V10l-7-5.05V3h13v18h-6Z"/>`),
		g.Group(children),
	)
}