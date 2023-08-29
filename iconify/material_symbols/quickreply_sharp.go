package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickreplySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 23v-5h-2v-6h5l-1.7 4h2.2L19 23ZM2 22V2h20v8h-7v8H6l-4 4Z"/>`),
		g.Group(children),
	)
}