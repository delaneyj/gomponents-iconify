package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V8q0-.85.6-1.438t1.45-.587h1.925L15 15h-2.85L2.075 4.925q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L15.15 18H4ZM8.85 6H20q.825 0 1.413.588T22 8v8q0 .65-.35 1.15t-.925.725L8.85 6Zm8.65 3q-.625 0-1.063.438T16 10.5q0 .625.438 1.063T17.5 12q.625 0 1.063-.438T19 10.5q0-.625-.438-1.063T17.5 9ZM7 13v1q0 .425.288.713T8 15q.425 0 .713-.288T9 14v-1h1q.425 0 .713-.288T11 12q0-.425-.288-.713T10 11H9v-1q0-.425-.288-.713T8 9q-.425 0-.713.288T7 10v1H6q-.425 0-.713.288T5 12q0 .425.288.713T6 13h1Z"/>`),
		g.Group(children),
	)
}