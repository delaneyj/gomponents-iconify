package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 13h8v8h-8v-8Zm0-2V3h8v8h-8Zm-2 0H3V3h8v8Zm0 2v8H3v-8h8Z"/>`),
		g.Group(children),
	)
}