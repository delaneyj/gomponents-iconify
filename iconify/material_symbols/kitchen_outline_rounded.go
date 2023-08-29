package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 8q-.425 0-.713-.288T8 7V6q0-.425.288-.713T9 5q.425 0 .713.288T10 6v1q0 .425-.288.713T9 8Zm0 9q-.425 0-.713-.288T8 16v-3q0-.425.288-.713T9 12q.425 0 .713.288T10 13v3q0 .425-.288.713T9 17Zm-3 5q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Zm0-2h12v-9H6v9ZM6 9h12V4H6v5Z"/>`),
		g.Group(children),
	)
}