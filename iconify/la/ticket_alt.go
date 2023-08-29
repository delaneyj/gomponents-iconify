package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 7v7h1c1.191 0 2 .809 2 2c0 1.191-.809 2-2 2H2v7h28v-7h-1c-1.191 0-2-.809-2-2c0-1.191.809-2 2-2h1V7zm2 2h24v3.188c-1.715.449-3 1.957-3 3.812c0 1.855 1.285 3.363 3 3.813V23H4v-3.188c1.715-.449 3-1.957 3-3.812c0-1.855-1.285-3.363-3-3.813z"/>`),
		g.Group(children),
	)
}