package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M56 152V56a8 8 0 0 1 8-8h96a8 8 0 0 1 5.66 13.66L123.31 104l58.35 58.34a8 8 0 0 1-11.32 11.32L112 115.31l-42.34 42.35A8 8 0 0 1 56 152Zm160 56H40a8 8 0 0 0 0 16h176a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}