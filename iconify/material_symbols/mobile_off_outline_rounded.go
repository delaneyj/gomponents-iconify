package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L1.4 4.2q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM5 6.425l2 2V18h9.6l2.4 2.4v.6q0 .825-.588 1.413T17 23H7q-.825 0-1.413-.588T5 21V6.425ZM7 20v1h10v-1H7ZM19 3v13.15l-2-2V6H8.825L5.15 2.3q.25-.575.738-.937T7 1h10q.825 0 1.413.588T19 3ZM7 4h10V3H7v1Zm0 16v1v-1ZM7 4V3v1Z"/>`),
		g.Group(children),
	)
}