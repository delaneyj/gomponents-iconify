package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperModeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.85 12l2.45 2.45q.3.3.288.7t-.288.7q-.3.3-.713.313t-.712-.288L4.7 12.7q-.3-.3-.3-.7t.3-.7l3.175-3.175q.3-.3.713-.287t.712.312q.275.3.288.7t-.288.7L6.85 12ZM5 17h2v1h10v-1h2v4q0 .825-.588 1.413T17 23H7q-.825 0-1.413-.588T5 21v-4ZM7 7H5V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v1Zm0 13v1h10v-1H7ZM7 4h10V3H7v1Zm10.15 8L14.7 9.55q-.3-.3-.288-.7t.288-.7q.3-.3.713-.312t.712.287L19.3 11.3q.3.3.3.7t-.3.7l-3.175 3.175q-.3.3-.713.288t-.712-.313q-.275-.3-.287-.7t.287-.7L17.15 12ZM7 4V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}