package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfloatLandscapeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v7h-2V6H4v12h11v2Zm9.075-5.5l1.425-1.425L9.4 10H12V8H6v6h2v-2.575ZM17 20v-7h5v7Zm-5-8Z"/>`),
		g.Group(children),
	)
}