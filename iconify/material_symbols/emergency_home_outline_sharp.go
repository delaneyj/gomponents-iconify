package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyHomeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6.625h2v6.75h-2v-6.75ZM12 16q-.425 0-.713-.288T11 15q0-.425.288-.713T12 14q.425 0 .713.288T13 15q0 .425-.288.713T12 16Zm0 6.8L1.2 12L12 1.2L22.8 12L12 22.8Zm0-2.8l8-8l-8-8l-8 8l8 8Zm0-8Z"/>`),
		g.Group(children),
	)
}