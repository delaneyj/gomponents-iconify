package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptioningSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 0 1-8-8a7.92 7.92 0 0 1 1.69-4.9L7.2 8.61A3 3 0 0 0 6 11v2a3 3 0 0 0 5.59 1.5a1 1 0 1 0-1.72-1a1 1 0 0 1-1.58.19A.93.93 0 0 1 8 13v-2a1 1 0 0 1 .67-.92L12 13.46A3 3 0 0 0 14.54 16l2.36 2.36A7.92 7.92 0 0 1 12 20Zm6.31-3.1l-1.52-1.52a2.94 2.94 0 0 0 .8-.88a1 1 0 1 0-1.72-1a1 1 0 0 1-.55.41L14 12.59V11a1 1 0 0 1 1.88-.48a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.38a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 12 10.62l-.35-.35a1 1 0 0 0-.1-.79a3.08 3.08 0 0 0-.46-.59a2.94 2.94 0 0 0-1.67-.84L7.1 5.69A7.92 7.92 0 0 1 12 4a8 8 0 0 1 8 8a7.92 7.92 0 0 1-1.69 4.9Z"/>`),
		g.Group(children),
	)
}