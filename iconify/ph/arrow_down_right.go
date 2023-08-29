package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 88v104a8 8 0 0 1-8 8H88a8 8 0 0 1 0-16h84.69L58.34 69.66a8 8 0 0 1 11.32-11.32L184 172.69V88a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}