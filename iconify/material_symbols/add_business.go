package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddBusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 23v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM2 20v-6H1v-2l1-5h15l1 5v2h-1v3h-2v-3h-4v6H2Zm2-2h5v-4H4v4ZM2 6V4h15v2H2Z"/>`),
		g.Group(children),
	)
}