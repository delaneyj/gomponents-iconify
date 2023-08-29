package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Factory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V9.975L9 7v2l5-2v3h8v12H2Zm9-4h2v-4h-2v4Zm-4 0h2v-4H7v4Zm8 0h2v-4h-2v4Zm6.8-9.5h-4.625l.85-6.5H21l.8 6.5Z"/>`),
		g.Group(children),
	)
}