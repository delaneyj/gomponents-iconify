package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.325 18.475L20 17.15V5H7.85l-2-2H20q.825 0 1.413.588T22 5v12q0 .45-.163.813t-.512.662Zm-18.15-15.3L5 5H4v12h10.15L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4l-4.3-4.3H16v2H8v-2H4q-.825 0-1.413-.588T2 17V5q0-.925.588-1.375l.587-.45ZM9.1 11.95Zm4.875-.825Z"/>`),
		g.Group(children),
	)
}