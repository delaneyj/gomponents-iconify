package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactoryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V9.975L9 7v2l5-2v3h8v12H2Zm2-2h16v-8h-8V9.95l-5 2V10l-3 1.325V20Zm7-2h2v-4h-2v4Zm-4 0h2v-4H7v4Zm8 0h2v-4h-2v4Zm7-8h-5l1-8h3l1 8ZM4 20h16H4Z"/>`),
		g.Group(children),
	)
}