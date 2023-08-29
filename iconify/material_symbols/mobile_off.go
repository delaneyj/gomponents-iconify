package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM5 6.425l2 2V18h9.6l2.4 2.4v.6q0 .825-.588 1.413T17 23H7q-.825 0-1.413-.588T5 21V6.425ZM19 3v13.15l-2-2V6H8.825L5.15 2.3q.25-.575.738-.937T7 1h10q.825 0 1.413.588T19 3Z"/>`),
		g.Group(children),
	)
}