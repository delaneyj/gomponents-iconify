package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tickets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17.5 4.23V8m0 2v4m0 2v3.77M.5 3.5v6a2.5 2.5 0 0 1 0 5v6H1l3.04-.434a56.277 56.277 0 0 1 15.92 0L23 20.5h.5v-6a2.5 2.5 0 0 1 0-5v-6H23l-3.04.434a56.285 56.285 0 0 1-15.92 0L1 3.5H.5Z"/>`),
		g.Group(children),
	)
}