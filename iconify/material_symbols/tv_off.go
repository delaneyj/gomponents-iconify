package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.325 18.475L5.85 3H20q.825 0 1.413.588T22 5v12q0 .45-.163.813t-.512.662Zm-18.15-15.3v2.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4l-4.3-4.3H16v2H8v-2H4q-.825 0-1.413-.588T2 17V5q0-.925.588-1.375l.587-.45Z"/>`),
		g.Group(children),
	)
}