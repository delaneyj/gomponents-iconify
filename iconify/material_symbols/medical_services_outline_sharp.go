package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalServicesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V6h6V2h8v4h6v16H2Zm2-2h16V8H4v12Zm6-14h4V4h-4v2ZM4 20V8v12Zm7-5v3h2v-3h3v-2h-3v-3h-2v3H8v2h3Z"/>`),
		g.Group(children),
	)
}