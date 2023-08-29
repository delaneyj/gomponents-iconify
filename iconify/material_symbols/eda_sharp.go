package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdaSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23v-7h5.45l3.4 4.525l6.175-6.175l2.35-1.425l2.65 1.975l-8.1 8.1H3Zm4-11V2h2v10h2V1h2v11h2V3h2v9.575l-4.95 4.9L9.45 14H3V4h2v8h2Z"/>`),
		g.Group(children),
	)
}