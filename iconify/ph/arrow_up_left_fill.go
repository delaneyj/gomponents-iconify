package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M197.66 197.66a8 8 0 0 1-11.32 0L116 127.31l-46.34 46.35A8 8 0 0 1 56 168V64a8 8 0 0 1 8-8h104a8 8 0 0 1 5.66 13.66L127.31 116l70.35 70.34a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}