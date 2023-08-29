package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SingleBedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19H6l-.65-2H4v-7h2V5h12v5h2v7h-1.35L18 19h-1l-.65-2h-8.7L7 19Zm6-9h3V7h-3v3Zm-5 0h3V7H8v3Zm-2 5h12v-3H6v3Zm12 0H6h12Z"/>`),
		g.Group(children),
	)
}