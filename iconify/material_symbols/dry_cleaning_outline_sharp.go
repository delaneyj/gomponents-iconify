package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DryCleaningOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22v-6H3v-4l8-3.55V7.8q-.875-.3-1.438-1.063T9 5q0-1.275.863-2.138T12 2q1.25 0 2.125.875T15 5h-2q0-.425-.288-.713T12 4q-.425 0-.713.288T11 5q0 .425.288.713T12 6h1v2.45L21 12v4h-4v6H7Zm2-7h6h-6Zm-4-1h2v-1h10v1h2v-.7l-7-3.1l-7 3.1v.7Zm4 6h6v-5H9v5Z"/>`),
		g.Group(children),
	)
}