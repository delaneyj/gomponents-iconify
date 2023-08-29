package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridFourXFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22v-3H2v-2h3v-4H2v-2h3V7H2V5h3V2h2v3h4V2h2v3h4V2h2v3h3v2h-3v4h3v2h-3v4h3v2h-3v3h-2v-3h-4v3h-2v-3H7v3H5Zm2-5h4v-4H7v4Zm6 0h4v-4h-4v4Zm-6-6h4V7H7v4Zm6 0h4V7h-4v4Z"/>`),
		g.Group(children),
	)
}