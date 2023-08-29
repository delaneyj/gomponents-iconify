package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foundation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-3H3v-2h2v-3H2l10-9l10 9h-3v3h2v2h-2v3h-2v-3h-4v3h-2v-3H7v3H5Zm2-5h4V6.6l-4 3.6V15Zm6 0h4v-4.8l-4-3.6V15Z"/>`),
		g.Group(children),
	)
}