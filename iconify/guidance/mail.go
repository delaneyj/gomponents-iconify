package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5 6.5l11.375 7h.25l11.375-7m0-2.5v16.5H23c-3-.5-8-.75-11-.75S4 20 1 20.5H.5V4c3-.5 8.5-.75 11.5-.75s8.5.25 11.5.75Z"/>`),
		g.Group(children),
	)
}