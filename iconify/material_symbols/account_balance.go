package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBalance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17v-7h2v7H5Zm6 0v-7h2v7h-2Zm-9 4v-2h20v2H2Zm15-4v-7h2v7h-2ZM2 8V6l10-5l10 5v2H2Z"/>`),
		g.Group(children),
	)
}