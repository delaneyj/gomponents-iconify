package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apartment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V7h4V3h10v8h4v10h-8v-4h-2v4H3Zm2-2h2v-2H5v2Zm0-4h2v-2H5v2Zm0-4h2V9H5v2Zm4 4h2v-2H9v2Zm0-4h2V9H9v2Zm0-4h2V5H9v2Zm4 8h2v-2h-2v2Zm0-4h2V9h-2v2Zm0-4h2V5h-2v2Zm4 12h2v-2h-2v2Zm0-4h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}