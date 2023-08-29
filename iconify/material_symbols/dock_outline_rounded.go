package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q0-.425.288-.713T9 21h6q.425 0 .713.288T16 22q0 .425-.288.713T15 23H9q-.425 0-.713-.288T8 22Zm0-3q-.825 0-1.413-.588T6 17V3q0-.825.588-1.413T8 1h8q.825 0 1.413.588T18 3v14q0 .825-.588 1.413T16 19H8Zm0-3v1h8v-1H8Zm0-2h8V6H8v8ZM8 4h8V3H8v1Zm0 0V3v1Zm0 12v1v-1Z"/>`),
		g.Group(children),
	)
}