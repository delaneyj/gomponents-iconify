package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16q0 .65-.35 1.15t-.925.725L18.85 16H20V8h-9.15l-2-2H20q.825 0 1.413.588T22 8v8Zm-4.5-4q-.625 0-1.063-.438T16 10.5q0-.625.438-1.063T17.5 9q.625 0 1.063.438T19 10.5q0 .625-.438 1.063T17.5 12Zm-8.35 0Zm5.7 0ZM7 13H6q-.425 0-.712-.288T5 12q0-.425.288-.713T6 11h1v-1q0-.425.288-.713T8 9q.425 0 .713.288T9 10v1h1q.425 0 .713.288T11 12q0 .425-.288.713T10 13H9v1q0 .425-.288.713T8 15q-.425 0-.713-.288T7 14v-1Zm-3 5q-.825 0-1.413-.588T2 16V8q0-.85.6-1.438t1.45-.587h1.925L8 8H4v8h9.15L2.075 4.925q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L15.15 18H4Z"/>`),
		g.Group(children),
	)
}