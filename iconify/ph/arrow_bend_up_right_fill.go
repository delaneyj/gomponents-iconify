package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendUpRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m229.66 109.66l-48 48A8 8 0 0 1 168 152v-40h-40a88.1 88.1 0 0 0-88 88a8 8 0 0 1-16 0A104.11 104.11 0 0 1 128 96h40V56a8 8 0 0 1 13.66-5.66l48 48a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}