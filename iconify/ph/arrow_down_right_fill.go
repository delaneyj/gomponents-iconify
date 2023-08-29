package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 88v104a8 8 0 0 1-8 8H88a8 8 0 0 1-5.66-13.66L128.69 140L58.34 69.66a8 8 0 0 1 11.32-11.32L140 128.69l46.34-46.35A8 8 0 0 1 200 88Z"/>`),
		g.Group(children),
	)
}