package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.15l-2-2V9h-8.15L10 7.15V5H7.85l-2-2H10q.825 0 1.413.588T12 5v2h8q.825 0 1.413.588T22 9v10.15Zm-2.2 3.475L18.15 21H4q-.825 0-1.413-.588T2 19V4.8l-.6-.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.288.688t-.288.712q-.275.275-.687.288t-.713-.263ZM4 19h2v-2H4v2Zm0-4h2v-2H4v2Zm0-4h2V9H4v2Zm4 8h2v-2H8v2Zm0-4h2v-2H8v2Zm4 4h4.15l-2-2H12v2Zm6-6h-2v-2h2v2Z"/>`),
		g.Group(children),
	)
}