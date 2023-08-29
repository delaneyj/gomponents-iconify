package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoTouch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5.5 9v15M9 9V3.5a3 3 0 0 1 3-3h.5v12v-7a3 3 0 0 1 3-3h.5V16v-3.5a3 3 0 0 1 3-3h.5v10M9.17 2.5H8.5a3 3 0 0 0-3 3m-5-5l23 23m-5.924-2.424A3.497 3.497 0 0 0 16 24m-3.5-4v-1c0-.84.295-1.61.788-2.212"/>`),
		g.Group(children),
	)
}