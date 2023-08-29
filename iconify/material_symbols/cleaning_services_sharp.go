package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleaningServicesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V11h6V1h6v10h6v12H3Zm2-2h2v-4h2v4h2v-4h2v4h2v-4h2v4h2v-8H5v8Z"/>`),
		g.Group(children),
	)
}