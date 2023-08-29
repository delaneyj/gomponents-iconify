package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.24 13.14a1 1 0 0 0-1.37.36a1 1 0 0 1-1.58.19A.93.93 0 0 1 8 13v-2a1 1 0 0 1 1.88-.48a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.38a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 9 8a3 3 0 0 0-3 3v2a3 3 0 0 0 5.59 1.5a1 1 0 0 0-.35-1.36Zm6 0a1 1 0 0 0-1.37.36a1 1 0 0 1-1.58.19A.93.93 0 0 1 14 13v-2a1 1 0 0 1 1.88-.48a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.38a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 15 8a3 3 0 0 0-3 3v2a3 3 0 0 0 5.59 1.5a1 1 0 0 0-.35-1.36ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}