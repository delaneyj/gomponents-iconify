package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoLuggageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 16.125l-2-2V8h-6.125l-2-2V3.5q0-.625.438-1.063T10.374 2H13.5q.625 0 1.063.438T15 3.5V6h2q.825 0 1.413.588T19 8v8.125Zm-3-3l-1.5-1.5V9H16v4.125Zm-3.25-3.25L11.875 9h.875v.875ZM10.375 6H13.5V3.5h-3.125V6ZM8 22q-.425 0-.713-.288T7 21q-.825 0-1.413-.588T5 19V8q0-.6.313-1.113t.887-.712L8.025 8H7v11h10v-2l2 2q0 .825-.588 1.413T17 21q0 .425-.288.713T16 22q-.425 0-.713-.288T15 21H9q0 .425-.288.713T8 22Zm0-4V9.475h1.5V18H8Zm3.25-6.775l1.5 1.525V18h-1.5v-6.775ZM14.5 14.5L16 16v2h-1.5v-3.5Zm-.55-3.45ZM11.6 14.4Zm8.875 8.9L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425Z"/>`),
		g.Group(children),
	)
}