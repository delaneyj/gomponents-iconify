package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditcardAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v9h-2v-1H3v8h11v2H1V3Zm2 6h18V5H3v4Zm2 5h5v2H5v-2Zm16 0v3h3v2h-3v3h-2v-3h-3v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}