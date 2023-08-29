package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintErrorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21q-.425 0-.713-.288T18 20q0-.425.288-.713T19 19q.425 0 .713.288T20 20q0 .425-.288.713T19 21Zm-1-4v-5h2v5h-2ZM6 21v-4H2V8h19.825v2H16v5H8v4h8v2H6ZM6 7V3h12v4H6Z"/>`),
		g.Group(children),
	)
}