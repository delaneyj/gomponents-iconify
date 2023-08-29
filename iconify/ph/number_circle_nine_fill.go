package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleNineFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M145.33 118A20 20 0 1 1 138 90.68a20 20 0 0 1 7.31 27.32ZM232 128A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-86-51.18A36 36 0 1 0 127.94 144q.94 0 1.89-.06l-16.7 28a8 8 0 0 0 2.77 11a8 8 0 0 0 11-2.77L159.18 126A36.05 36.05 0 0 0 146 76.82Z"/>`),
		g.Group(children),
	)
}