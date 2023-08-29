package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoLuggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.375 6H13.5V3.5h-3.125V6ZM8 22q-.425 0-.713-.288T7 21q-.825 0-1.413-.588T5 19V8q0-.6.313-1.113t.887-.712L19 19q0 .825-.588 1.413T17 21q0 .425-.288.713T16 22q-.425 0-.713-.288T15 21H9q0 .425-.288.713T8 22Zm0-4h1.5V9.475H8V18Zm3.25 0h1.5v-5.25l-1.5-1.525V18Zm3.25 0H16v-2l-1.5-1.5V18Zm5.975 5.3L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM19 16.125l-3-3V9h-1.5v2.625l-1.75-1.75V9h-.875l-3-3V3.5q0-.625.438-1.063T10.374 2H13.5q.625 0 1.063.438T15 3.5V6h2q.825 0 1.413.588T19 8v8.125Z"/>`),
		g.Group(children),
	)
}