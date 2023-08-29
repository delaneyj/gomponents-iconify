package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddBusinessOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 23v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM2 20v-6H1v-2l1-5h15l1 5v2h-1v3h-2v-3h-4v6H2Zm2-2h5v-4H4v4Zm-.95-6h12.9h-12.9ZM2 6V4h15v2H2Zm1.05 6h12.9l-.6-3H3.65l-.6 3Z"/>`),
		g.Group(children),
	)
}