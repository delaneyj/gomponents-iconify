package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineDownLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.166 21.986a1.95 1.95 0 0 1-2.154-2.153L3.352 7.77c.181-1.625 2.161-2.32 3.317-1.163l1.268 1.267L13.24 2.57a1.95 1.95 0 0 1 2.758 0L21.427 8a1.95 1.95 0 0 1 0 2.757l-5.304 5.304l1.268 1.267c1.156 1.156.461 3.136-1.164 3.317l-12.061 1.34Z"/>`),
		g.Group(children),
	)
}