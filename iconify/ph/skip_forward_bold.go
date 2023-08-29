package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 28a12 12 0 0 0-12 12v62L74.55 31A20 20 0 0 0 44 47.88v160.24A20 20 0 0 0 74.55 225L188 154v62a12 12 0 0 0 24 0V40a12 12 0 0 0-12-12ZM68 200.73V55.27L184.3 128Z"/>`),
		g.Group(children),
	)
}